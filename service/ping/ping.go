package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

var (
	sha                 = "unknown"
	version             = "dev"
	date                = "unknown"
	pod                 string
	validIpAddressRegex = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	validHostnameRegex  = `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://netping.org"},
		AllowMethods: []string{"GET"},
	}))
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

func ping(c *gin.Context) {
	hostname := c.Param("hostname")
	glog.Info("ping: " + hostname)

	// validate hostname
	var validip = regexp.MustCompile(validIpAddressRegex).MatchString
	var validhn = regexp.MustCompile(validHostnameRegex).MatchString
	glog.Info("ValidIpAddress: " + strconv.FormatBool(validip(hostname)))
	glog.Info("ValidHostname: " + strconv.FormatBool(validhn(hostname)))

	if !validip(hostname) && !validhn(hostname) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed ip address or hostname"})
	} else {
		out, err := exec.Command("ping", "-c3", "-i0.2", "-W2", hostname).CombinedOutput()
		if err != nil {
			glog.Error("ping: " + err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": string(out),
		})
	}
	glog.Infof("ping: %d", http.StatusOK)
}
