package Domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `bson:"_id`
	Email    string
	Password string
	Role     string
	tokens   string
	UserProfile Profile
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