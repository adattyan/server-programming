package main

import (
	"kadai/models"
	"log"
)

func main() {
	//追加する構造体
	p1 := models.Contact{
		Id:     1,
		Name:   "adachi",
		Number: "123-456-7890",
		Email:  "aa@email",
	}

	//連絡先追加
	models.AddContact(p1)

	p2 := models.Contact{
		Id:     2,
		Name:   "tanaka",
		Number: "100-000-0890",
		Email:  "tanaka@email",
	}
	models.AddContact(p2)

	//リスト表示
	err := models.GetList()
	if err != nil {
		log.Println("Error : ", err)
	}

	err = models.GetContact(1)
	if err != nil {
		log.Println("Error : ", err)
	}
}
