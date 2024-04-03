package main

import (
	"encoding/json"
	"errors"
)

type ID int
type Record interface {
	json.Marshaler
	json.Unmarshaler
}

// для типа type RawMessage []byte в пакете encoding/json
// декларированы оба метода выше,
// то есть *json.RawMessage реализует интерфейс Record

/* хорошо */
type Storage2 interface {
	Insert(Record, ID)                // метод принимает интерфейс
	Get(ID) (*json.RawMessage, error) // метод возвращает конкретный тип
}

type MapStore map[ID]*json.RawMessage

func (m MapStore) Insert(r Record, id ID) {
	// делаем приведение типов,
	// освобождая от этого вызывающего
	m[id] = r.(*json.RawMessage)
}
func (m MapStore) Get(id ID) (*json.RawMessage, error) {
	r, ok := m[id]
	// проверяем, есть ли запись в хранилище
	if !ok {
		return r, errors.New("not found")
	}
	return r, nil
}

// конструктор
func NewMapStore() MapStore {
	s := make(MapStore)
	return s
}
