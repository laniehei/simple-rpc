package main

import (
	"log"
	"net/rpc"
	"os"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatalf("connection error: %s", err)
		os.Exit(1)
	}

	a := Item{"first", "first item"}
	b := Item{"second", "second item"}
	c := Item{"third", "third item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)

}
