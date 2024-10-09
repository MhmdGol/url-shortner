package controller

import (
	"context"
	urlshortnerapiv1 "url-shortner/api/grpc/url-shortner/v1"
	"url-shortner/internal/service"
)

type URLShortnerController struct {
	srv service.URLShortnerService
	urlshortnerapiv1.UnimplementedURLShortnerServiceServer
}

var _ urlshortnerapiv1.URLShortnerServiceServer = (*URLShortnerController)(nil)

func NewURLShortnerController() *URLShortnerController {
	return &URLShortnerController{
		srv: service.URLShortnerService{},
	}
}

func (u *URLShortnerController) ShortURL(
	ctx context.Context,
	in *urlshortnerapiv1.ShortURLRequest,
) (*urlshortnerapiv1.ShortURLResponse, error) {
	shortURL := u.srv.ShortURL(ctx, in.GetUrl())

	return &urlshortnerapiv1.ShortURLResponse{
		ShortenUrl: shortURL,
	}, nil
}
