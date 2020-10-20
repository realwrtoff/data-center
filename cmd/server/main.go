package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/hpifu/go-kit/hconf"
	"github.com/hpifu/go-kit/henv"
	"github.com/hpifu/go-kit/hflag"
	"github.com/hpifu/go-kit/hrule"
	"github.com/hpifu/go-kit/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/olivere/elastic/v7"
	"github.com/realwrtoff/data-center/internal/router"
	"github.com/realwrtoff/data-center/internal/scheduler"
	"github.com/realwrtoff/data-center/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// AppVersion name
var AppVersion = "unknown"

type Options struct {
	Service struct {
		Port string `hflag:"usage: service port"  hrule:"atLeast 1"`
	}
	Redis struct {
		Addr     string `hflag:"usage: redis addr" hrule:"atLeast 10"`
		Password string `hflag:"usage: redis password"`
	}
	Es struct {
		Uri string `hflag:"usage: elasticsearch address"`
	}
	Mysql struct {
		Host     string `hflag:"usage: mysql host" hrule:"atLeast 7"`
		Port     string `hflag:"usage: mysql port" hrule:"atLeast 1"`
		UserName string `hflag:"usage: mysql user" hrule:"atLeast 1"`
		Password string `hflag:"usage: mysql password"`
		DbName   string `hflag:"usage: mysql dbname" hrule:"atLeast 1"`
	}
	Logger struct {
		Run logger.Options
	}
}

func main() {
	version := hflag.Bool("v", false, "print current version")
	configfile := hflag.String("c", "configs/server.json", "config file path")
	if err := hflag.Bind(&Options{}); err != nil {
		panic(err)
	}
	if err := hflag.Parse(); err != nil {
		panic(err)
	}
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	// load config
	options := &Options{}
	config, err := hconf.New("json", "local", *configfile)
	if err != nil {
		panic(err)
	}
	if err := config.Unmarshal(options); err != nil {
		panic(err)
	}
	if err := henv.NewHEnv("SERV").Unmarshal(options); err != nil {
		panic(err)
	}
	if err := hflag.Unmarshal(options); err != nil {
		panic(err)
	}
	if err := hrule.Evaluate(options); err != nil {
		panic(err)
	}

	runLog, err := logger.NewLogger(&options.Logger.Run)
	if err != nil {
		panic(err)
	}

	rds := redis.NewClient(&redis.Options{
		Addr:         options.Redis.Addr,
		Password:     options.Redis.Password,
		MaxRetries:   1,
		MinIdleConns: 1,
	})
	if _, err := rds.Ping().Result(); err != nil {
		panic(err)
	}
	runLog.Infof("ping redis %v ok\n", options.Redis.Addr)

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", options.Mysql.UserName, options.Mysql.Password, options.Mysql.Host, options.Mysql.Port, options.Mysql.DbName)
	mdb, err := gorm.Open("mysql", dbLink)
	if err != nil {
		panic(err)
	}
	// 设置连接池信息
	// mdb.DB().SetMaxIdleConns(10)
	// mdb.DB().SetConnMaxLifetime(100)

	es, err := elastic.NewClient(
		elastic.SetURL(options.Es.Uri),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	watcher := scheduler.NewWatcher(rds)
	go watcher.Run()
	// init services
	svc := service.NewService(rds, mdb, es, runLog)

	// init gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Pub", "Authorization", "Accept", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	// set handler
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})
	router.InitCompanyRouter(r, svc)
	router.InitProjectRouter(r, svc)

	// run server
	server := &http.Server{
		Addr:    options.Service.Port,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// graceful quit
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	runLog.Infof("%v shutdown ...", os.Args[0])

	_ = mdb.Close()
	_ = rds.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		runLog.Errorf("%v shutdown fail or timeout", os.Args[0])
		return
	}
	_ = runLog.Out.(*rotatelogs.RotateLogs).Close()
}
