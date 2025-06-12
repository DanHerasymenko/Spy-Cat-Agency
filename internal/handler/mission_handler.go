package handler

import (
	"SpyCatAgency/internal/model"
	"SpyCatAgency/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MissionHandler struct {
	service *service.MissionService
}

func NewMissionHandler(service *service.MissionService) *MissionHandler {
	return &MissionHandler{service: service}
}

func (h *MissionHandler) RegisterRoutes(router *gin.Engine) {
	missions := router.Group("/api/missions")
	{
		missions.POST("", h.Create)
		missions.PUT("/:id", h.Update)
		missions.DELETE("/:id", h.Delete)
		missions.GET("/:id", h.GetByID)
		missions.GET("", h.List)
		missions.POST("/:id/assign", h.AssignCat)
		missions.POST("/:id/targets", h.AddTarget)
		missions.DELETE("/targets/:id", h.DeleteTarget)
		missions.PUT("/targets/:id", h.UpdateTarget)
	}
}

// @Summary Create a new mission
// @Description Create a mission with 1–3 targets
// @Tags Missions
// @Accept json
// @Produce json
// @Param body body model.MissionCreate true "Mission create body"
// @Success 201 {object} model.Mission
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions [post]
func (h *MissionHandler) Create(ctx *gin.Context) {
	var create model.MissionCreate
	if err := ctx.ShouldBindJSON(&create); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mission, err := h.service.Create(ctx.Request.Context(), create)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, mission)
}

// @Summary Create a new mission
// @Description Create a mission with 1–3 targets
// @Tags Missions
// @Accept json
// @Produce json
// @Param body body model.MissionCreate true "Mission create body"
// @Success 201 {object} model.Mission
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions [post]
func (h *MissionHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var update model.MissionUpdate
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mission, err := h.service.Update(ctx.Request.Context(), uint(id), update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mission)
}

// @Summary Delete a mission
// @Description Delete a mission by ID (only if not assigned to a cat)
// @Tags Missions
// @Produce plain
// @Param id path int true "Mission ID"
// @Success 204 {string} string "No content"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions/{id} [delete]
func (h *MissionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get mission by ID
// @Description Retrieve mission details, including assigned cat and targets
// @Tags Missions
// @Produce json
// @Param id path int true "Mission ID"
// @Success 200 {object} model.Mission
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions/{id} [get]
func (h *MissionHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	mission, err := h.service.GetByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mission)
}

// @Summary List all missions
// @Description Get all created missions
// @Tags Missions
// @Produce json
// @Success 200 {array} model.Mission
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions [get]
func (h *MissionHandler) List(ctx *gin.Context) {
	missions, err := h.service.List(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, missions)
}

// @Summary Assign cat to mission
// @Description Assign a cat to a mission (1 cat per mission)
// @Tags Missions
// @Accept json
// @Produce plain
// @Param id path int true "Mission ID"
// @Param body body model.CatAssign true "Cat assign body"
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions/{id}/assign [post]
func (h *MissionHandler) AssignCat(ctx *gin.Context) {
	missionID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission id"})
		return
	}

	var request model.CatAssign

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignCat(ctx.Request.Context(), uint(missionID), request.CatID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Add target to mission
// @Description Add a target to an existing mission (only if mission is not completed)
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path int true "Mission ID"
// @Param body body model.TargetCreate true "Target create body"
// @Success 201 {object} model.Target
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions/{id}/targets [post]
func (h *MissionHandler) AddTarget(ctx *gin.Context) {
	missionID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission id"})
		return
	}

	var targetCreate model.TargetCreate
	if err := ctx.ShouldBindJSON(&targetCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target, err := h.service.AddTarget(ctx.Request.Context(), uint(missionID), targetCreate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, target)
}

// @Summary Delete target
// @Description Delete a target by ID (only if target is not completed)
// @Tags Missions
// @Produce plain
// @Param id path int true "Target ID"
// @Success 204 {string} string "No content"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions/targets/{id} [delete]
func (h *MissionHandler) DeleteTarget(ctx *gin.Context) {
	targetID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid target id"})
		return
	}

	if err := h.service.DeleteTarget(ctx.Request.Context(), uint(targetID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// @Summary Update target
// @Description Update notes or completion status for a target
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path int true "Target ID"
// @Param body body model.TargetUpdate true "Target update body"
// @Success 200 {object} model.Target
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/missions/targets/{id} [put]
func (h *MissionHandler) UpdateTarget(ctx *gin.Context) {
	targetID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid target id"})
		return
	}

	var update model.TargetUpdate
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target, err := h.service.UpdateTarget(ctx.Request.Context(), uint(targetID), update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, target)
}
