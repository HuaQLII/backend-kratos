package data

import (
	"verify-code/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// TODO new Redis client
	data := &Data{}

	options, err := redis.PareURL("redis://localhost:6379/1?dial_timeout=1")
	if err != nil {
		data.Rdb = nil

	}
	rdb := redis.NewClient(options)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
