package connection 

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"log"
	"fmt"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/santos95/go-etl/config"
)

func GetMongoConnection(uri string) *mongo.Client {

	// define the context
	ctx := context.TODO()

	// establish connection to mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {

		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {

		log.Fatal(err)
	} else {

		log.Println("Connected to Mongo")
	}

	return client
}

func GetSqlServerConnectionString(server string, dbname string, user string, password string) string {

	// response
	var connStr string 

	//decode password 
	decPass := config.DecodePassString(password)

	connStr = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, user, decPass, dbname)

	return connStr
}

func GetSqlServerConnection(connStr string) *sql.DB {

	// open the database connection
	db, err := sql.Open("sqlserver", connStr)

	if err != nil {

		fmt.Println(connStr)
		log.Fatalf("Failed to open the database connection: %v", err)
	} else {

		fmt.Println("Database Connection opened!")
	}

	// defer db.Close()

	// test connection
	err = db.Ping() 
	if err != nil {

		log.Fatalf("Failed to connecto to the database: %v", err)
	}

	fmt.Println("Connected to sql server.")

	return db 
}
