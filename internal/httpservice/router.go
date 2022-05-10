package httpservice

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vivekweb2013/deck-api/internal/config"
	"github.com/vivekweb2013/deck-api/internal/deck"
)

// Run starts the http server.
func Run(c config.Config) error {
	gin.SetMode(gin.ReleaseMode)
	if c.HTTP.Debug {
		gin.SetMode(gin.DebugMode)
	}

	createDeckService := deck.NewCreateDeckService()
	createDeckHandler := NewCreateDeckHandler(createDeckService)
	openDeckService := deck.NewOpenDeckService()
	openDeckHandler := NewOpenDeckHandler(openDeckService)
	drawCardsService := deck.NewDrawCardsService()
	drawCardsHandler := NewDrawCardsHandler(drawCardsService)

	router := gin.Default()
	v1 := router.Group("api/v1")
	v1.POST("/decks", createDeckHandler.CreateDeck)
	v1.GET("/decks/:id", openDeckHandler.OpenDeck)
	v1.POST("/decks/:id/draw", drawCardsHandler.DrawCards)

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
