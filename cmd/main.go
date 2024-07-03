package main

import (
	"fmt"
	"gostarter/internal/server"
)

func main() {
	s := server.New()

	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
