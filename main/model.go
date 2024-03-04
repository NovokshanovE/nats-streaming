package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

type delivery struct {
	Delivery_id string `json:"order_uid"`
}

/*
TODO: Необходимо допилить структуры
*/
type Order struct {
	Order_uid    string     `json:"order_uid"`
	Track_number string     `json:"track_number"`
	Entry        string     `json:"entry"`
	Delivery     []delivery `json:"delivery"`
}

type Cache struct {
	Orders map[int]Order
}

// var body []byte
var cache = Cache{
	Orders: make(map[int]Order),
}

func (c Cache) from_json(json_str string) {
	order := Order{}

	if err := json.Unmarshal([]byte(json_str), &order); err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру

	c.Orders[len(c.Orders)] = order
	log.Print(cache.Orders[len(cache.Orders)-1])

}

func (c Cache) from_db(id int, json_str string) {
	order := Order{}

	if err := json.Unmarshal([]byte(json_str), &order); err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру

	c.Orders[id] = order
	log.Print(cache.Orders[len(cache.Orders)-1])
}

func (c Cache) by_id(id int) Order {

	return c.Orders[id]
}

func to_json() string {
	return ""
}

func start_program() {

}

func load_from_db() {
	db, err := sql.Open("postgres", "user=postgres password=2307 dbname=nats_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from message")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var data string

		if err := rows.Scan(&id, &data); err != nil {
			log.Fatal(err)
		}

		fmt.Print(reflect.TypeOf(data))
		fmt.Print(data)
		cache.from_db(id, data)
		// cache[id] = data
	}
}

// req := Order{
// 	Order_uid:    "wcubje",
// 	Track_number: "hjejce",
// 	Entry:        "checjw",
// 	Delivery: []delivery{{
// 		Delivery_id: "7867697cdcd",
// 	}},
// }
// if err := json.Decoder(r.Body).Decode(&req); err != nil {
// 	// обработка ошибки
// }
// encoder := json.NewEncoder(io.Discard) // Создание энкодера, вывод в /dev/null
// err = encoder.Encode(req)
// fmt.Print(err)
// _, err = db.Exec("select * from message")
// body, err = json.Marshal(req)
// var json_str string = ""
// for i := 0; i < len(body); i++ {
// 	json_str += string(body[i])
// 	// fmt.Print(string(body[i]))
// }
