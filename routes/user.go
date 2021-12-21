package routes

import (
	"errors"

	"github.com/TranThaiNgoc/Fiber-GORM/database"
	"github.com/TranThaiNgoc/Fiber-GORM/models"
	"github.com/gofiber/fiber/v2"
)

//khoi tao bien "struct" co nhieu kieu du lieu trong do
type User struct {
	// This is not the model, more like a serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//dat ten la "user" tham so truyen vao la models.User trong file models, User la gia tri tra ve
//khai bao func bao gom ten ham, ten bien -> kieu bien, kieu tra ve
//co mot dang khac nua la method goi la receiver argument
//// Struct type - `Rectangle`
// type Rectangle struct {
// 	X, Y float64
// }

// // Method with receiver `Rectangle`
// func (p Rectangle) Acreage() float64 {
// 	return p.Y * p.X
// }
//giong nhau deu la tai su dung bien struct nhung cach dung khac nhau
//func tra ve se la kieu du lieu bien struc
//method reciver argument dung bien struct de su dung trong ham

func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

//con tro (Pointer)
//cu phap "&" dung de lay dia chi cua bien do
//cu phap "*a" dung de luu dia chi cua bien do
//su khac biet giua toan tu va pointer la toan tu so sanh chi co the thay doi trong func vi no la ban nhap
//con pointer la thay doi bang cach lay dia chi cua bien do va thay doi value
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	// đặt 2 biến trong params
	id, err := c.ParamsInt("id")

	// đặt biến user trong model.User
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findUser(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"fist_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findUser(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Successfully deleted User")
}
