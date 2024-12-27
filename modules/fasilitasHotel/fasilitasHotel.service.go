package fasilitashotel

import (
	"booking-hotel/databases"
	"booking-hotel/libs"
	"booking-hotel/modules/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CreateFasilitasHotel godoc
// @Summary Create Fasilitas Hotel
// @Description Create Fasilitas Hotel
// @Tags Fasilitas Hotel
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param CreateFasilitasHotel body RequestFasilitasHotel true "Create Fasilitas Hotel"
// @Success 201 {string} string "Success Create Fasilitas Hotel"
// @Router /1.0/fasilitas-hotel [post]
func CreateFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 && claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitasHotel := FasilitasHotel{}
	if err := c.BodyParser(&fasilitasHotel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err := validate.Struct(fasilitasHotel)
	if err != nil {
		err := err.(validator.ValidationErrors)
		errors := map[string]string{}
		for _, err := range err {
			errors[err.Field()] = err.Tag()
		}
		return libs.ResponseError(c, errors, 400)
	}
	// checkHotel := make(chan bool)
	// checkFasilitas := make(chan bool)
	// checkInFasilitasHotel := make(chan bool)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go checkFasilitasHotelIsExist(fasilitasHotel.Hotel_id, fasilitasHotel.Fasilitas_id, &wg, checkInFasilitasHotel)

	// wg.Add(1)
	// go checkIsFasilitasExist(fasilitasHotel.Fasilitas_id, &wg, checkFasilitas)

	// wg.Add(1)
	// go checkIsHotelExist(fasilitasHotel.Hotel_id, &wg, checkHotel)

	// go func() {
	//     wg.Wait()
	//     close(checkHotel)
	//     close(checkFasilitas)
	//     close(checkInFasilitasHotel)
	// }()

	// hotelExists := <-checkHotel
	// fasilitasExists := <-checkFasilitas
	// inFasilitasHotel := <-checkInFasilitasHotel

	// if !hotelExists && !fasilitasExists  {
	// 	return libs.ResponseError(c, "Hotel and Fasilitas not found", 404)
	// }
	// if !hotelExists {
	//     return libs.ResponseError(c, "Hotel not found", 404)
	// }
	// if !fasilitasExists {
	//     return libs.ResponseError(c, "Fasilitas not found", 404)
	// }
	// if inFasilitasHotel {
	//     return libs.ResponseError(c, "Fasilitas hotel already exist", 400)
	// }

	if claims.(*user.Claims).Hak_akses_id == 3 {
		fasilitasHotel.Hotel_id = claims.(*user.Claims).Hotel_id
	}

	if err := databases.DB.Table("fasilitas_hotel").Create(&fasilitasHotel); err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	return libs.ResponseSuccess(c, "Success create fasilitas hotel", 201)
}

// func checkFasilitasHotelIsExist(hotel_id int, fasilitas_id int, wg *sync.WaitGroup, ch chan bool) {
// 	var fasilitasHotel []FasilitasHotel
// 	defer wg.Done()
// 	if err := databases.DB.Table("fasilitas_hotel").Where("hotel_id = ? AND fasilitas_id = ?", hotel_id, fasilitas_id).Find(&fasilitasHotel); err.Error != nil {
// 		ch <- false
// 		return
// 	} else if len(fasilitasHotel) != 0 {
// 		ch <- true
// 		return
// 	}

// }

// func checkIsHotelExist(hotel_id int, wg *sync.WaitGroup, ch chan bool) {
// 	var hotel hotel.Hotel
// 	defer wg.Done()
// 	if err := databases.DB.Table("hotel").First(&hotel, hotel_id); err.Error != nil {
// 		ch <- false
// 		return
// 	} else if err.RowsAffected == 0 {
// 		ch <- false
// 		return
// 	} else {
// 		ch <- true
// 		return
// 	}
// }

// func checkIsFasilitasExist(fasilitas_id int, wg *sync.WaitGroup, ch chan bool) {
// 	var fasilitas fasilitas.Fasilitas
// 	defer wg.Done()
// 	if err := databases.DB.Table("fasilitas").First(&fasilitas, fasilitas_id); err.Error != nil {
// 		ch <- false
// 		return
// 	} else if err.RowsAffected == 0 {
// 		ch <- false
// 		return
// 	} else {
// 		ch <- true
// 		return
// 	}
// }

// GetAllFasilitasHotel godoc
// @Summary Get All Fasilitas Hotel
// @Description Get All Fasilitas Hotel
// @Tags Fasilitas Hotel
// @Accept json
// @Produce json
// @Success 200 {object} ResponseFasilitasHotel
// @Router /1.0/fasilitas-hotel [get]
func GetAllFasilitasHotel(c *fiber.Ctx) error {
	var fasilitasHotel []ResponseFasilitasHotel
	if err := databases.DB.Preload("Hotel").Preload("Fasilitas").Table("fasilitas_hotel").Find(&fasilitasHotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}

// GetFasilitasHotelById godoc
// @Summary Get Fasilitas Hotel By Id
// @Description Get Fasilitas Hotel By Id
// @Tags Fasilitas Hotel
// @Accept json
// @Produce json
// @Param id path int true "Fasilitas Hotel ID"
// @Success 200 {object} ResponseFasilitasHotel
// @Router /1.0/fasilitas-hotel/{id} [get]
func GetFasilitasHotelById(c *fiber.Ctx) error {
	fasilitasHotel := ResponseFasilitasHotel{}
	id := c.Params("id")
	err := databases.DB.Preload("Hotel").Preload("Fasilitas").Table("fasilitas_hotel").First(&fasilitasHotel, id)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}

// UpdateFasilitasHotel godoc
// @Summary Update Fasilitas Hotel
// @Description Update Fasilitas Hotel
// @Tags Fasilitas Hotel
// @Accept json
// @Produce json
// @Param id path int true "Fasilitas Hotel ID"
// @Param Authorization header string true "Bearer"
// @Param UpdateFasilitasHotel body RequestFasilitasHotel true "Update Fasilitas Hotel"
// @Success 200 {string} string "Success Update Fasilitas Hotel"
// @Router /1.0/fasilitas-hotel/{id} [put]
func UpdateFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 && claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	fasilitasHotel := FasilitasHotel{}
	if err := c.BodyParser(&fasilitasHotel); err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	}
	id := c.Params("id")

	// checkHotel := make(chan bool)
	// checkFasilitas := make(chan bool)
	// checkInFasilitasHotel := make(chan bool)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go checkFasilitasHotelIsExist(fasilitasHotel.Hotel_id, fasilitasHotel.Fasilitas_id, &wg, checkInFasilitasHotel)

	// wg.Add(1)
	// go checkIsFasilitasExist(fasilitasHotel.Fasilitas_id, &wg, checkFasilitas)

	// wg.Add(1)
	// go checkIsHotelExist(fasilitasHotel.Hotel_id, &wg, checkHotel)

	// go func() {
	// 	wg.Wait()
	// 	close(checkHotel)
	// 	close(checkFasilitas)
	// 	close(checkInFasilitasHotel)
	// }()

	// hotelExists := <-checkHotel
	// fasilitasExists := <-checkFasilitas
	// inFasilitasHotel := <-checkInFasilitasHotel

	// if !hotelExists && !fasilitasExists {
	// 	return libs.ResponseError(c, "Hotel and Fasilitas not found", 404)
	// }
	// if !hotelExists {
	// 	return libs.ResponseError(c, "Hotel not found", 404)
	// }
	// if !fasilitasExists {
	// 	return libs.ResponseError(c, "Fasilitas not found", 404)
	// }
	// if inFasilitasHotel {
	// 	return libs.ResponseError(c, "Fasilitas hotel already exist", 400)
	// }

	err := databases.DB.Table("fasilitas_hotel").Where("fasilitas_hotel_id = ?", id).Updates(&fasilitasHotel)
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success update fasilitas hotel", 200)
}

// DeleteFasilitasHotel godoc
// @Summary Delete Fasilitas Hotel
// @Description Delete Fasilitas Hotel
// @Tags Fasilitas Hotel
// @Accept json
// @Produce json
// @Param id path int true "Fasilitas Hotel ID"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} string "Success Delete Fasilitas Hotel"
// @Router /1.0/fasilitas-hotel/{id} [delete]
func DeleteFasilitasHotel(c *fiber.Ctx) error {
	claims := c.Locals("user")
	if claims == nil {
		return libs.ResponseError(c, "Unauthorized", 401)
	}
	if claims.(*user.Claims).Hak_akses_id != 1 && claims.(*user.Claims).Hak_akses_id != 3 {
		return libs.ResponseError(c, "Forbidden", 403)
	}
	id := c.Params("id")
	err := databases.DB.Table("fasilitas_hotel").Where("fasilitas_hotel_id = ?", id).Delete(&FasilitasHotel{})
	if err.Error != nil {
		return libs.ResponseError(c, err.Error.Error(), 400)
	}
	if err.RowsAffected == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, "Success delete fasilitas hotel", 200)
}

// GetFasilitasByHotelId godoc
// @Summary Get Fasilitas By Hotel Id
// @Description Get Fasilitas By Hotel Id
// @Tags Fasilitas Hotel
// @Accept json
// @Produce json
// @Param id path int true "Hotel ID"
// @Success 200 {object} ResponseFasilitasHotel
// @Router /1.0/fasilitas-hotel/by-hotel/{id} [get]
func GetFasilitasByHotelId(c *fiber.Ctx) error {
	id := c.Params("id")
	var fasilitasHotel []ResponseFasilitasHotel
	if err := databases.DB.Preload("Fasilitas").Preload("Hotel").Table("fasilitas_hotel").Where("hotel_id = ?", id).Find(&fasilitasHotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	} else if len(fasilitasHotel) == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}

// GetFasilitasByFasilitasId godoc
// @Summary Get Fasilitas By Fasilitas Id
// @Description Get Fasilitas By Fasilitas Id
// @Tags Fasilitas Hotel
// @Accept json
// @Produce json
// @Param id path int true "Fasilitas ID"
// @Success 200 {object} ResponseFasilitasHotel
// @Router /1.0/fasilitas-hotel/by-fasilitas/{id} [get]
func GetFasilitasByFasilitasId(c *fiber.Ctx) error {
	id := c.Params("id")
	var fasilitasHotel []ResponseFasilitasHotel
	if err := databases.DB.Preload("Fasilitas").Preload("Hotel").Table("fasilitas_hotel").Where("fasilitas_id = ?", id).Find(&fasilitasHotel).Error; err != nil {
		return libs.ResponseError(c, err.Error(), 400)
	} else if len(fasilitasHotel) == 0 {
		return libs.ResponseError(c, "Data not found", 404)
	}
	return libs.ResponseSuccess(c, fasilitasHotel, 200)
}
