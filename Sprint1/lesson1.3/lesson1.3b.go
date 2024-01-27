package main

import (
	"io"
	"net/http"
)

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин <input type="text" name="login"></label>
            <label>Пароль <input type="password" name="password"></label>
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Auth(login, password string) bool {
	return login == `guest` && password == `demo`
}

func mainPageAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Добро пожаловать!")
		} else {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		}
		return
	} else {
		io.WriteString(w, form)
	}
}

func main() {
	err := http.ListenAndServe(`:8080`, http.HandlerFunc(mainPageAuth))
	if err != nil {
		panic(err)
	}
}
