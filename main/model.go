package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/lib/pq"
)

type delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type payment struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	Payment_dt    int    `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}

type item struct {
	Chrt_id      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         string `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}

/*
TODO: Необходимо допилить структуры
*/
type Order struct {
	Order_uid    string   `json:"order_uid"`
	Track_number string   `json:"track_number"`
	Entry        string   `json:"entry"`
	Delivery     delivery `json:"delivery"`
	Payment      payment  `json:"payment"`
	Items        []item   `json:"items"`
}

type Cache struct {
	Orders map[string]Order
}

// var body []byte
var cache = Cache{
	Orders: make(map[string]Order),
}

func (c Cache) from_json(json_str string) {
	order := Order{}

	if err := json.Unmarshal([]byte(json_str), &order); err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру

	c.Orders[order.Order_uid] = order
	// log.Print(cache.Orders[len(cache.Orders)-1])

}

func (c Cache) from_db(id string, json_str string) {
	order := Order{}

	if err := json.Unmarshal([]byte(json_str), &order); err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру

	c.Orders[id] = order
	log.Print(cache.Orders["b563feb7b2b84b6test"])
}

func (c Cache) by_id(id string) Order {
	// if c.Orders[id] != nil {

	// }
	return c.Orders[id]
}

func (c Cache) to_json(id string) string {
	bytes, err := json.MarshalIndent(c.by_id(id), "", "\t")
	if err != nil {
		// log.Panic()
		// fmt.Errorf()
		panic(err)
	} // Десериализация JSON в структуру
	// order := ""
	// for i := 0; i < len(order_bytes); i++ {
	// 	order += order_bytes[i]
	// }
	return string(bytes)
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
		var id string
		var data string

		if err := rows.Scan(&id, &data); err != nil {
			log.Fatal(err)
		}

		// fmt.Print(reflect.TypeOf(data))
		// fmt.Print(data)
		cache.from_db(id, data)

		// cache[id] = data
	}
	log.Print(cache)
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
