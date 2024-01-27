package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"strconv"
	"strings"
	"time"
)

type Video struct {
	Id          string    // video_id
	Title       string    // title
	PublishTime time.Time // publish_time
	Tags        []string  // tags
	Views       int       // views
}

func readVideoCSV(ctx context.Context, db *sql.DB, csvFile string) error {
	// открываем csv файл
	file, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer file.Close()
	videos := make([]Video, 0, 1000)

	// определим индексы нужных полей
	const (
		Id          = 0 // video_id
		Title       = 2 // title
		PublishTime = 5 // publish_time
		Tags        = 6 // tags
		Views       = 7 // views
	)

	// конструируем Reader из пакета encoding/csv
	// он умеет читать строки csv-файла
	r := csv.NewReader(file)
	// пропустим первую строку с именами полей
	if _, err := r.Read(); err != nil {
		return err
	}

	for {
		// csv.Reader за одну операцию Read() считывает одну csv-запись
		// в виде []string
		l, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// инициализируем целевую структуру,
		// в которую будем делать разбор csv-записи
		v := Video{
			Id:    l[Id],
			Title: l[Title],
		}
		// парсинг строковых записей в типизированные поля структуры
		if v.PublishTime, err = time.Parse(time.RFC3339, l[PublishTime]); err != nil {
			return err
		}
		tags := strings.Split(l[Tags], "|")
		for i, v := range tags {
			tags[i] = strings.Trim(v, `"`)
		}
		v.Tags = tags
		if v.Views, err = strconv.Atoi(l[Views]); err != nil {
			return err
		}
		// добавляем полученную структуру в слайс
		videos = append(videos, v)
		if len(videos) == 1000 {
			if err = insertVideos(ctx, db, videos); err != nil {
				return err
			}
			videos = videos[:0]
		}
	}
	// добавляем оставшиеся записи
	return insertVideos(ctx, db, videos)
}

func insertVideos(ctx context.Context, db *sql.DB, videos []Video) error {
	// начинаем транзакцию

	tx, err := db.Begin()
	if err != nil {
		return err

	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		"INSERT INTO videos (video_id, title, publish_time, tags, views)"+
			" VALUES(?,?,?,?,?)")
	// пример с использованием PrepareContext
	// не забывайте закрывать запрос
	defer stmt.Close()
	for _, v := range videos {
		// все изменения записываются в транзакцию
		_, err := stmt.ExecContext(ctx, v.Id, v.Title, v.PublishTime,
			strings.Join(v.Tags, `|`), v.Views)
		if err != nil {
			// если ошибка, то откатываем изменения
			return err
		}
	}
	// завершаем транзакцию
	return tx.Commit()
}

func main() {
	db, err := sql.Open("sqlite", "newvideo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	_, err = db.ExecContext(ctx, `CREATE TABLE videos (
        "video_id" TEXT,
        "title" TEXT,
        "publish_time" TEXT,
        "tags" TEXT,
        "views" INTEGER
      )`)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	err = readVideoCSV(ctx, db, "USvideos.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))

	if err != nil {
		log.Fatal(err)
	}
}
