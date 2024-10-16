package main

// 構造体の宣言
type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {

	//構造体の初期化その1
	p1 := Person{
		Name:  "adachi",
		Age:   21,
		Email: "aa@mail",
	}

	//構造体の初期化その2
	p2 := new(Person)
	p2.Name = "adachi"
	p2.Age = 21
	p2.Email = "aa@mail"

	//構造体の初期化その2
	//キー名を指定しない、順番で
	p3 := Person{
		"adachi", 21, "aa@mail",
	}

	//ポインタ
	p4 := &Person{
		Name:  "adachi",
		Age:   21,
		Email: "aa@mail",
	}
}
