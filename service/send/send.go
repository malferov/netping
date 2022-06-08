package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"net/smtp"
	"os"
)

type Payload struct {
	Subject string `binding:"required"`
	Message string `binding:"required"`
	Email   string `binding:"required,email"`
	Channel string `binding:"required"`
}

var (
	sha     = "unknown"
	version = "dev"
	date    = "unknown"
	pod     string
	email   string
	from    = "send@netping.org"
	mxs     = "mxs.mail.ru:25" //"gmail-smtp-in.l.google.com:25"
	tok     string
	channel = "general"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://netping.org"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))
	r.GET("/", statusOk)
	g := r.Group("/send")
	{
		g.GET("/version", getVersion)
		g.POST("/submit", submit)
	}
	return r
}

func main() {
	var port string
	flag.StringVar(&port, "port", "5000", "server listening port")
	flag.Parse()

	pod, _ = os.Hostname()
	email = os.Getenv("EMAIL")
	tok = os.Getenv("BOT_TOKEN")

	router := setupRouter()
	router.Run(":" + port)
}

func getVersion(c *gin.Context) {
	body := gin.H{
		"sha":      sha,
		"version":  version,
		"date":     date,
		"hostname": pod,
		"ginmode":  gin.Mode(),
		"lang":     "golang",
	}
	c.JSON(http.StatusOK, body)
}

func statusOk(c *gin.Context) {
	c.Status(http.StatusOK)
}

func submit(c *gin.Context) {
	var pld Payload
	var err error
	err = c.BindJSON(&pld)
	if err != nil {
		glog.Error("Invalid payload: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		header := []byte("From: " + from + "\r\n" +
			"To: " + email + "\r\n" +
			"Subject: message from netping.org\r\n" + "\r\n")
		msg := []byte("Subject: " + pld.Subject + "\r\n" +
			"Message: " + pld.Message + "\r\n" +
			"Email: " + pld.Email + "\r\n")
		if pld.Channel == "email" {
			err = smtp.SendMail(mxs, nil, from, []string{email}, append(header, msg...))
		} else if pld.Channel == "slack" {
			var request *http.Request
			var response *http.Response
			post := struct {
				Channel string `json:"channel"`
				Text    string `json:"text"`
			}{
				Channel: channel,
				Text:    string(msg),
			}
			var jsonPost []byte
			jsonPost, err = json.Marshal(post)
			if err != nil {
				glog.Error("Marshal post: " + err.Error())
			} else {
				glog.Infof("Marshal jsonPost: %+v", string(jsonPost))
				request, err = http.NewRequest("POST",
					"https://slack.com/api/chat.postMessage", bytes.NewBuffer(jsonPost))
				request.Header.Set("Content-Type", "application/json; charset=utf-8")
				request.Header.Set("Authorization", "Bearer "+tok)
				client := &http.Client{}
				response, err = client.Do(request)
				if err != nil {
					glog.Error("Request post: " + err.Error())
				} else {
					defer response.Body.Close()
					if response.StatusCode == http.StatusOK {
						var body struct {
							Ok    bool
							Error string
						}
						json.NewDecoder(response.Body).Decode(&body)
						glog.Infof("Response status code: %d Body: %+v", response.StatusCode, body)
						if !body.Ok {
							err = fmt.Errorf(body.Error)
						}
					} else {
						err = fmt.Errorf("Response status code: %d", response.StatusCode)
					}
				}
			}
		} else {
			glog.Error("Invalid payload: " + "unknown channel")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "unknown channel",
			})
			return
		}

		if err != nil {
			glog.Error("Can't send via <" + pld.Channel + ">: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			glog.Infof("submit: %d", http.StatusOK)
			c.JSON(http.StatusOK, gin.H{
				"message": "submitted",
			})
		}
	}
}
