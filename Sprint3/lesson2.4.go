package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "videos.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRowContext(context.Background(),
		"SELECT title, likes, comments_disabled "+
			"FROM videos ORDER BY likes DESC LIMIT 1")
	var (
		title  string
		likes  int
		comdis bool
	)
	// порядок переменных должен соответствовать порядку колонок в запросе
	err = row.Scan(&title, &likes, &comdis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s | %d | %t \r\n", title, likes, comdis)
}
