package handler

import (
	"SpyCatAgency/internal/model"
	"SpyCatAgency/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CatHandler struct {
	service *service.CatService
}

func NewCatHandler(service *service.CatService) *CatHandler {
	return &CatHandler{service: service}
}

func (h *CatHandler) RegisterRoutes(router *gin.Engine) {
	cats := router.Group("/api/cats")
	{
		cats.POST("/create", h.Create)
		cats.PUT("/:id/salary", h.Update)
		cats.DELETE("/:id", h.Delete)
		cats.GET("/:id", h.GetByID)
		cats.GET("/list", h.List)
	}
}

// @Summary Create a new spy cat
// @Description Create a new spy cat
// @Tags Cats
// @Accept json
// @Produce json
// @Param body body model.CatCreate true "CreateCat request body"
// @Success 201 {object} model.Cat
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/cats/create [post]
func (h *CatHandler) Create(c *gin.Context) {
	var create model.CatCreate
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat, err := h.service.Create(c.Request.Context(), create)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cat)
}

// @Summary Update cat salary
// @Description Update salary of a spy cat by ID
// @Tags Cats
// @Accept json
// @Produce json
// @Param id path int true "Cat ID"
// @Param body body model.CatUpdate true "Update salary request body"
// @Success 200 {object} model.Cat
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/cats/{id}/salary [put]
func (h *CatHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var update model.CatUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat, err := h.service.Update(c.Request.Context(), uint(id), update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cat)
}

// @Summary Delete a spy cat
// @Description Remove a spy cat by ID
// @Tags Cats
// @Produce plain
// @Param id path int true "Cat ID"
// @Success 204 {string} string "No content"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/cats/{id} [delete]
func (h *CatHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(ctx.Request.Context(), uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// @Summary Get a spy cat
// @Description Retrieve a spy cat by ID
// @Tags Cats
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {object} model.Cat
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/cats/{id} [get]
func (h *CatHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	cat, err := h.service.GetByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cat)
}

// @Summary List all spy cats
// @Description Retrieve a full list of all registered spy cats in the agency
// @Tags Cats
// @Produce json
// @Success 200 {array} model.Cat
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/cats/list [get]
func (h *CatHandler) List(c *gin.Context) {
	cats, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cats)
}
