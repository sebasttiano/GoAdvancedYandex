package lesson1_5

import (
	"encoding/json"
	"log"
	"net/http"
)

// Use2r — основной объект для теста.
type User2 struct {
	ID        string
	FirstName string
	LastName  string
}

// UserViewHandler — хендлер, который нужно протестировать.
func UserViewHandler(users map[string]User2) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			http.Error(rw, "user_id is empty", http.StatusBadRequest)
			return
		}

		user, ok := users[userId]
		if !ok {
			http.Error(rw, "user not found", http.StatusNotFound)
			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(rw, "can't provide a json. internal error",
				http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonUser)
	}
}

func main4() {
	users := make(map[string]User2)
	u1 := User2{
		ID:        "u1",
		FirstName: "Misha",
		LastName:  "Popov",
	}
	u2 := User2{
		ID:        "u2",
		FirstName: "Sasha",
		LastName:  "Popov",
	}
	users["u1"] = u1
	users["u2"] = u2

	http.HandleFunc("/users", UserViewHandler(users))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
