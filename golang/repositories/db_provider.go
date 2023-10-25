package repositories

import (
	"context"
	"db.sampes.puc/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var dbCache *gorm.DB

func getDbInstance() *gorm.DB {
	if dbCache == nil {
		//dsn := "sqlserver://amos-junior.korp:123qwe@dev-env-mssql.com?database=amos-junior_BASE_ERP"
		dsn := "root:1234@tcp(127.0.0.1:3306)/teste?charset=utf8mb4&parseTime=True&loc=Local"
		dbCache, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		autoMigrate()
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return dbCache.WithContext(ctx)
}

func autoMigrate() {
	dbCache.AutoMigrate(&entities.Robozinho{})

	/*dbCache.Exec(`
		CREATE TABLE Robozinho (
			id INT PRIMARY KEY IDENTITY (1,1) NOT NULL,
			nome VARCHAR (MAX) NOT NULL,
			peso DECIMAL(19, 6) NOT NULL,
			data_cadastro DATETIME2 NOT NULL,
			data_atualizacao DATETIME2 NULL,
		)
	`)*/
}
