package main

import (
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
    "mini-project-sanbercode/controllers"
    "mini-project-sanbercode/database"

    _ "github.com/lib/pq"
)

var (
    DB  *sql.DB
    err error
)

func main() {
    // ENV Configuration
    err = godotenv.Load("config/.env")
    if err != nil {
        log.Fatalf("failed to load environment file: %v", err)
    }

    dbName := os.Getenv("DB_NAME")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    psqlInfo := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
        dbHost, dbPort, dbUser, dbPassword, dbName,
    )

    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("failed to open database: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("DB Connection Failed: %v", err)
    } else {
        fmt.Println("DB Connection Success")
    }

    database.DbMigrate(DB)
    defer DB.Close()

    // Router GIN
    router := gin.Default()
    router.GET("/persons", controllers.GetAllPerson)
    router.POST("/persons", controllers.InsertPerson)
    router.PUT("/persons/:id", controllers.UpdatePerson)
    router.DELETE("/persons/:id", controllers.DeletePerson)

    router.Run("localhost:8080")
}
