package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodb/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb connection string. it will be used to connect with mongodb
const connectionString = "mongodb+srv://fahadpathan56:fahadpathan@cluster0.oxoqi6z.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix" // database name
const colName = "watchlist" // collection name

var collection *mongo.Collection // collection object/instance. by using this variable we can perform CRUD operations on "watchlist" collection

// connect with mongodb
func init() {
	clientOptions := options.Client().ApplyURI(connectionString) // clientOptions is used to set options for client. it is needed to connect with mongodb.
	client, err := mongo.Connect(context.TODO(), clientOptions) // connect with mongodb. it returns client and error. context.TODO() will be used in place of context.Background(). It is used to avoid memory leaks.
	// context.TODO() works like context.Background(). It returns empty context.Background() if the Go version is greater than 1.7 otherwise it returns context.Background(). we directly didn't use context.Background() because it will give warning in Go version 1.7 or greater. So, we used context.TODO() to avoid warning.

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection(colName) // create collection instance. it returns collection. we will use this collection instance to perform CRUD operations on "watchlist" collection.

	fmt.Println("Collection instance created!")
}

// insert a record in mongodb
func insertOneMovie(movie model.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
	// return insertResult, nil
}

// update a record in mongodb
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // convert string id to mongodb hex id
	filter := bson.M{"_id": id} // create filter. it is used to find the record in mongodb
	update := bson.M{"$set": bson.M{"watched": true}} // create update. it is used to update the record in mongodb
	updateResult, err := collection.UpdateOne(context.Background(), filter, update) // update the record in mongodb

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// delete one record from mongodb
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // convert string id to mongodb hex id
	filter := bson.M{"_id": id} // create filter. it is used to find the record in mongodb
	deleteResult, err := collection.DeleteOne(context.Background(), filter) // delete the record from mongodb

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

// delete all records from mongodb
func deleteAllMovies() {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{}) // delete all records from mongodb
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

// find all records from mongodb
func findAllMovies() []primitive.M {
	var movies []primitive.M // create a slice of primitive.M. primitive.M is used to get the BSON document as a map[string]interface{}. It is similar to map[string]interface{}.
	cur, err := collection.Find(context.Background(), bson.D{}) // find all records from mongodb. it returns mongo.Cursor and error. mongo.Cursor is an interface type that is used to iterate through all the records. bson.D{} specifies all the documents in the collection.

	if err != nil {
		log.Fatal(err)
	}

	// iterate through all the records
	for cur.Next(context.Background()) {
		var movie bson.M // create a variable to hold single record
		err := cur.Decode(&movie) // decode the record and store it in movie variable. cur.Decode() decodes the current record into the value pointed to by movie. It returns error if the record cannot be decoded. It returns io.EOF if there are no more records in the cursor.
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie) // append the record in movies slice
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background()) // close the cursor
	return movies // return the slice
}




// Actual controller functions
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded") // x-www-form-urlencoded is used to submit the form data in the request body. It is the default value of Content-Type header. json is also used to submit the form data in the request body. It is used when we want to submit the form data in json format.
	allMovies := findAllMovies() // find all records from mongodb
	json.NewEncoder(w).Encode(allMovies) // encode allMovies in json format and send it in response
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	params := mux.Vars(r) // get params from request. params is a map of string keys and string values. It returns a map of route variables that are set for the current request. It returns empty map if there are no variables.
	movie := findAllMovies() // find all records from mongodb

	// iterate through all the records. find the record with the given id and send it in response. it is a map of string keys and interface values.
	for _, item := range movie {
		if item["_id"] == params["id"] {
			json.NewEncoder(w).Encode(item) // encode item in json format and send it in response
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Netflix{}) // encode empty Netflix struct in json format and send it in response
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Mehtods", "POST") // Allow-Control-Allow-Mehtods is used to allow the methods. It is used when we want to send the request from one domain to another domain.
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie) // decode the request body and store it in movie variable. r.Body is an io.ReadCloser. It is the request body. It implements io.Reader and io.Closer interfaces. io.Reader is used to read the request body. io.Closer is used to close the request body.
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Mehtods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Mehtods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Mehtods", "DELETE")
	deleteAllMovies()
	json.NewEncoder(w).Encode("All Movies Deleted")
}