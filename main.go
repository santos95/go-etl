package main

import (
    "flag"
    "fmt"
    "strconv"
    "log"
    "github.com/santos95/go-etl/config"
    "github.com/santos95/go-etl/connection"
    "github.com/santos95/go-etl/process"
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
    port, _ := strconv.Atoi(configValues.Port)
    fmt.Println("Port: ", port)
    user := configValues.User
    fmt.Println("User:", user)
    pass := configValues.Password 
    batchSize, _ := strconv.Atoi(configValues.Batchsize)
    fmt.Println("batchsize: ", batchSize)

    connStr := connection.GetSqlServerConnectionString(sqlServer, database, user, pass)

    fmt.Println("connection string: ", connStr)
    
    // get sql server connection
    db := connection.GetSqlServerConnection(connStr)

    // close db connection
    defer db.Close()

    // test connection
    var version string 
    err = db.QueryRow("SELECT @@VERSION").Scan(&version)

    if err != nil {

        log.Fatalf("Query failed: %v", err)
    }

    fmt.Println("SqlServer version: ", version)

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

    fmt.Println("Start ETL-PROCESS")
    fmt.Println("--------------------------------------------")
    fmt.Println("--------------------------------------------")

    // get endDate from the last execution
    startDate := process.GetLasExecution(db)
    fmt.Println("StartDate: ", startDate)

    if *processType == "incremental" {
        batch = 1
    } else {
        batch = 0
    }

    fmt.Println("This are the values ", layout, " dataProcessed: " , dataProcessed, " DataErased: ", dataErased, " Batch: ", batch)

    

}