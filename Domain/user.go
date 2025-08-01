package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id      primitive.ObjectID `bson:"_id`
	Email    string
	FirstName string
	LastName string
	Password string
	Role     string
	Tokens   string
	UserProfile Profile
	CreatedAt   time.Time         
 	UpdatedAt   time.Time          

}

type Profile struct {
	Bio string
	Gender string
	ProfilePicture string
	PhoneNumber string
	UserAddress Address
}

type Address struct {
	Country     string 
    Region      string 
    City        string 
    Street      string
    PostalCode  string 
}