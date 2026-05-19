package main

import (
	"log"

	"github.com/kudo07/go-grpc/module1/proto"
)

func main() {
	person := proto.Person{
		Name: "abc",
	}

	log.Println(person.GetName())
}