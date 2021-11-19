package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	err = json.Unmarshal(bs, &lp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Back into data structure", lp2)
}
