package Repository
import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func ConnectDb(password string,userName string) *mongo.Database{

	clientOption := options.Client().ApplyURI("mongodb+srv://" + userName + ":" + password + "@cluster0.xolx70t.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	Database := client.Database("TaskManager")

	fmt.Println("Database Connected Successfuly")

	return Database

}


func Disconnect(db *mongo.Database){
	err := db.Client().Disconnect(context.TODO())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connection Closed")
}