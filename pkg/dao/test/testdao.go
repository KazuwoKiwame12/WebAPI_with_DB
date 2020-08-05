package testdao

import (
	"log"

	"github.com/KazuwoKiwame12/WebAPI_with_DB/pkg/db"
)

//Shogi ...型の構造体
//あとでjson形式にするので、jsonのタグをあらかじめつけておきます
type Shogi struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//FetchIndex ...shogi型のデータを返す
func FetchIndex() []Shogi {
	db := db.Connect()
	defer db.Close()

	//rowを取得
	rows, err := db.Query("SELECT * FROM shogi")
	if err != nil {
		log.Fatal(err.Error())
	}
	//Opening型のスライスに格納します
	shogiArgs := make([]Shogi, 0)
	for rows.Next() {
		var shogi Shogi
		err = rows.Scan(&shogi.ID, &shogi.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		shogiArgs = append(shogiArgs, shogi)
	}
	return shogiArgs
}
