package model

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func init() {
	Cache = cache.New(cache.NoExpiration, 5*time.Minute)
}
