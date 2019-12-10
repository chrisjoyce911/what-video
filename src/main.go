package main

import (
	"log"

	"net/http"
	"what-video/internal/config"
	"what-video/internal/database"
	"what-video/internal/keywords"
	"what-video/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {

	c := config.NewConfig().Env()
	var err error

	database := database.NewDB(
		c.DbUser,
		c.DbPassword,
		c.DbHost,
		c.DbName,
		c.DbPort)

	defer database.DB.Close()

	err = keywords.Migrate(*database)
	if err != nil {
		log.Panicln(err)
	}

	err = users.Migrate(*database)
	if err != nil {
		log.Panicln(err)
	}

	// users

	// uu := []users.User{}
	// err = db.Select(&uu, "SELECT * FROM users")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("users...")
	// log.Println(uu)

	gin.ForceConsoleColor()
	r := setupRouter()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	err = r.Run(":80")
	if err != nil {
		log.Println(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	//apiRouter := r.Group("/api")

	// apiRouter.POST("/roles", CreateRole)
	// apiRouter.POST("/users", CreateUser)
	// apiRouter.PUT("/users/:user_id", UpdateUser)
	// apiRouter.DELETE("/users/:user_id", DeleteUser)

	// apiRouter.GET("/role/:user_id", GetUserRole)
	// apiRouter.GET("/users/:user_id", GetUser)

	return r
}

// func CreateRole(context *gin.Context) {
// 	var reqdata CreateRoleReq
// 	err := context.BindJSON(&reqdata)
// 	if err != nil {
// 		err.Error()
// 	}

// 	role := &Role{
// 		Slug:  reqdata.Slug,
// 		Title: reqdata.Title,
// 	}
// 	newRole := db.Create(role)
// 	context.JSON(200, newRole)
// }

// func CreateUser(context *gin.Context) {
// 	var reqdata CreateUserReq
// 	err := context.BindJSON(&reqdata)
// 	if err != nil {
// 		err.Error()
// 	}

// 	var roles []Role
// 	db.Where("slug IN (?)", reqdata.Roles).Find(&roles)

// 	user := &users.User{UserName: reqdata.Mobile}
// 	newUser := db.Create(user)

// 	// Instead of line 30 & 31 for creating newUser
// 	// you can also use these two line of codes(34 & 35).
// 	//user := &User{Mobile:reqdata.Mobile}
// 	//newUser := db.Create(&user).Association("Roles").Append(roles)

// 	context.JSON(201, newUser)
// }

// func GetUser(context *gin.Context) {
// 	userId := context.Param("user_id")

// 	var roles []Role
// 	var user users.User
// 	db.Where("ID=  ?", userId).First(&user)
// 	db.Model(&user).Association("Roles").Find(&roles)

// 	//user.Roles = roles
// 	context.JSON(200, user)
// }

// func GetUserRole(context *gin.Context) {
// 	userId := context.Param("user_id")

// 	var roles []Role
// 	var user users.User
// 	db.Where("ID=  ?", userId).First(&user)
// 	db.Model(&user).Association("Roles").Find(&roles)

// 	context.JSON(200, roles)
// }

// func UpdateUser(context *gin.Context) {
// 	userId := context.Param("user_id")
// 	var reqdata UpdateUserReq
// 	err := context.BindJSON(&reqdata)
// 	if err != nil {
// 		err.Error()
// 	}

// 	var user users.User
// 	db.Where("ID=  ?", userId).First(&user)

// 	var roles []Role
// 	db.Where("slug IN (?)", reqdata.Roles).Find(&roles)

// 	user.UserName = reqdata.Mobile
// 	db.Save(&user)
// 	db.Model(&user).Association("Roles").Replace(roles)

// 	context.JSON(200, user)
// }

// func DeleteUser(context *gin.Context) {
// 	userId := context.Param("user_id")

// 	var user users.User
// 	db.Where("ID=  ?", userId).First(&user)
// 	db.Delete(&user)
// 	db.Model(&user).Association("Roles").Clear()

// 	context.JSON(200, user)
// }

func AddNumber(a int, b int) string {

	log.Printf("%d %d", a, b)
	return "2"
}
