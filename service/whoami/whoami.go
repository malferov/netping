package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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
	r.Use(cors.Default())
	r.GET("/", statusOk)
	g := r.Group("/whoami")
	{
		g.GET("/version", getVersion)
		g.GET("/ip", getIp)
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

func getIp(c *gin.Context) {
	glog.Info("whoami: begin")
	c.JSON(http.StatusOK, gin.H{
		"message": c.ClientIP(),
	})
	glog.Infof("whoami: %d", http.StatusOK)
}
