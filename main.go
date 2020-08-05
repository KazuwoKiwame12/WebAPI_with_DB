package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	testdao "github.com/KazuwoKiwame12/WebAPI_with_DB/pkg/dao/test"
)

func main() {
	r := mux.NewRouter()
	// localhost:8080/shogi/で将棋の駒を確認することが出来る
	r.HandleFunc("/shogi/", showShogiIndex)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func showShogiIndex(w http.ResponseWriter, r *http.Request) {
	shogi := testdao.FetchIndex()
	//json形式に変換します
	bytes, err := json.Marshal(shogi)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(string(bytes)))
}
