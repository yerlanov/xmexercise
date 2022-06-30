package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yerlanov/xmexercise/internal/company"
	"github.com/yerlanov/xmexercise/internal/company/ipapi"
	"github.com/yerlanov/xmexercise/internal/company/middleware"
	"github.com/yerlanov/xmexercise/internal/company/service"
	"net/http"
	"strconv"
)

type Handler struct {
	Service service.Service
	IpApi   ipapi.Service
}

func (h *Handler) Register(router *gin.RouterGroup) {
	router.POST("/", middleware.LocationAccessMiddleware(middleware.AuthMiddleware(h.CreateCompany), h.IpApi))
	router.DELETE("/:id", middleware.LocationAccessMiddleware(middleware.AuthMiddleware(h.DeleteCompany), h.IpApi))
	router.PUT("/:id", h.UpdateCompany)
	router.GET("/list", h.ListCompany)
	router.GET("/filter", h.ListWithFilter)
}

func (h *Handler) CreateCompany(c *gin.Context) {
	var req company.Company
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res, err := h.Service.Create(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) DeleteCompany(c *gin.Context) {
	req := c.Param("id")

	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("int64 required")))
		return
	}

	affected, err := h.Service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success",
		"message": fmt.Sprintf("%d rows affected", affected)})
}

func (h *Handler) UpdateCompany(c *gin.Context) {
	reqId := c.Param("id")

	id, err := strconv.ParseInt(reqId, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("int64 required")))
		return
	}

	var body company.Company
	if err = c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updated, err := h.Service.Update(c, body, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *Handler) ListCompany(c *gin.Context) {
	list, err := h.Service.List(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) ListWithFilter(c *gin.Context) {
	params := c.Request.URL.Query()
	filter := make(map[string]string)

	for key, values := range params {
		filter[key] = values[0]
	}

	list, err := h.Service.ListWithFilter(c, filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, list)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
