package services

import (
	"context"
	"log"

	. "github.com/Carjul/web-service-gin/config"
	"github.com/Carjul/web-service-gin/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func GetArticulos() []bson.M {
	collection := Instance.Database.Collection("Articulos")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var articulos []bson.M
	if err = cursor.All(context.Background(), &articulos); err != nil {
		log.Fatal(err)
	}
	return articulos
}

func GetArticulo(id string) bson.M {
	objID, err0 := primitive.ObjectIDFromHex(id)
	if err0 != nil {
		log.Fatal(err0)
	}
	collection := Instance.Database.Collection("Articulos")

	filter := bson.M{"_id": objID}
	var articulo bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&articulo)
	if err != nil {
		log.Fatal(err)
	}
	return articulo
}

func CreateArticulo(articulo models.Articulo) {
	collection := Instance.Database.Collection("Articulos")
	_, err := collection.InsertOne(context.Background(), articulo)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateArticulo(id string, articulo models.Articulo) {
	objID, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		log.Fatal(err2)
	}
	collection := Instance.Database.Collection("Articulos")
	filter := bson.M{"id": objID}
	update := bson.M{"$set": articulo}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

}

func DeleteArticulo(id string) {
	objID, err3 := primitive.ObjectIDFromHex(id)
	if err3 != nil {
		log.Fatal(err3)
	}
	collection := Instance.Database.Collection("Articulos")
	filter := bson.M{"id": objID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
