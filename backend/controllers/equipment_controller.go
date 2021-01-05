package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B6104641/app/ent"
	"github.com/B6104641/app/ent/equipment"
	"github.com/gin-gonic/gin"
)

// EquipmentController struct
type EquipmentController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateEquipment handles POST requests for adding equipment entities
// @Summary Create equipment
// @Description Create equipment
// @ID create-equipment
// @Accept   json
// @Produce  json
// @Param equipment body ent.Equipment true "Equipment entity"
// @Success 200 {object} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments [post]
func (ctl *EquipmentController) CreateEquipment(c *gin.Context) {
	obj := ent.Equipment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "equipment binding failed",
		})
		return
	}

	e, err := ctl.client.Equipment.
		Create().
		SetEQUIPMENTNAME(obj.EQUIPMENTNAME).
		SetEQUIPMENTPRICE(obj.EQUIPMENTPRICE).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving equipment failed",
		})
		return
	}

	c.JSON(200, e)
}

// GetEquipment handles GET requests to retrieve a equipment entity
// @Summary Get a equipment entity by ID
// @Description get equipment by ID
// @ID get-equipment
// @Produce  json
// @Param id path int true "Equipment ID"
// @Success 200 {object} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments/{id} [get]
func (ctl *EquipmentController) GetEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	e, err := ctl.client.Equipment.
		Query().
		Where(equipment.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, e)
}

// ListEquipment handles request to get a list of equipment entities
// @Summary List equipment entities
// @Description list equipment entities
// @ID list-equipment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments [get]
func (ctl *EquipmentController) ListEquipment(c *gin.Context) {
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

	equipments, err := ctl.client.Equipment.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, equipments)
}

// DeleteEquipment handles DELETE requests to delete a equipment entity
// @Summary Delete a equipment entity by ID
// @Description get equipment by ID
// @ID delete-equipment
// @Produce  json
// @Param id path int true "Equipment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments/{id} [delete]
func (ctl *EquipmentController) DeleteEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Equipment.
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

// UpdateEquipment handles PUT requests to update a equipment entity
// @Summary Update a equipment entity by ID
// @Description update equipment by ID
// @ID update-equipment
// @Accept   json
// @Produce  json
// @Param id path int true "Equipment ID"
// @Param equipment body ent.Equipment true "Equipment entity"
// @Success 200 {object} ent.Equipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /equipments/{id} [put]
func (ctl *EquipmentController) UpdateEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Equipment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "equipment binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	e, err := ctl.client.Equipment.
		UpdateOneID(int(id)).
		SetEQUIPMENTNAME(obj.EQUIPMENTNAME).
		SetEQUIPMENTPRICE(obj.EQUIPMENTPRICE).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update Equipment failed",
		})
		return
	}

	c.JSON(200, e)
}

// NewEquipmentController creates and registers handles for the oder controller
func NewEquipmentController(router gin.IRouter, client *ent.Client) *EquipmentController {
	ec := &EquipmentController{
		client: client,
		router: router,
	}

	ec.register()

	return ec

}

func (ctl *EquipmentController) register() {
	equipments := ctl.router.Group("/equipments")

	equipments.POST("", ctl.CreateEquipment)
	equipments.GET(":id", ctl.GetEquipment)
	equipments.GET("", ctl.ListEquipment)
	equipments.DELETE(":id", ctl.DeleteEquipment)
	equipments.PUT(":id", ctl.UpdateEquipment)

}
