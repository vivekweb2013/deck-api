package httpservice

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vivekweb2013/deck-api/internal/config"
)

// Run starts the http server.
func Run(c config.Config) error {
	gin.SetMode(gin.ReleaseMode)
	if c.HTTP.Debug {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	address := net.JoinHostPort(c.HTTP.Host, c.HTTP.Port)
	server := http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()
}
