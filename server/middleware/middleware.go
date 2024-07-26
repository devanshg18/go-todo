package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection

func init() {
    loadTheEnv()
    createDBInstance()
}
func loadTheEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading the environment file")
    }
}
func createDBInstance() {
    connectionString := os.Getenv("DB_URI")
    dbName := os.Getenv("DB_NAME")
    collectionName := os.Getenv("DB_COLLECTION_NAME")
    clientOptions := options.Client().ApplyURI(connectionString)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")

    collection = client.Database(dbName).Collection(collectionName)
    fmt.Println("Collection Instance created")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/x-www-form")
    w.Header().Set()("Access-Control-Allow-Origin" , "*")
    payload :=GetAllTasks()
    json.NewEncoder(w).Encode(payload)
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin" , "*")
    w.Header().Set("Access-Control-Allow0Methods" , "POsT")
    w.Header().Set("Access-Control-Allow-Header" , "Contrnt-Type")
    var task models.ToDoList
    json.NewDecoder(r.Body).Decode(&task)
    insertOneTask(task)
    json.NewEncoder(w).Encode(task)
}
func TaskComplete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin" , "*")
    w.Header().Set("Access-Control-Allow0Methods" , "POsT")
    w.Header().Set("Access-Control-Allow-Header" , "Contrnt-Type")

    params:= mux.Var(r)
    TaskComplete(params["id"])
    json.NewEncoder(w).Encode(params["id"])
}
func UndoTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Contrnt-Type" ,"application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin" , "*")
    w.Header().Set("Access-Control-Allow-Methods" , "PUT")
    w.Header().Set("Access-Control-Allow-Headers" , "Content-Type")
    params := mux.Vars(r)
    UndoTask(params["id"])
    json.NewEncoder(w).Encode(params["id"])

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Contrnt-Type" ,"application/x-www-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods" , "DELETE")
    w.Header().Set("Access-Control-Allow-Headers" , "Content-Type")
    params := mux.Vars(r)
    deleteOneTask(params["id"])
}
func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type" ,)
    w.Header().Set("Access-Control-Allow-Origin")
    count:=deleteAllTasks()
    json.NewEncoder(w).Encode(count)
    
}
func GetAllTasks()[]primitive.M{
    cur , err := collection.Find(context.Background() , bson.D{{}})
    if err != nil {
    log.Fatal(err)
    }   
    var results []bson.M
    for cur.Next(context.Background()){ 
        var result bson.M
        e :=cur.Decode(&result)
        if e != nil {
            log.Fatal(e)
        }
        results = append(results , result)
    }   
    if err := cur.Err(); err != nil{
        log.Fatal(err)
    }
    cur.Close(context.Background())
    return results
}
func TaskComplete (task string){
    id , _ := primitive.ObjectIDFromHex(task)
    filter := bson.M{"_id":id}
    update := bson.M("$set":bson.M{"status":true})
    result , err :=collection.UpdateOne(context.Background() , filter , update)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Modified Count" , result.ModifiedCount)
}
func deleteOneTask(){
    id,_ := primitive.ObjectIDFromHex(task)
    filter := bson.M{"_id":id}
    d,err := collection.DeleteOne(context.Background(),filter)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Deleted Document" , d.)
}
func insertOneTask(task models.ToDoList){
    insertRes , err :=collection.insertOneTask(context.Background() , task )
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Inserted a single record" , insertRes.Inserted)
}
func UndoTask(task string){

}
func deleteAllTasks(){
    d, err:=collection.DeleteMany(context.Background() , bson.d)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Deleted document" , d.DeletedCount)
    return d.DeletedCount
}
