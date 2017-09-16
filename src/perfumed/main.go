package main

import (
	"server"
)

func main() {
	cnf := server.NewConfig()
	s, err := server.NewServer(cnf)

	if err != nil {
		panic(err)
	}

	s.Listen()
}
