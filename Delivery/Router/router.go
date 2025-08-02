package Router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoute(db *mongo.Database, router *gin.Engine) {
	NewUserRoute(db, router)
}
