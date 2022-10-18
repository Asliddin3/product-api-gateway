package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	// "github.com/Asliddin3/product-api-gateway/api/handlers/models"
	"github.com/Asliddin3/product-api-gateway/genproto/product"
	l "github.com/Asliddin3/product-api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary create products with info
// @Description this func create products with ful info
// @Tags product
// @Accept json
// @Produce json
// @Param product body product.ProductFullInfo true "Products"
// @Success 200 {object} product.ProductFullInfoResponse
// @Router /products/fullinfo [post]
func (h *handlerV1) CreateProductFullInfo(c *gin.Context) {
	var (
		body        product.ProductFullInfo
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
	response, err := h.serviceManager.ProductService().CreateProduct(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete product", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary get all products
// @Description this func get all products
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} product.Products
// @Router /products/info [get]
func (h *handlerV1) GetProducts(c *gin.Context) {

	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.ProductService().GetProducts(ctx, &product.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete product", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary create type
// @Description this func create type
// @Tags type
// @Accept json
// @Produce json
// @Param type body product.TypeRequest true "Type"
// @Success 200 {object} product.Type
// @Router /products/type [post]
func (h *handlerV1) CreateType(c *gin.Context) {
	var (
		body        product.TypeRequest
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

	response, err := h.serviceManager.ProductService().CreateType(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create type", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)

}

// @Summary update product
// @Description this func update product
// @Tags product
// @Accept json
// @Produce json
// @Param request body product.Product true "Products"
// @Success 200 {object} product.Product
// @Router /products/update [patch]
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	var (
		body        product.Product
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
	response, err := h.serviceManager.ProductService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update product", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}


// PingExample godoc
// @Summary delete products
// @Description this func delete products
// @Tags product
// @Accept json
// @Produce json
// @Param request body product.GetProductId true "Product"
// @Success 200 "success" product.Empty
// @Router /product/delete/{id} [delete]
func (h *handlerV1) DeleteProduct(c *gin.Context) {
	var (
		body        product.GetProductId
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
	response, err := h.serviceManager.ProductService().DeleteProduct(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete product", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}


// @Summary create category
// @Description this func create category
// @Tags category
// @Accept json
// @Produce json
// @Param type body product.CategoryRequest true "Category"
// @Success 200 {object} product.Category
// @Router /products/category [post]
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		body        product.CategoryRequest
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

	response, err := h.serviceManager.ProductService().CreateCategory(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create category", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)

}

// @Summary create products
// @Description this func create products
// @Tags product
// @Accept json
// @Produce json
// @Param request body product.ProductRequest true "Products"
// @Success 200 {object} product.ProductResponse
// @Router /products [post]
func (h *handlerV1) CreateProduct(c *gin.Context) {
	var (
		body        product.ProductRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	fmt.Println(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.ProductService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create product", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// PingExample godoc
// @Summary get products
// @Description this func get products
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /products/{id} [get]
func (h *handlerV1) GetProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.ProductService().GetProduct(
		ctx, &product.GetProductId{
			Id: id,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get product", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
