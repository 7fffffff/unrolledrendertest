package main

import (
	"fmt"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	renderer := render.New(render.Options{
		Layout: "layout",
	})
	http.HandleFunc("/example1", func(w http.ResponseWriter, req *http.Request) {
		recorder := httptest.NewRecorder()
		renderer.HTML(recorder, http.StatusOK, "example1", nil)
		if recorder.Body.String() != "example1" {
			panic(fmt.Errorf("expected example1, got %s", recorder.Body.String()))
		}
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/example2", func(w http.ResponseWriter, req *http.Request) {
		recorder := httptest.NewRecorder()
		renderer.HTML(recorder, http.StatusOK, "example2", nil)
		if recorder.Body.String() != "example2" {
			panic(fmt.Errorf("expected example2, got %s", recorder.Body.String()))
		}
		w.Write([]byte("OK"))
	})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
