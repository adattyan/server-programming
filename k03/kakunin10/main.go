package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"mail"`
}

func main() {
	//JOSN文字列の宣言
	jsonData := `{"name":"adachi", "age":21, "email":"a@mail"}`

	//デコード
	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatal(err)
	}

	//構造体の内容を出力
	fmt.Println(person)

}
