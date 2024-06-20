package config

import (
	"context"
	"fmt"
	"log"
	"os"


	"go.mongodb.org/mongo-driver/mongo"
	
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var Instance MongoInstance

func Dbconect() error {
	uri := os.Getenv("URI")
	if uri == "" {
		return fmt.Errorf("debes configurar la variable de entorno 'MONGODB_URI'")
	}

	var err error
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("error al conectar con MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("error al hacer ping a MongoDB: %v", err)
	} else {
		fmt.Println("Canectado a la base de datos")
	}

	var db = client.Database("Cbsanantero")

	Instance = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}

func DesconectarDB() {
	if Instance.Client != nil {
		if err := Instance.Client.Disconnect(context.Background()); err != nil {
			log.Printf("Error al desconectar la base de datos: %v", err)
		} else {
			log.Println("Desconexión exitosa de MongoDB")
		}
	}
}
