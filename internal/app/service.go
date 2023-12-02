package app

import (
	"github.com/markLishansky/url-short/internal/store"
	desc "github.com/markLishansky/url-short/pkg"
)

type Service interface {
	desc.ShorterServer
}

type service struct {
	provider store.DataStore
}

func NewService(provider store.DataStore) (Service, error) {
	return &service{provider: provider}, nil
}
