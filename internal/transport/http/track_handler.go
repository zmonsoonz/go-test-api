package transport_http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	"github.com/zmonsoonz/go-test-api/internal/platform/httpx"
	"github.com/zmonsoonz/go-test-api/internal/transport/middleware"
)

type TrackHandler struct {
	trackService ports.TrackService
}

func NewTrackHandler(trackService ports.TrackService) *TrackHandler {
	return &TrackHandler{trackService: trackService}
}

func (h *TrackHandler) InitTrackRoutes(r gin.IRouter) {
	tracks := r.Group("/tracks")
	{
		tracks.POST("/", h.Create)
		tracks.GET("/:id", h.GetTrackById)
		tracks.GET("/", h.GetAll)
		tracks.DELETE("/:id", h.Delete)
	}
}
func (h *TrackHandler) Create(c *gin.Context) {
	var input domain.Track
	userId, err := middleware.GetUserId(c);
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
         
	id, err := h.trackService.Create(input, userId)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *TrackHandler) GetAll(c *gin.Context)  {
	userId, err := middleware.GetUserId(c);
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
         
	tracks, err := h.trackService.GetAll(userId)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": tracks,
	})
}

func (h *TrackHandler) GetTrackById(c *gin.Context)  {
	userId, err := middleware.GetUserId(c);
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	track, err := h.trackService.GetById(id, userId)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, track)
}

func (h *TrackHandler) Delete(c *gin.Context)  {
	userId, err := middleware.GetUserId(c);
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.trackService.Delete(id, userId)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, httpx.StatusResponse{
		Status: "ok",
	})
}