package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"os"
	"regexp"
)

var (
	version = "dev"
	date    = "unknown"
	pod     string
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", statusOk)
	g := r.Group("/ping")
	{
		g.GET("/version", getVersion)
		g.GET("/v1/:hostname", ping)
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

func ping(c *gin.Context) {
	hostname := c.Param("hostname")
	glog.Info("ping: " + hostname)
	// validate hostname
	var alpha = regexp.MustCompile(`^[[:alpha:]]+$`).MatchString
	if !alpha(hostname) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed hostname"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "hostname is " + hostname,
		})
	}
	glog.Infof("ping: %d", http.StatusOK)
}
