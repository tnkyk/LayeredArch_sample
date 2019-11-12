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

	//サーバー起動
	todoPersistence := persistence.NewTodoPersistence()
	todoUseCase := usecase.NewTodoUseCase(todoPersistence)
	todoHandler := handler.NewTodokHandler(todoUseCase)

	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	//ルーティングの設定
	Router := httprouter.New()
	Router.GET("/api/todos", todoHandler.Index)
	Router.GET("/api/todo", todoHandler.GetOneTodo)
	Router.PUT("/api/todo/upsert", todoHandler.UpsertTodo)
	Router.DELETE("/api/todo/delete", todoHandler.DeleteTodo)

	Router.GET("/api/users", userHandler.Index)
	Router.GET("/api/user", todoHandler.GetOneTodo)
	Router.PUT("/api/user/upsert", userHandler.UpsertUser)
	Router.DELETE("/api/user/delete", userHandler.DeleteUser)
	//"3000"だとエラーになる
	port := ":3000"
	fmt.Println(`Server Start >> http:// localhost:%d`, port)
	log.Fatal(http.ListenAndServe(port, Router))
}
