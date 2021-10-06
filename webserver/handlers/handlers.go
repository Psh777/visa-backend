package handlers

import (
	"github.com/Psh777/gin_handler_answer"
	"github.com/Psh777/visa-backend/singleton"
	"github.com/gin-gonic/gin"
	"time"
)

func IndexHandler(w *gin.Context) {
	//answer.HandlerError(w, "0")
	uptime := singleton.GetUptime()
	v, n := singleton.GetData()
	answer.HandlerInterface(w, Out{
		Uptime:  time.Now().Unix() - uptime,
		Now:     time.Now(),
		NowUTC:  time.Now().UTC(),
		Version: v,
		Node:    n,
	})
}

type Out struct {
	Node    string    `json:"node"`
	Version string    `json:"version"`
	Uptime  int64     `json:"uptime"`
	Now     time.Time `json:"now"`
	NowUTC  time.Time `json:"now_utc"`
}
