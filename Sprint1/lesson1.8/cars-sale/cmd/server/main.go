package main

import (
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strings"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")
	if carID == "" {
		http.Error(rw, "carID param is missed", http.StatusBadRequest)
		return
	}
	rw.Write([]byte(carFunc(carID)))
}

func brandHandle(writer http.ResponseWriter, request *http.Request) {
	list := make([]string, 0)
	carBrand := strings.ToLower(chi.URLParam(request, "brand"))
	for _, c := range cars {
		bm := strings.Split(strings.ToLower(c), ` `)
		if bm[0] == carBrand {
			list = append(list, c)
		}
	}
	io.WriteString(writer, strings.Join(list, ", "))
}

func modelHandle(writer http.ResponseWriter, request *http.Request) {
	car := strings.ToLower(chi.URLParam(request, "brand") + ` ` + chi.URLParam(request, "model"))
	for _, c := range cars {
		if strings.ToLower(c) == car {
			io.WriteString(writer, c)
			return
		}
	}
	http.Error(writer, "unknown model: "+car, http.StatusNotFound)

}

func CarRouter() chi.Router {
	r := chi.NewRouter()

	r.Route("/cars", func(r chi.Router) {
		r.Get("/", carsHandle)
		r.Route("/{brand}", func(r chi.Router) {
			r.Get("/", brandHandle)
			r.Route("/{model}", func(r chi.Router) {
				r.Get("/", modelHandle)
			})
		})
	})
	r.Get("/car/{id}", carHandle)
	return r
}

func main() {

	log.Fatal(http.ListenAndServe(":8080", CarRouter()))
}
