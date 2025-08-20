package jarvis

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ChatRequest represents the incoming chat request
type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

// ChatResponse represents the chat response
type ChatResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

// Server holds the Jarvis instance and handles HTTP requests
type Server struct {
	jarvis *Jarvis
	mutex  sync.Mutex
}

// NewServer creates a new server instance
func NewServer() *Server {
	return &Server{
		jarvis: Start(),
	}
}

// StartServer starts the HTTP server on the specified port
func (s *Server) StartServer(port string) error {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Configure CORS to allow requests from frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"} // Common Vite/React ports
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Chat endpoint
	r.POST("/chat", s.handleChat)

	log.Printf("Starting server on port %s", port)
	return r.Run(":" + port)
}

// handleChat processes chat requests
func (s *Server) handleChat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ChatResponse{
			Error: "Invalid request format: " + err.Error(),
		})
		return
	}

	// Use mutex to ensure thread safety for chat history
	s.mutex.Lock()
	defer s.mutex.Unlock()

	response, err := s.jarvis.Chat(req.Message)
	if err != nil {
		log.Printf("Chat error: %v", err)
		c.JSON(http.StatusInternalServerError, ChatResponse{
			Error: "Failed to process chat: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ChatResponse{
		Response: response,
	})
}
