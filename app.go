package main

import (
	"embed"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rg-app/app/controller"
	"github.com/rg-app/app/handler"
	"github.com/rg-app/app/model"
	"github.com/rg-app/config"
)

//go:embed app/views/*
var Resources embed.FS

func main() {
	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.Student{})
	studentHandle := handler.NewStudentHandler(conn)

	mainAPI := controller.NewAPI(studentHandle, Resources)
	mainAPI.Start()
}

// flyctl secrets set DATABASE_URL=postgres://aditira:SemogaBerkah@sd-app-db.internal:5432/postgres
