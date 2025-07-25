package main

import (
	"cli-tracker/internal/commands"
	"cli-tracker/internal/storage"
)

func main() {

	st := storage.NewJsonStorage("some.json")

	handler := commands.NewHandler(st)

	handler.ParseFlags()

}
