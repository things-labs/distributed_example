package services

import (
	context "context"
)

type ProdService struct{}

func (*ProdService) GetProductStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 100,
	}, nil
}

func (sf *ProdService) GetProdStocks(ctx context.Context, in *QuerySize) (*ProdResponseList, error) {
	return &ProdResponseList{
		Prodres: []*ProdResponse{{ProdStock: 100}, {ProdStock: 101}, {ProdStock: 102}},
	}, nil
}
