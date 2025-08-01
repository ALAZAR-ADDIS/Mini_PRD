package Router

import (
	"Mini_PRD/Delivery/Controller"
	"Mini_PRD/Infrastructure"
	"Mini_PRD/Repository"
	"Mini_PRD/Usecase"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRoute(db *mongo.Database, router *gin.Engine) {

	userRep := Repository.NewUserRepo(db,context.TODO())
	userInf := Infrastructure.NewUserServices()
	userUscase := Usecase.NewUserUscase(userRep,userInf)
	userControler := Controller.NewUserController(*userUscase)
	router.POST("/sing-in",userControler.SignInHandler)
	router.POST("/sing-up",userControler.SignUpHandler)
	router.POST("/sing-out",userControler.SignOutHandler)
	router.POST("/promote/:id",userControler.PromoteHandler)
	

}