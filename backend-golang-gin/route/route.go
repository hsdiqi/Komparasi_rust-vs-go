package route

import (
	"backend-golang-gin/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handler.ProductHandler) {
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status":"ok"}) })
	r.GET("/products", h.List)
	r.GET("/products/search", h.Search)
	r.GET("/products/:id", h.Detail)
	r.POST("/products", h.Create)
	r.PUT("/products/:id", h.Update)
	r.DELETE("/products/:id", h.Delete)
}
