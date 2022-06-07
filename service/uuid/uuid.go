package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"net/http"
	"os"
)

var (
	sha     = "unknown"
	version = "dev"
	date    = "unknown"
	pod     string
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://netping.org"},
		AllowMethods: []string{"GET"},
	}))
	r.GET("/", statusOk)
	g := r.Group("/uuid")
	{
		g.GET("/version", getVersion)
		g.POST("/generate", generate)
	}
	return r
}

func main() {
	var port string
	flag.StringVar(&port, "port", "5000", "server listening port")
	flag.Parse()
	pod, _ = os.Hostname()
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

func generate(c *gin.Context) {
	uuid := uuid.NewString()
	c.JSON(http.StatusOK, gin.H{
		"uuid": uuid,
	})
	glog.Infof("uuid: %s", uuid)
}
