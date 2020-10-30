package internal

import (
	"log"

	"github.com/levigross/grequests"
)

const STORE_API_URL = "https://apis.vinmonopolet.no/stores/v0/details"

type Store struct {
}

// GetAllStores returns a JSON payload with all VM stores
func (s *Store) GetAllStores() string {
	resp, err := grequests.Get(STORE_API_URL, nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	return resp.String()
}
