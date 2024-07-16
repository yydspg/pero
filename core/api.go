package core

import (
	"github.com/rs/zerolog"
	"os"
)

type API struct {
	zerolog.Logger
	serviceDB *ServiceDB
	itemDB    *ItemDB
}

func NewAPI() *API {
	return &API{
		Logger:    zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}),
		serviceDB: newServiceDB(),
		itemDB:    newItemDB(),
	}
}
