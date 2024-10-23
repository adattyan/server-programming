package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 連絡先構造体
type Contact struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Email  string `json:"email"`
}

// 連絡先リスト
var contacts []Contact

// 連絡先の追加
func AddContact(contact Contact) {
	contacts = append(contacts, contact)
}

// 連絡先の表示
func GetList() error {
	data, err := json.MarshalIndent(contacts, "", " ")
	if err != nil {
		return err
	}
	//リスト表示
	fmt.Println("↓↓　連絡先一覧　↓↓")
	fmt.Println(string(data))
	fmt.Println("↑↑　ここまで　↑↑")

	return nil
}

// 指定されたIDの連絡先を表示
func GetContact(id int) error {
	for _, contact := range contacts {
		if contact.Id == id {
			//ID発見
			data, err := json.Marshal(contact)
			if err != nil {
				return err
			}
			fmt.Printf("ID : %d を取得\n", id)
			fmt.Println(string(data))
			return nil
		}
	}

	//見つからなかった
	return errors.New("指定されたIDは見つかりませんでした")
}
