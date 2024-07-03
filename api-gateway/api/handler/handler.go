package handler

import (
	user "api-gateway/proto/user"
	prod "api-gateway/proto/product"
)

type Handler struct{
	user  user.UserServiceClient
	product prod.ProductServiceClient
}

func NewHandler(user user.UserServiceClient,product prod.ProductServiceClient) *Handler {
	return &Handler{user: user,product: product}
}
