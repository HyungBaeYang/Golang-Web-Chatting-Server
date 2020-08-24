package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
)

const (
	currentUserKey  = "oauth2_current_user" // 세션에 저장 되는 User Key
	sessionDuration = time.Hour             // 로그인 세션 유지 시간.
)

type User struct {
	UUID      string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"user"`
	AvatarUrl string    `json:"avatar_url"`
	Expired   time.Time `json:"expired"`
}

func (u *User) Vaild() bool {

	// 현재 시간 기준으로 만료되는 시간 확인.
	return u.Expired.Sub(time.Now()) > 0
}

func (u *User) Refresh() {
	// 만료 시간 연장 시킴.
	u.Expired = time.Now().Add(sessionDuration)
}

func GetCurrentUser(r *http.Request) *User {
	s := sessions.GetSession(r)

	if s.Get(currentUserKey) == nil {
		fmt.Print("Empty Session Current User Key ! ")
		return nil
	}
	data := s.Get(currentUserKey).([]byte)
	var u User
	json.Unmarshal(data, &u)
	return &u
}

func SetCurrentUser(r *http.Request, u *User) {
	if u != nil {
		u.Refresh()
	}

	// 세션에 CurrentUser 정보를 Json 형태로 저장함.
	s := sessions.GetSession(r)
	val, _ := json.Marshal(u)
	s.Set(currentUserKey, val)
}
