package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/tnkyk/LayeredArch_sample/usecase"
)

type UserHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
	GetOneUser(http.ResponseWriter, *http.Request, httprouter.Params)
	UpsertUser(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteUser(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

func (uh *userHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	type UserField struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//response : Todo　API　のレスポンス
	type response struct {
		Users []UserField `json:"users"`
	}
	ctx := r.Context()

	users, err := uh.userUsecase.UserGetAll(ctx)
	if err != nil {
		log.Println(err)
	}
	res := new(response)
	for _, user := range users {
		var uf UserField = UserField(user)
		res.Users = append(res.Users, uf)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func (uh *userHandler) GetOneUser(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	type UserField struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	ctx := r.Context()
	param := r.URL.Query()
	name := param.Get("name")

	user, err := uh.userUsecase.UserGetByName(ctx, name)
	if err != nil {
		log.Println(err)
		return
	}
	var res UserField = UserField(*user)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func (uh *userHandler) UpsertUser(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	ctx := r.Context()
	type NewUser struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var newUser NewUser

	newUser.ID = u.String()
	newUser.Name = r.FormValue("name")
	newUser.Email = r.FormValue("email")
	newUser.Password = r.FormValue("password")
	err = uh.userUsecase.UpsertUser(ctx, u.String(), r.FormValue("name"), r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(newUser)
}

func (uh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	ctx := r.Context()
	err := uh.userUsecase.DeleteUser(ctx, r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
}
