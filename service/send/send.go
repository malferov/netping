package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"net/smtp"
	"os"
)

type Payload struct {
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
	Email   string `json:"email" binding:"required"`
}

var (
	sha     = "unknown"
	version = "dev"
	date    = "unknown"
	pod     string
	email   string
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
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
	err := c.BindJSON(&pld)
	if err != nil {
		glog.Error("Invalid payload" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		msg := []byte("From: send@netping.org\r\n" +
			"To: " + email + "\r\n" +
			"Subject: message from netping.org\r\n" +
			"\r\n" +
                        "Subject: " + pld.Subject + "\r\n" +
			"Message: " + pld.Message + "\r\n" +
                        "Email: " + pld.Email + "\r\n")
		err := smtp.SendMail("gmail-smtp-in.l.google.com:25", nil, pld.Email, []string{email}, msg)
		if err != nil {
			glog.Error("Can't send email" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "submitted",
			})
		}
	}
	glog.Infof("submit: %d", http.StatusOK)
}
