package persistence

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
)

func TestGetAll(t *testing.T) {
	tp1, _ := time.Parse("2006年01月02日 15時04分05秒", "2019年11月02日 16時05分00秒")
	tp2, _ := time.Parse("2006年01月02日 15時04分05秒", "2019年11月02日 15時37分20秒")
	cases := []struct {
		Id        string
		Title     string
		CreatedAt time.Time
	}{
		{"4f214757-1ec1-49d7-8465-cac3212ae169", "example3", tp1},
		{"e51cd7d4-335d-4aa8-9cd1-85ea46569492", "example2", tp2},
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:3000/api/todos"), nil)
	ctx := req.Context()
	for _, cs := range cases {
		todos, err := NewTodoPersistence().GetAll(ctx)
		t.Log(cs.CreatedAt)
		if err != nil {
			t.Fatalf(`want get todos`)
		}
		for key, todo := range todos {
			if model.Todo(cs) == todo {
				break
			} else if key < len(todos) {
				continue
			} else {
				t.Fatalf(`wantId:%s,wantTitle:%s,wantCreatedAt:%s
				but getId:%s,getTitle:%s,getCreatedAt:%s`, cs.Id, cs.Title, cs.CreatedAt, todo.Id, todo.Title, todo.CreatedAt)
			}
		}
	}
}

func TestGetById(t *testing.T) {
	tp1, _ := time.Parse("2006年01月02日 15時04分05秒", "2019年11月02日 16時05分00秒")
	tp2, _ := time.Parse("2006年01月02日 15時04分05秒", "2019年11月02日 15時37分20秒")
	cases := []struct {
		Id        string
		Title     string
		CreatedAt time.Time
	}{
		{"4f214757-1ec1-49d7-8465-cac3212ae169", "example3", tp1},
		{"e51cd7d4-335d-4aa8-9cd1-85ea46569492", "example2", tp2},
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:3000/api/todo"), nil)
	ctx := req.Context()
	for _, cs := range cases {
		todo, err := NewTodoPersistence().GetById(ctx, cs.Id)
		if err != nil {
			t.Fatalf(`want get todo`)
		}
		if todo.Title != cs.Title {
			t.Fatalf("wantTitle:%s,getTitle:%s", cs.Title, todo.Title)
		}
		if todo.CreatedAt != cs.CreatedAt {
			t.Fatalf("wantCreatedAt:%s,getCreatedAt:%s", cs.CreatedAt, todo.CreatedAt)
		}
	}
}

func TestUpsertTodo(t *testing.T) {
	// tp1, _ := time.Parse("2006年01月02日 15時04分05秒", "2019年11月02日 16時05分00秒")
	// tp2, _ := time.Parse("2006年01月02日 15時04分05秒", "2019年11月02日 15時37分20秒")
	// cases := []struct {
	// 	CurrentId        string
	// 	CurrentTitle     string
	// 	CurrentCreatedAt time.Time
	// }{
	// 	{"4f214757-1ec1-49d7-8465-cac3212ae169", "example3", tp1},
	// 	{"e51cd7d4-335d-4aa8-9cd1-85ea46569492", "example2", tp2},
	// }
	// req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:3000/api/upsert"))
	// ctx := req.Context()

}
