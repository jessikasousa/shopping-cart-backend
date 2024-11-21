package db

import (
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DSN do banco de dados não foi encontrado no arquivo .env")
    }

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
    }

    DB = database
    log.Println("Conexão ao banco de dados bem-sucedida!")
}
