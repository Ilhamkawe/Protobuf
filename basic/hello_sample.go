package basic

import (
	"log"
	"protobuf/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Muhammad Ilham Kusumawardhana",
	}

	log.Println(&h)
}
