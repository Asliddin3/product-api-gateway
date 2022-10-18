package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Asliddin3/product-api-gateway/genproto/store"
	l "github.com/Asliddin3/product-api-gateway/pkg/logger"
)


// @Summary get store
// @Description this func get store
// @Tags store
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /store/{id} [get]
func (h *handlerV1) GetStore(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body:=&store.GetStoreInfoById{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.StoreService().GetStore(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get store", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}


// @BasePath /api/v1

// PingExample godoc
// @Summary create store
// @Description this func create store with ful info
// @Tags store
// @Accept json
// @Produce json
// @Param product body store.StoreRequest true "Store"
// @Success 200 {object} store.StoreResponse
// @Router /store [post]
func (h *handlerV1) CreateStore(c *gin.Context) {
	var (
		body        store.StoreRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.StoreService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create store", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}
