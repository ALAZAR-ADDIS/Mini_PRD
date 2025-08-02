package Repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Db      *mongo.Database
	context context.Context
}

func NewUserRepo(db *mongo.Database, context context.Context) *UserRepo {
	return &UserRepo{db, context}
}

//Crude operation methodes for the User
