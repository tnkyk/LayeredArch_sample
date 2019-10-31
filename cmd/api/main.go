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

	//ルーティングの設定
	router := httprouter.New()
	router.GET("/api/todos", todoHandler.Index)

	//サーバー起動
	port := ":3000" //"3000"だとエラーになる
	fmt.Println(`Server Start >> http:// localhost:%d`, port)
	log.Fatal(http.ListenAndServe(port, router))
}
