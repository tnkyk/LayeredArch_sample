package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tnkyk/LayeredArch_sample/usecase"
)

type TodoHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
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

//Index: Get /todos -> todoデータ一覧取得
func (th todoHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	//request：TodoAPIのパラメータ
	type requset struct {
		Begin uint `query:begin`
		Limit uint `query:limit`
	}

	type TodoField struct {
		Id        int64     `json:"id"`
		Title     string    `json:"title"`
		Author    string    `json:"author"`
		CreatedAt time.Time `json:"created_at"`
	}
	//response : Todo　API　のレスポンス
	type response struct {
		Todos []TodoField `json:"todos"`
	}

	ctx := r.Context()

	//ユースケースの呼び出し
	todos, err := th.todoUseCase.GetAll(ctx)
	if err != nil {
		http.Error(w, "Internal Sever Error", 500)
		return
	}

	//取得したドメインモデルをresponseに変換
	res := new(response)
	for _, todo := range todos {
		var tf TodoField
		tf.Id = int64(todo.Id)
		tf.Title = todo.Title
		tf.Author = todo.Author
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
