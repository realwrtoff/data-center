package middleware

import "testing"

func TestGenerateToken(t *testing.T) {
	userName := "jim"
	password := "123456"
	role := 1
	token, err := GenerateToken(userName, password, role)
	if err != nil {
		t.Errorf("%v, error[%s]", token, err.Error())
	} else {
		t.Log(token)
	}
}

func TestParseToken(t *testing.T) {
	userName := "jim"
	password := "123456"
	role := 1
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJqaW0iLCJwYXNzd29yZCI6IjEyMzQ1NiIsInJvbGUiOjEsImV4cCI6MTU5Nzc0ODQzNCwiaXNzIjoiZ2luLWJsb2cifQ.IuIKDo9nptcUXzQ3nBxTO17nvEdGAfhYhaayhZLFmDI"
	token, _ := GenerateToken(userName, password, role)
	claim, err := ParseToken(token)
	if err != nil {
		t.Errorf("error[%s]", err.Error())
	} else {
		if claim.UserName != userName {
			t.Errorf("user name [%s] != [%s]", claim.UserName, userName)
		}
		if claim.Password != password {
			t.Errorf("user password [%s] != [%s]", claim.UserName, userName)
		}
		if claim.Role != role {
			t.Errorf("user role [%s] != [%s]", claim.UserName, userName)
		}
		t.Logf("claim [%v]", claim)
	}
}

func TestJWT(t *testing.T) {
	JWT()
}
