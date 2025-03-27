package pokeapi

import (
	"net/http"
	"time"

	"github.com/snansidansi/pokedex-cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheCheckIntervall, maxCacheAge time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheCheckIntervall, maxCacheAge),
	}
}
