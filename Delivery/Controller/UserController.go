package Controller

import (
	"Mini_PRD/Usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase Usecase.UserUsecase
}

func NewUserController(userUsecase Usecase.UserUsecase) *UserController {
	return &UserController{userUsecase}
}

func (UC *UserController) SignInHandler(ctx *gin.Context) {

}

func (UC *UserController) SignUpHandler(ctx *gin.Context) {

}
func (UC *UserController) SignOutHandler(ctx *gin.Context) {

}
func (UC *UserController) PromoteHandler(ctx *gin.Context) {

}
