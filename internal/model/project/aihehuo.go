package project

import (
	"encoding/json"
	"time"
)

type AiHeHuoProject struct {
	ProjectId       int         `json:"project_id"`
	Title           string      `json:"title"`
	PublisherId     int         `json:"publisher_id"`
	PublisherName   string      `json:"publisher_name"`
	Province        string      `json:"province"`
	City            string      `json:"city"`
	PartnersCount   string      `json:"partners_count"`
	TeamMembers     string      `json:"team_members"`
	Investment      string      `json:"investment"`
	ReviewerComment string      `json:"reviewer_comment"`
	FollowersCount  int         `json:"followers_count"`
	CommentsCount   int         `json:"comments_count"`
	ViewsCount      int         `json:"views_count"`
	Recruit         int         `json:"recruit"`
	Description     string      `json:"description"`
	Bp              interface{} `json:"bp,omitempty"`
	Cover           string      `json:"cover,omitempty"`
	Images          interface{} `json:"images,omitempty"`
	JobDescriptions interface{} `json:"job_descriptions,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
}

type AiHeHuoBP struct {
	Name                string `json:"name"`
	CreatedAtFormatted  string `json:"created_at_formatted"`
	TotalViewers        int    `json:"total_viewers"`
	Permission          string `json:"permission"`
	PermissionFormatted string `json:"permission_formatted"`
	Alert               struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Buttons     []struct {
			Title    string `json:"title"`
			Redirect struct {
				CheckVerification bool `json:"check_verification"`
			} `json:"redirect"`
		} `json:"buttons"`
	} `json:"alert"`
}

type AiHeHuoBPJobDescription struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Idea struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"idea"`
	Description string `json:"description"`
	Role        string `json:"role"`
	Exp         string `json:"exp"`
	Benefit     string `json:"benefit"`
	CanRemote   bool   `json:"can_remote"`
	CanPt       bool   `json:"can_pt"`
	Title       string `json:"title"`
	Tags        []struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"tags"`
	Inquiries struct {
		Total int `json:"total"`
		Data  []struct {
			ID                 int       `json:"id"`
			CreatedAtFormatted string    `json:"created_at_formatted"`
			Avatar             string    `json:"avatar"`
			IsNew              time.Time `json:"is_new"`
		} `json:"data"`
	} `json:"inquiries"`
	CanEnquire       bool   `json:"can_enquire"`
	ShortQuestionOne string `json:"short_question_one"`
	ShortQuestionTwo string `json:"short_question_two"`
}

type AiHeHuoPublisher struct {
	PublisherID    int         `json:"publisher_id"`
	Name           string      `json:"name"`
	Province       string      `json:"province"`
	City           string      `json:"city"`
	Industry       string      `json:"industry"`
	Experience     string      `json:"experience"`
	StatusRole     string      `json:"status_role"`
	LastAccessedAt string      `json:"last_accessed_at"`
	Bio            string      `json:"bio"`
	Verified       int         `json:"verified"`
	ZhimaVerified  int         `json:"zhima_verified"`
	IsVip          int         `json:"is_vip"`
	AdminFeatured  int         `json:"admin_featured"`
	FeaturedAt     time.Time   `json:"featured_at"`
	AdminApproval  string      `json:"admin_approval"`
	Avatar         string      `json:"avatar"`
	Skills         interface{} `json:"skills"`
	CreatedAt      time.Time   `json:"created_at"`
}

type AiHeHuoSkill struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

//值传递 不可修改字段
func (AiHeHuoProject) TableName() string {
	return "aihehuo_project"
}

func (AiHeHuoPublisher) TableName() string {
	return "aihehuo_publisher"
}

// 引用传递，可修改字段，修改json.dumps的字段
func (ap *AiHeHuoProject)JsonDeal() {
	bpBytes := ap.Bp.([]byte)
	if len(bpBytes) > 0 {
		bp := AiHeHuoBP{}
		_ = json.Unmarshal(bpBytes, &bp)
		ap.Bp = bp
	}
	imgBytes := ap.Images.([]byte)
	if len(imgBytes) > 0 {
		var images []string
		_ = json.Unmarshal(imgBytes, &images)
		ap.Images = images
	}
	jdsBytes := ap.JobDescriptions.([]byte)
	if len(jdsBytes) > 0 {
		var jds []AiHeHuoBPJobDescription
		_ = json.Unmarshal(jdsBytes, &jds)
		ap.JobDescriptions = jds
	}
}

// 引用传递
func (ap *AiHeHuoPublisher)JsonDeal() {
	bts := ap.Skills.([]byte)
	if len(bts) > 0 {
		var skills []AiHeHuoSkill
		_ = json.Unmarshal(bts, &skills)
		ap.Skills = skills
	}
}
