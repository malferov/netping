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
	validPortRegex      = `^[1-9][0-9]+$`
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://netping.org"},
		AllowMethods: []string{"GET"},
	}))
	r.GET("/", statusOk)
	g := r.Group("/portcheck")
	{
		g.GET("/version", getVersion)
		g.GET("/v1/:hostname/:port", portCheck)
	}
	return r
}

func main() {
	var serverPort string
	flag.StringVar(&serverPort, "serverPort", "5000", "server listening port")
	flag.Parse()

	pod, _ = os.Hostname()

	router := setupRouter()
	router.Run(":" + serverPort)
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

func portCheck(c *gin.Context) {
	hostname := c.Param("hostname")
	port := c.Param("port")
	glog.Info("portcheck: " + hostname + ":" + port)

	// validate hostname
	var validip = regexp.MustCompile(validIpAddressRegex).MatchString
	var validhn = regexp.MustCompile(validHostnameRegex).MatchString
	var validp = regexp.MustCompile(validPortRegex).MatchString
	glog.Info("ValidIpAddress: " + strconv.FormatBool(validip(hostname)))
	glog.Info("ValidHostname: " + strconv.FormatBool(validhn(hostname)))
	glog.Info("ValidPort: " + strconv.FormatBool(validp(port)))

	if !validp(port) || !validip(hostname) && !validhn(hostname) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed ip address, hostname or port"})
	} else {
		out, err := exec.Command(
			// probe any port via http
			"curl", "-sSv", "--max-time", "1", "--max-filesize", "256", hostname+":"+port,
		).CombinedOutput()

		if err != nil {
			glog.Error("portcheck: " + err.Error())
		}

		msg := "Can't check"
		if len(out) > 0 {
			msg = string(out)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	}
	glog.Infof("portcheck: %d", http.StatusOK)
}
