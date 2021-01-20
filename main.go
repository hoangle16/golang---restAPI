package main

import (
	"log"
	"net/http"
	"restful_api/database"
	"restful_api/packages/user"
	delivery "restful_api/packages/user/handler/http"
)

func main() {
	router := http.NewServeMux()
	sqliteDB := database.GetSqliteConnectionPool()

	userRepo := user.NewSqliteUserRepository(sqliteDB)
	userService := user.NewService(userRepo)

	delivery.MakeUserHandler(router, userService)

	log.Fatal(http.ListenAndServe(":6969", router))
}
