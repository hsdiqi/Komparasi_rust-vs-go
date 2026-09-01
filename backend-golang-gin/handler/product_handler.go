package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"backend-golang-gin/model"
	"backend-golang-gin/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct { service *service.ProductService }
func NewProductHandler(s *service.ProductService) *ProductHandler { return &ProductHandler{service: s} }

func parsePageLimit(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 { page = 1 }
	if limit < 1 { limit = 10 }
	if limit > 100 { limit = 100 }
	return page, limit
}

func (h *ProductHandler) List(c *gin.Context) {
	page, limit := parsePageLimit(c)
	data, err := h.service.List(c.Request.Context(), page, limit)
	if err != nil { c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	c.JSON(http.StatusOK, model.ProductListResponse{Status:"success", Data:data, Meta:model.Meta{Page:page, Limit:limit}})
}

func (h *ProductHandler) Detail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data, err := h.service.Detail(c.Request.Context(), id)
	if err == sql.ErrNoRows { c.JSON(http.StatusNotFound, model.ErrorResponse{Status:"error", Message:"product not found"}); return }
	if err != nil { c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	c.JSON(http.StatusOK, model.ProductResponse{Status:"success", Data:data})
}

func (h *ProductHandler) Search(c *gin.Context) {
	page, limit := parsePageLimit(c)
	keyword := c.DefaultQuery("keyword", "")
	data, err := h.service.Search(c.Request.Context(), keyword, page, limit)
	if err != nil { c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	c.JSON(http.StatusOK, model.ProductListResponse{Status:"success", Data:data, Meta:model.Meta{Page:page, Limit:limit}})
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req model.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	data, err := h.service.Create(c.Request.Context(), req)
	if err != nil { c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	c.JSON(http.StatusCreated, model.ProductResponse{Status:"success", Data:data})
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req model.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	data, err := h.service.Update(c.Request.Context(), id, req)
	if err == sql.ErrNoRows { c.JSON(http.StatusNotFound, model.ErrorResponse{Status:"error", Message:"product not found"}); return }
	if err != nil { c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	c.JSON(http.StatusOK, model.ProductResponse{Status:"success", Data:data})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.service.Delete(c.Request.Context(), id); err != nil { c.JSON(http.StatusInternalServerError, model.ErrorResponse{Status:"error", Message:err.Error()}); return }
	c.JSON(http.StatusOK, model.ProductResponse{Status:"success", Data:nil})
}
