package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DiscountHandler struct {
	DiscountUsecase domain.DiscountUsecase
}

func NewDiscountHandler(e *gin.Engine, discountUsecase domain.DiscountUsecase) {
	handler := &DiscountHandler{
		DiscountUsecase: discountUsecase,
	}

	e.POST("/api/v1/seasoning-discount", handler.AddSeasoningDiscount)
	e.POST("/api/v1/product/:productID/event-discount", handler.AddEventDiscount)

	store := e.Group("/api/v1/store/:storeID")
	{
		store.GET("/discounts", handler.GetAllShippingByStoreId)
		store.POST("/shipping-discount", handler.AddShippingDiscount)
	}
}

func (s *DiscountHandler) GetAllShippingByStoreId(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	discounts, err := s.DiscountUsecase.GetShippingByStoreID(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}
	c.JSON(200, discounts)

}

func (s *DiscountHandler) AddSeasoningDiscount(c *gin.Context) {
	var discount swagger.SeasoningDiscount

	if err := c.BindJSON(&discount); err != nil {
		logrus.Error(err)
		return
	}

	err := s.DiscountUsecase.AddSeasoning(c, &discount)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (s *DiscountHandler) AddShippingDiscount(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	var discount swagger.ShippingDiscount

	if err := c.BindJSON(&discount); err != nil {
		logrus.Error(err)
		return
	}

	err = s.DiscountUsecase.AddShipping(c, &discount, storeID)

	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (s *DiscountHandler) AddEventDiscount(c *gin.Context) {
	productID, err := strconv.ParseInt(c.Param("productID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	var discount swagger.EventDiscount

	if err := c.BindJSON(&discount); err != nil {
		logrus.Error(err)
		return
	}

	err = s.DiscountUsecase.AddEvent(c, &discount, productID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}
