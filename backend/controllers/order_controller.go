package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/B6104641/app/ent"
	"github.com/B6104641/app/ent/company"
	"github.com/B6104641/app/ent/equipment"
	"github.com/B6104641/app/ent/order"
	"github.com/B6104641/app/ent/user"
	"github.com/gin-gonic/gin"
)

//OrderController struct
type OrderController struct {
	client *ent.Client
	router gin.IRouter
}

//Order struct
type Order struct {
	Equipment int
	Company   int
	User      int
	Amount    string
	Price     int
	Added     string
}

// CreateOrder handles POST requests for adding order entities
// @Summary Create order
// @Description Create order
// @ID create-order
// @Accept   json
// @Produce  json
// @Param order body Order true "Order entity"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders [post]
func (ctl *OrderController) CreateOrder(c *gin.Context) {
	obj := Order{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "order binding failed",
		})
		return
	}

	e, err := ctl.client.Equipment.
		Query().
		Where(equipment.IDEQ(int(obj.Equipment))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "equipment not found",
		})
		return
	}

	cp, err := ctl.client.Company.
		Query().
		Where(company.IDEQ(int(obj.Company))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "company not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}
	a, err := strconv.Atoi(obj.Amount)
	time, err := time.Parse(time.RFC3339, obj.Added)
	o, err := ctl.client.Order.
		Create().
		SetEquipment(e).
		SetCompany(cp).
		SetUser(u).
		SetAMOUNT(a).
		SetPRICE(obj.Price).
		SetADDEDTIME(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, o)
}

// DeleteOrder handles DELETE requests to delete a order entity
// @Summary Delete a order entity by ID
// @Description get order by ID
// @ID delete-order
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [delete]
func (ctl *OrderController) DeleteOrder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Order.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// GetOrder handles GET requests to retrieve a order entity
// @Summary Get a order entity by ID
// @Description get order by ID
// @ID get-order
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [get]
func (ctl *OrderController) GetOrder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	o, err := ctl.client.Order.
		Query().
		Where(order.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, o)
}

// UpdateOrder handles PUT requests to update a order entity
// @Summary Update a order entity by ID
// @Description update order by ID
// @ID update-order
// @Accept   json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body ent.Order true "Order entity"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [put]
func (ctl *OrderController) UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.Order{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "order binding failed",
		})
		return
	}
	obj.ID = int(id)
	o, err := ctl.client.Order.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, o)
}

// ListOrder handles request to get a list of order entities
// @Summary List order entities
// @Description list order entities
// @ID list-order
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders [get]
func (ctl *OrderController) ListOrder(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	Orders, err := ctl.client.Order.
		Query().
		WithEquipment().
		WithCompany().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, Orders)
}

// NewOrderController creates and registers handles for the order controller
func NewOrderController(router gin.IRouter, client *ent.Client) *OrderController {
	oc := &OrderController{
		client: client,
		router: router,
	}

	oc.register()

	return oc

}

func (ctl *OrderController) register() {
	orders := ctl.router.Group("/orders")
	orders.GET("", ctl.ListOrder)
	// CRUD
	orders.POST("", ctl.CreateOrder) //
	orders.GET(":id", ctl.GetOrder)
	orders.PUT(":id", ctl.UpdateOrder)
	orders.DELETE(":id", ctl.DeleteOrder)//
}
