package commands

import (
	"cli-tracker/internal/storage"
	"flag"
	"fmt"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) ParseFlags() {
	var add string
	var update int
	var list string
	flag.StringVar(&add, "add", "", "Use an `add` to add task in a tracker")
	flag.IntVar(&update, "mark", 0, "Mark status that your track has")
	flag.StringVar(&list, "list", "", "Use a `list` with parameters to see a list of tracks (done, in progress)")
	flag.Parse()

	if add != "" {
		id, err := h.storage.Create(add)
		if err != nil {
			fmt.Printf("something wrong: %s\n", err.Error())
			return
		}

		fmt.Printf("✅ Task successfully created (ID:%d)\n", id)
	}
	if list != "" && update == 0 {
		switch list {
		case "done":
			tracks, err := h.storage.FindByStatus("done")
			if err != nil {
				fmt.Printf("something wrong: %s\n", err.Error())
				return
			}
			fmt.Println(tracks)
		case "in progress":
			tracks, err := h.storage.FindByStatus("in process")
			if err != nil {
				fmt.Printf("something wrong: %s\n", err.Error())
				return
			}
			fmt.Println(tracks)
		default:
			return

		}
	}

	if update == 0 {
		return
	} else {
		track, err := h.storage.Update(update, list)
		if err != nil {
			fmt.Printf("something wrong: %s\n", err.Error())
		}
		fmt.Printf("✅ Task updated: %v", track)
	}
}
