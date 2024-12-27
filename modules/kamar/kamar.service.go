package kamar

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CreateKamar godoc
// @Summary Create Kamar
// @Description Create Kamar
// @Tags Kamar
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreateKamar body RequestKamar true "Create Kamar"
// @Success 201 {string} string "Success Create Kamar"
// @Router /1.0/kamar [post]
func CreateKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 && claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	kamar := RequestKamar{}

	if err := c.BodyParser(&kamar); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	if kamar.Hotel_id != claims.(*user.Claims).Hotel_id {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	if err := validate.Struct(kamar); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	kamar.Status_kamar_id = 1
	if err := databases.DB.Table("kamar").Create(&kamar).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create kamar", 201)
}

// GetAllKamar godoc
// @Summary Get All Kamar
// @Description Get All Kamar
// @Tags Kamar
// @Accept json
// @Produce json
// @Success 200 {object} Kamar
// @Router /1.0/kamar [get]
func GetAllKamar(c *fiber.Ctx) error {
	var kamar []Kamar
	if err := databases.DB.Preload("Status_kamar").Table("kamar").Find(&kamar).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, kamar, 200)
}

// GetKamarById godoc
// @Summary Get Kamar By ID
// @Description Get Kamar By ID
// @Tags Kamar
// @Accept json
// @Produce json
// @Param id path int true "Kamar ID"
// @Success 200 {object} Kamar
// @Router /1.0/kamar/{id} [get]
func GetKamarById(c *fiber.Ctx) error {
	kamar := Kamar{}
	id := c.Params("id")
	err := databases.DB.Preload("Status_kamar").Table("kamar").First(&kamar, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, kamar, 200)
}

// UpdateKamar godoc
// @Summary Update Kamar
// @Description Update Kamar
// @Tags Kamar
// @Accept json
// @Produce json
// @Param id path int true "Kamar ID"
// @Param Authorization header string true "Bearer"
// @Param UpdateKamar body RequestKamar true "Update Kamar"
// @Success 200 {string} string "Success Update Kamar"
// @Router /1.0/kamar/{id} [put]
func UpdateKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 && claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	kamar := RequestKamar{}
	id := c.Params("id")

	if err := c.BodyParser(&kamar); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}

	if kamar.Hotel_id != claims.(*user.Claims).Hotel_id {
		return libs.ResponseError(c, "Forbidden", 403)
	}

	if err := validate.Struct(kamar); err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, e := range err {
			errors[e.Field()] = e.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	kamar.Status_kamar_id = 1
	err := databases.DB.Table("kamar").Where("kamar_id = ?", id).Updates(&kamar)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update kamar", 200)
}

// DeleteKamar godoc
// @Summary Delete Kamar
// @Description Delete Kamar
// @Tags Kamar
// @Accept json
// @Produce json
// @Param id path int true "Kamar ID"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} string "Success Delete Kamar"
// @Router /1.0/kamar/{id} [delete]
func DeleteKamar(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 && claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	kamar := Kamar{}
	id := c.Params("id")
	err := databases.DB.Table("kamar").Where("kamar_id = ?", id).Delete(&kamar)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete kamar", 200)
}
