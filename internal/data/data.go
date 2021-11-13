package data

import (
	"liquid/internal/conf"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

type Data struct{}

func NewData(c *conf.Data) (*Data, error) {
	//load config...
	return &Data{}, nil
}
