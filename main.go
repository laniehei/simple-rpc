package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type Item struct {
	Title string
	Body  string
}

type API struct {
}

var database []Item

func main() {
	var err error
	var api API

	if err = rpc.Register(
		&api,
	); nil != err {
		log.Fatalf("error while registering API: %s", err)
		os.Exit(1)
	}

	port := ":4040"

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening: %s", err)
		os.Exit(1)
	}

	log.Printf("serving rpc on port %s", port)
	if err = http.Serve(listener, nil); nil != err {
		log.Fatalf("error serving: %s", err)
	}

}

func (a *API) GetByName(item Item, reply *Item) error {

	for _, val := range database {
		if val.Title == item.Title {
			reply = &val
		}
	}

	return nil
}

func (a *API) CreateItem(item Item, reply *Item) error {
	database = append(database, item)

	reply = &item

	return nil

}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)

	spew.Dump(database)

	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = edit
			changed = edit
		}
	}

	if changed.Title == "" {
		return fmt.Errorf("item could not be found and edited: %s", edit.Title)
	}

	reply = &changed

	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			reply = &item
			break
		}
	}

	if reply.Title == "" {
		return fmt.Errorf("couldn't delete %s", item.Title)

	}

	return nil
}
