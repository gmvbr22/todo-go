package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * Faz a ligação com o banco de dados
 * 
 * Tempo limite de 10 segundos
 */
func ConnectDB(url string, dbName string, collection string) (*mongo.Database, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	
	db := client.Database(dbName)
	return db, nil
}