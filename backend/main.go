package main

import (
	"encoding/json"
	"example_jwt/jt"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var Users []User

func main() {
	Users = append(Users, User{
		Username: "Ivan",
		Password: "Check",
		Email:    "IvanChek)",
	})

	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(c.Handler)

	r.Post("/login", LoginHandler)
	r.Get("/protected", ProtectedHandler)
	r.Post("/auth", Auth)
	//r.Post()
	http.ListenAndServe(":8088", r)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userValid(u.Username, u.Password) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Пользователь с такой учеткой существует")
		return
	}
	Users = append(Users, u)
	tokenString, err := jt.CreateToken(u.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	err := jt.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Invalid token err: ", err)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprint(w, "Welcome to the protected area")
}

func userValid(username string, password string) bool {
	for _, u := range Users {
		if u.Username == username {
			if u.Password == password {
				return true
			}
		}
	}
	return false
}

type Response struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Пришел запрос")
	w.Header().Set("Content-Type", "application/json")

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("Значение запроса пользователя %v", u)
	if userValid(u.Username, u.Password) {
		fmt.Println("Юзер существует и валиден")
		tokenString, err := jt.CreateToken(u.Username)
		if err != nil {
			fmt.Println("Ошибка в поиске пользователя: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Token: tokenString})
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Неверные учетные данные")
	}
}
