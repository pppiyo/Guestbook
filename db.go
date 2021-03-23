package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connection = Connection()

func Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func Save(note *Note) error {
	note.ID = primitive.NewObjectID()
	note.Time = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	notesCollection := connection.Database("guestbook").Collection("notes")
	_, err := notesCollection.InsertOne(ctx, note)
	return err
}

func getList() (notes []Note) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	notesCollection := connection.Database("guestbook").Collection("notes")

	cur, err := notesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		result := Note{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, result)
	}

	return
}

func Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	notesCollection := connection.Database("guestbook").Collection("notes")
	_, err := notesCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
