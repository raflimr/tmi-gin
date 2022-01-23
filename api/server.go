package api

import (
	"fmt"
	"os"
	"time"

	db "tmi-gin/db/sqlc"
	"tmi-gin/token"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DBDriver            string
	DBSource            string
	ServerAddress       string
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
	Email               string
	Password            string
}

// Server serves HTTP requests for our banking service.
type Server struct {
	config     Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config Config, store db.Store) (*Server, error) {
	tokenKey := os.Getenv("TOKEN_SYMMETRIC_KEY")

	if tokenKey == "" {
		tokenKey = "-"
	}

	tokenMaker, err := token.NewPasetoMaker(tokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/mahasiswas", server.createMahasiswa)
	router.POST("/mahasiswas/login", server.loginMahasiswa)
	router.GET("/mahasiswas/home-dashboard", server.homeDashboard)
	router.GET("/mahasiswas/practice", server.listPractices)
	router.GET("/mahasiswas/practice/:id_category", server.listPracticesByCategory)
	router.POST("/mahasiswas/forgot-password", server.postEmailOtp)
	router.PUT("/mahasiswas/forgot-password/send", server.updateOTP)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.PUT("/mahasiswas/update/:id", server.updateMahasiswa)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
