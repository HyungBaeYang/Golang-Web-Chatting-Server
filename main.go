package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/websocket"
	"github.com/unrolled/render"

	"github.com/julienschmidt/httprouter"
)

const (
	// 어플리케이션에서 사용할 세션의 키 정보
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"

	socketBufferSize = 1024
)

var (
	//mongoSession *mgo.mongoSession
	renderer = render.New()
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
)

func main() {

	// 라우터 생성
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "index", map[string]interface{}{"host": r.Host})
	})

	ne := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	ne.Use(sessions.Sessions(sessionKey, store))

	ne.UseHandler(router)

	ne.Run(":3000")

}
