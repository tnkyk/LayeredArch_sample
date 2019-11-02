package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/tnkyk/LayeredArch_sample/usecase"
)

type TodoHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
	GetOneTodo(http.ResponseWriter, *http.Request, httprouter.Params)
	UpsertTodo(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteTodo(http.ResponseWriter, *http.Request, httprouter.Params)
}

//この構造体は元々TodoUseCaseinterfaceと紐づいていて、Indexメソッドの宣言の際にこの構造体と新たに紐づけられる
type todoHandler struct {
	todoUseCase usecase.TodoUseCase
}

// NewTodoUseCase : Todo データに関する Handler を生成
func NewTodokHandler(tu usecase.TodoUseCase) TodoHandler {
	return &todoHandler{
		todoUseCase: tu,
	}
}

//Index: Get /todos -> todoデータ一覧取得  ポインタにしないと、レシーバの値を影響させることが出来ない
func (th *todoHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	type TodoField struct {
		Id    string `json:"id"`
		Title string `json:"title"`
		//Author    string    `json:"author"`
		CreatedAt time.Time `json:"created_at"`
	}
	//response : Todo　API　のレスポンス
	type response struct {
		Todos []TodoField `json:"todos"`
	}

	ctx := r.Context()

	//ユースケースの呼び出し
	todos, err := th.todoUseCase.TodoGetAll(ctx)
	if err != nil {
		http.Error(w, "Internal Sever Error", 500)
		return
	}

	//取得したドメインモデルをresponseに変換
	res := new(response)
	for _, todo := range todos {
		var tf TodoField
		tf.Id = todo.Id
		tf.Title = todo.Title
		//tf.Author = todo.Author
		tf.CreatedAt = todo.CreatedAt
		res.Todos = append(res.Todos, tf)
	}

	//クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (th *todoHandler) GetOneTodo(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	//request：TodoAPIのパラメータ
	type TodoField struct {
		Id    string `json:"id"`
		Title string `json:"title"`
		//Author    string    `json:"author"`
		CreatedAt time.Time `json:"created_at"`
	}
	//response : Todo　API　のレスポンス
	type response struct {
		Todos []TodoField `json:"todos"`
	}

	ctx := r.Context()
	param := r.URL.Query()
	id := param.Get("id")
	log.Println(id)

	//ユースケースの呼び出し
	todo, err := th.todoUseCase.TodoGetById(ctx, id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Sever Error", 500)
		return
	}

	//取得したドメインモデルをresponseに変換
	res := new(response)

	var tf TodoField
	tf.Id = todo.Id
	tf.Title = todo.Title
	//tf.Author = todo.Author
	tf.CreatedAt = todo.CreatedAt
	res.Todos = append(res.Todos, tf)

	//クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (th *todoHandler) UpsertTodo(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	ctx := r.Context()
	type NewTodo struct {
		ID       string    `json:"id"`
		Title    string    `json:"title"`
		CreateAt time.Time `json:"create_at"`
	}
	var newTodo NewTodo

	newTodo.ID = u.String()
	newTodo.Title = r.FormValue("title")
	newTodo.CreateAt = time.Now()
	err = th.todoUseCase.UpsertTodo(ctx, u.String(), r.FormValue("title"), time.Now())

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(newTodo)

}

func (th *todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	ctx := r.Context()
	err := th.todoUseCase.DeleteTodo(ctx, r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
}
