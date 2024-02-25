package controller

import (
	"fmt"
	"short_url/model"
	"short_url/model/short_url"
	"sync"
)

var (
	b *BaseController
	m sync.Once
)

type BaseController struct {
	shortUrl short_url.Service
}

func GetBaseController() (*BaseController, error) {
	if b == nil {
		return nil, fmt.Errorf("[BaseController] no init")
	}
	return b, nil
}

func Init() {
	model.Init()
	var err error
	m.Do(func() {
		b = &BaseController{}
		if b.shortUrl, err = short_url.GetServer(); err != nil {
			panic(err)
		}
	})
}