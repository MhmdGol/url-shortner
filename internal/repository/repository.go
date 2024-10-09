package repository

import "url-shortner/internal/model"

type URLShortnerRepository interface {
	Put(url model.URLModel) error
	Get(urlHashStr string) (model.URLModel, error)
}
