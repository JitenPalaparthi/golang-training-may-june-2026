package handlers

import (
	"encoding/json"
	"http-demo/helpers"
	"http-demo/models"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	FileName string
}

func Mewuser(filename string) *User {
	return &User{filename}
}

func (u *User) CreateUser(ctx *gin.Context) {

	// bytes, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
	// 	ctx.Abort()
	// 	return
	// }

	// var user models.User

	// err = json.Unmarshal(bytes, &user)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
	// 	ctx.Abort()
	// 	return
	// }

	// username := ctx.GetHeader("username")
	// password := ctx.GetHeader("password")

	// if username == "admin" && password == "admin123" {
	// 	ctx.Next()
	// } else {
	// 	slog.Warn("unauthorized resource access")
	// 	ctx.String(http.StatusUnauthorized, "unauthorized access")
	// 	ctx.Abort()
	// }

	var user models.User

	err := ctx.Bind(&user) // unmarshal
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
		ctx.Abort()
		return
	}

	user.Id = rand.IntN(10000)
	user.Status = "active"
	user.LastModified = uint64(time.Now().Unix())

	bytes, err := json.Marshal(user)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
		ctx.Abort()
		return
	}

	err = helpers.SaveUser(u.FileName, append(bytes, byte('\n')))
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
		ctx.Abort()
		return
	}
	slog.Debug("user successfully saved")
	//ctx.String(201, "User Successfully saved")
	ctx.JSON(201, struct {
		Message string
		User    models.User
	}{Message: "Success", User: user})
}

func (u *User) GetUsers(ctx *gin.Context) {
	users, err := helpers.GetUsers(u.FileName)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, "oops! something went wrong, try again or consult admin")
		ctx.Abort()
		return
	}
	ctx.JSON(200, users)
}
