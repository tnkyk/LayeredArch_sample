package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	handler "github.com/tnkyk/LayeredArch_sample/handler/rest"
	"github.com/tnkyk/LayeredArch_sample/infra/persistence"
	"github.com/tnkyk/LayeredArch_sample/usecase"
)

func main() {
	todoPersistence := persistence.NewTodoPersistence()
	todoUseCase := usecase.NewTodoUseCase(todoPersistence)
	todoHandler := handler.NewTodokHandler(todoUseCase)

	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	//ルーティングの設定
	router := httprouter.New()
	router.GET("/api/todos", todoHandler.Index)
	router.GET("/api/todo", todoHandler.GetOneTodo)
	router.PUT("/api/todo/upsert", todoHandler.UpsertTodo)
	router.DELETE("/api/todo/delete", todoHandler.DeleteTodo)

	router.GET("/api/users", userHandler.Index)
	router.GET("/api/user", todoHandler.GetOneTodo)
	router.PUT("/api/user/upsert", userHandler.UpsertUser)
	router.DELETE("/api/user/delete", userHandler.DeleteUser)
	//サーバー起動
	port := ":3000" //"3000"だとエラーになる
	fmt.Println(`Server Start >> http:// localhost:%d`, port)
	log.Fatal(http.ListenAndServe(port, router))
}
