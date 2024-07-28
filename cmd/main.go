package main

import (
	"blooprint/internal/server"
	"fmt"
)

func main() {
	s := server.New()

	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
