package repository

import (
	"auth-services/model"
	"time"

	"github.com/go-resty/resty/v2"
)

type CartRepo interface {
	AddCart(model.RequestAddCart) (model.ResponseSuccessCart, *resty.Response, error)
	ListCart() ([]model.ResponseCart, *resty.Response, error)
	DetailCart(model.RequesCardById) (model.ResponseCart, *resty.Response, error)
	DeleteCart(model.RequesCardById) (model.ResponseCart, *resty.Response, error)
}

type cartRepo struct {
	client *resty.Client
}

// AddCart implements CartRepo.
func (r cartRepo) AddCart(req model.RequestAddCart) (model.ResponseSuccessCart, *resty.Response, error) {
	panic("unimplemented")
}

// DeleteCart implements CartRepo.
func (r cartRepo) DeleteCart(req model.RequesCardById) (model.ResponseCart, *resty.Response, error) {
	panic("unimplemented")
}

// DetailCart implements CartRepo.
func (r cartRepo) DetailCart(req model.RequesCardById) (model.ResponseCart, *resty.Response, error) {
	panic("unimplemented")
}

// ListCart implements CartRepo.
func (r cartRepo) ListCart() ([]model.ResponseCart, *resty.Response, error) {
	panic("unimplemented")
}

func NewCartRepo() CartRepo {
	return cartRepo{
		resty.New().SetTimeout(30 * time.Second).SetRetryCount(3),
	}
}
