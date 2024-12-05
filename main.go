package main

import (
    "flag"
    "fmt"
    "strconv"
    "log"
    "github.com/santos95/go-etl/config"
    "github.com/santos95/go-etl/connection"
    "time"
    "context"
)

func main() {

    // define a flag - name: type, default: "incremental", Description
    processType := flag.String("type", "incremental", "Type of process (Incremental | Initial)")

    // parse the flag
    flag.Parse()

    // initialize metadata
    const layout = "2006-01-02 15:04:05"
    var dataProcessed int64 = 0 
    var dataErased int64 = 0
    var batch int
    start := time.Now().Format(layout)
    fmt.Println("Execution Start: ", start)
    
    // get configuration values 
    configValues, err := config.GetConfigValues()
    if err != nil {

        log.Fatalf("Error getting config values: %v", err)
    }

    // origin configuration
    sqlServer := configValues.Server
    fmt.Println("SQLServer URI: : ", sqlServer)
    database := configValues.Database
    fmt.Println("Origin Database: ", database)
    user := configValues.User
    fmt.Println("User:", user)
    batchSize, _ := strconv.Atoi(configValues.Batchsize)
    fmt.Println("batchsize: ", batchSize)

    // target configuration
    fmt.Println("Target Configuration: ")
    mongouri := configValues.MongoURI
    fmt.Println("MongoURI:", mongouri)
    mongoDatabase := configValues.DatabaseMongo
    fmt.Println("MongoDatabase: ", mongoDatabase)
    collection := configValues.Collection 
    fmt.Println("Collection: ", collection)

    // establish mongodb connection 
    client := connection.GetMongoConnection(mongouri)
    ctx := context.TODO()
    defer client.Disconnect(ctx)
    
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Error pinging MONGODB: %v", err)
    } else {
        fmt.Println("Successfully connected to mongodb")
    }

    if *processType == "incremental" {
        batch = 1
    } else {
        batch = 0
    }

    fmt.Println("This are the values ", layout, " dataProcessed: " , dataProcessed, " DataErased: ", dataErased, " Batch: ", batch)

    

}