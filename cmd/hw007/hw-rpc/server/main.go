package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"

	db "github.com/eastonfly/gocourse/cmd/hw007/hw-rpc/db"
)

type Users struct {
	Name string
}

type Handler struct{}

func (rh *Handler) Register(user *Users, reply *int) error {
	if user.Name == " " {
		return errors.New("unable to create empty user")
	}
	if db.Check(user.Name) == nil {
		*reply = db.Create(user.Name)
		return nil
	} else {
		return errors.New("user already exists")
	}
}

func (rh *Handler) List(user *Users, reply *[]string) error {

	*reply = db.Base
	return nil

}

func main() {

	handler := new(Handler)
	rpc.Register(handler)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
