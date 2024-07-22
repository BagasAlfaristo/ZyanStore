package main

import (
	"ZyanStore/app/configs"
	"ZyanStore/app/databases"
	"ZyanStore/app/migrations"
	"ZyanStore/app/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.InitMigrations(dbMysql)

	e := echo.New()

	routers.InitRoutuer(e, dbMysql)
	e.Logger.Fatal(e.Start(":8080"))
}
