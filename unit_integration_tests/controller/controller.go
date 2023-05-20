package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abelwhite/pigsty/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://pigsty:pigsty@cluster0.yf0s3uc.mongodb.net/?retryWrites=true&w=majority"
const dbName = "pigsty"
const colName = "pigs"

// collection or a reference of mongo db
var collection *mongo.Collection

// connect with mongoDB
// create an init method which runs a first time
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOptions) //context for accesses an outside server
	//context is for how long for the request and how long it goes off,
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("collection instance is ready")
}

//MongoDb helpers - file

// insert 1 record
// create room usecase
func createRoom(room model.Room) {
	inserted, err := collection.InsertOne(context.Background(), room) //inserting data from room
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created 1 room in db with id: ", inserted.InsertedID)
}

// update 1 room use case
// func updateRoom(roomId string) {
// 	id, _ := primitive.ObjectIDFromHex(roomId)
// 	filter := bson.M{"_id": id}                             //filter to find what will be updated
// 	update := bson.M{"$set": bson.M{"name": "roomupdated"}} //this will update

// 	result, err := collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//how many values are updated
// 	fmt.Println("modified count: ", result.MatchedCount)

// }

// delete 1 record
func deleteOneRoom(roomId string) {
	id, _ := primitive.ObjectIDFromHex(roomId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Room got deleted with delete count: ", deleteCount)
}

// delete all records
func deleteAllRoom() int64 {
	//empty means everything is selected
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of rooms deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all rooms from database
// view Room Use Case: View Room
func viewRoom() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var rooms []primitive.M //use a loop to view the rooms

	for cursor.Next(context.Background()) {
		var room bson.M
		err := cursor.Decode(&room)
		if err != nil {
			log.Fatal(err)
		}
		//push rooms into room
		rooms = append(rooms, room)
	}
	defer cursor.Close(context.Background())
	return rooms

}

// actual controllers - file
func ViewRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allRooms := viewRoom()
	json.NewEncoder(w).Encode(allRooms)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") //only allow certain types

	var room model.Room
	_ = json.NewDecoder(r.Body).Decode(&room)
	createRoom(room)
	json.NewEncoder(w).Encode(room)
}

func DeleteOneRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow certain types

	params := mux.Vars(r)
	deleteOneRoom(params["id"]) //delete the movie
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow certain types

	count := deleteAllRoom()
	json.NewEncoder(w).Encode(count)
}

// --------------FOR PIGS--------------
// Use case
func createPig(pig model.Pig) {
	inserted, err := collection.InsertOne(context.Background(), pig) //inserting data from room
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created 1 pig in db with id: ", inserted.InsertedID)
}

// delete 1 record
func deleteOnePig(roomId string) {
	id, _ := primitive.ObjectIDFromHex(roomId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pig got deleted with delete count: ", deleteCount)
}

// delete all records
func deleteAllPigs() int64 {
	//empty means everything is selected
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of Pigs deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all rooms from database
// view Room Use Case: View Room
func viewPig() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var pigs []primitive.M //use a loop to view the rooms

	for cursor.Next(context.Background()) {
		var pig bson.M
		err := cursor.Decode(&pig)
		if err != nil {
			log.Fatal(err)
		}
		//push rooms into room
		pigs = append(pigs, pig)
	}
	defer cursor.Close(context.Background())
	return pigs

}

// actual controllers - file
func ViewPig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allRooms := viewPig()
	json.NewEncoder(w).Encode(allRooms)
}

func CreatePig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") //only allow certain types

	var pig model.Pig
	_ = json.NewDecoder(r.Body).Decode(&pig)
	createPig(pig)
	json.NewEncoder(w).Encode(pig)
}

func DeleteOnePig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow certain types

	params := mux.Vars(r)
	deleteOnePigsty(params["id"]) //delete the pig with the id
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllPigs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow certain types

	count := deleteAllPigsty()
	json.NewEncoder(w).Encode(count)
}

//--------For PIGSTY----

// Testing for use cases 7-8
func createPigsty(pigsty model.Pigsty) {
	inserted, err := collection.InsertOne(context.Background(), pigsty) //inserting data from room
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created 1 pigsty in db with id: ", inserted.InsertedID)
}

// delete 1 record
func deleteOnePigsty(pigstyId string) {
	id, _ := primitive.ObjectIDFromHex(pigstyId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pig got deleted with delete count: ", deleteCount)
}

// delete all records
func deleteAllPigsty() int64 {
	//empty means everything is selected
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of Pigsty deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all rooms from database
// view Room Use Case: View Room
func viewPigsty() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var pigstys []primitive.M //use a loop to view the rooms

	for cursor.Next(context.Background()) {
		var pigsty bson.M
		err := cursor.Decode(&pigsty)
		if err != nil {
			log.Fatal(err)
		}
		//push rooms into room
		pigstys = append(pigstys, pigsty)
	}
	defer cursor.Close(context.Background())
	return pigstys

}

// actual controllers - file
func ViewPigsty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allPigstys := viewPigsty()
	json.NewEncoder(w).Encode(allPigstys)
}

func CreatePigsty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") //only allow certain types

	var pigsty model.Pigsty
	_ = json.NewDecoder(r.Body).Decode(&pigsty)
	createPigsty(pigsty)
	json.NewEncoder(w).Encode(pigsty)
}

func DeleteOnePigsty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow certain types

	params := mux.Vars(r)
	deleteOnePigsty(params["id"]) //delete the pig with the id
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllPigsty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE") //only allow certain types

	count := deleteAllPigsty()
	json.NewEncoder(w).Encode(count)
}
