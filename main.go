package main

import (
	"bytes"
	"developers_today_test/db"
	"developers_today_test/server"
	"developers_today_test/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	log.Printf("[%d] \t %s \t Request: %s \t response: %s\n", c.Writer.Status(), c.Request.Method, c.Request.RequestURI, blw.body.String())
}

func main() {
	var config struct {
		DSN string `json:"sqlite_dsn"`
	}
	err := utils.ReadConfig(&config)
	if err != nil {
		log.Panicln(err)
	}
	err = db.SetHandler(config.DSN)
	if err != nil {
		log.Panicln(err)
	}
	r := gin.New()
	r.Use(Logger)
	server.SetServer(r)
	log.Fatalln(r.Run(":8080"))
}
