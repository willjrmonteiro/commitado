package database

import (
	"context"
	"log"
	"os"
	"time"

	"commitado/graphql/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetBill(id string) *model.Bill {
	billCollec := db.client.Database("commitment").Collection("bills")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var bill model.Bill
	err := billCollec.FindOne(ctx, filter).Decode(&bill)
	if err != nil {
		log.Fatal(err)
	}
	return &bill
}

func (db *DB) GetBills() []*model.Bill {
	billCollec := db.client.Database("commitment").Collection("bills")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var bills []*model.Bill
	cursor, err := billCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &bills); err != nil {
		panic(err)
	}

	return bills
}

func (db *DB) CreateBill(billInfo model.CreateBillInput) *model.Bill {
	billCollec := db.client.Database("commitment").Collection("bills")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := billCollec.InsertOne(ctx, bson.M{"name": billInfo.Name, "deadline": billInfo.Deadline, "status": billInfo.Status, "amount": billInfo.Amount})
	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnBill := model.Bill{ID: insertedID, Name: billInfo.Name, Deadline: billInfo.Deadline, Status: billInfo.Status, Amount: billInfo.Amount}
	return &returnBill
}

func (db *DB) UpdateBill(billId string, billInfo model.UpdateBillInput) *model.Bill {
	billCollec := db.client.Database("commitment").Collection("bills")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateBillInfo := bson.M{}

	if billInfo.Name != nil {
		updateBillInfo["name"] = billInfo.Name
	}
	if billInfo.Deadline != nil {
		updateBillInfo["deadline"] = billInfo.Deadline
	}
	if billInfo.Status != nil {
		updateBillInfo["status"] = billInfo.Status
	}
	if billInfo.Amount != nil {
		updateBillInfo["amount"] = billInfo.Amount
	}

	_id, _ := primitive.ObjectIDFromHex(billId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateBillInfo}

	results := billCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var bill model.Bill

	if err := results.Decode(&bill); err != nil {
		log.Fatal(err)
	}

	return &bill
}

func (db *DB) DeleteBill(billId string) *model.DeleteBillResponse {
	jobCollec := db.client.Database("commitment").Collection("bills")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(billId)
	filter := bson.M{"_id": _id}
	_, err := jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteBillResponse{DeletedBillID: billId}
}
