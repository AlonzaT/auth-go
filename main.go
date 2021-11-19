package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "tom",
	}

	p2 := person{
		First: "lisa",
	}

	lp := []person{p1, p2}

	bs, err := json.Marshal(lp)
	if err != nil {
		//Panic is for programming errors
		log.Panic(err)
	}

	fmt.Println(string(bs))

	lp2 := []person{}

	//Unmarshall takes data and a pointer to where
	//the data will be unmarshalled
	err = json.Unmarshal(bs, &lp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Back into data structure", lp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8888", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p3 := person{
		First: "jenny",
	}

	err := json.NewEncoder(w).Encode(p3)
	if err != nil {
		log.Println("Got bad data!", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {

}
