package order

import (
	"context"

	proto "order-service/proto/productproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductClient struct {
	client proto.ProductServiceClient
}

func NewProductClient(address string) (*ProductClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewProductServiceClient(conn)
	return &ProductClient{client: client}, nil
}

func (pc *ProductClient) GetProductPrice(productID string) (float32, error) {
	req := &proto.GetProductRequest{Id: productID}
	res, err := pc.client.GetProduct(context.Background(), req)
	if err != nil {
		return 0, err
	}
	return res.Price, nil
}
