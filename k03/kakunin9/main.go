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
	//構造体初期化
	person := Person{Name: "adachi", Age: 21, Email: "aa@mail"}

	//objからJSON文字列（エンコード）に変換
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	//出力
	fmt.Println(string(jsonData))
}
