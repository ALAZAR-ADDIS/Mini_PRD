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

	userRep := Repository.NewUserRepo(db, context.TODO())
	passwordService := Infrastructure.NewPasswordService()
	jwtService := Infrastructure.NewJWTService()
	authMiddleware := Infrastructure.NewAuthMiddleware(jwtService)

	userUscase := Usecase.NewUserUscase(userRep, passwordService, jwtService)
	userControler := Controller.NewUserController(userUscase)

	// Authentication routes
	router.POST("/sign-in", userControler.SignInHandler)
	router.POST("/sign-up", userControler.SignUpHandler)
	router.POST("/sign-out", authMiddleware.AuthMiddleware(), userControler.SignOutHandler)
	router.POST("/promote/:id", userControler.PromoteHandler)

}
