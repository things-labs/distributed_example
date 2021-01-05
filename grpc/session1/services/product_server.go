package services

import context "context"

// user's defined
type ProdService struct{}

func (*ProdService) GetProductStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: req.ProdId,
	}, nil
}
