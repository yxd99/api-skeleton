package handlers

import (
	"api-skeleton/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GreetHandler struct {
	GreetService services.GreetService
}

func NewGreetHandler() *GreetHandler {
	return &GreetHandler{
		GreetService: *services.NewGreetService(),
	}
}

func (h *GreetHandler) Greet(c *gin.Context) {
	name := c.Query("name")
	message := h.GreetService.GenerateGreet(name)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
