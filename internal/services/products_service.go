package services

import (
	"context"
	"errors"
	"gobid/internal/responses"
	"gobid/internal/store/pgstore"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductsService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewProductService(pool *pgxpool.Pool) ProductsService {
	return ProductsService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (ps *ProductsService) CreateProduct(
	ctx context.Context,
	sellerId uuid.UUID,
	productName string,
	description string,
	baseprice float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {
	id, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		SellerID:    sellerId,
		ProductName: productName,
		Description: description,
		Baseprice:   baseprice,
		AuctionEnd:  auctionEnd,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (ps *ProductsService) ReadProductsBySellerId(ctx context.Context, sellerId uuid.UUID) ([]responses.ProductResponse, error) {
	productsModel, err := ps.queries.GetProductsBySellerId(ctx, sellerId)

	if err != nil {
		return []responses.ProductResponse{}, err
	}

	productsResponse := make([]responses.ProductResponse, len(productsModel))
	for i, p := range productsModel {
		productsResponse[i] = responses.ProductResponse{
			ID:          p.ID,
			SellerID:    p.SellerID,
			ProductName: p.ProductName,
			Description: p.Description,
			Baseprice:   p.Baseprice,
			AuctionEnd:  p.AuctionEnd,
			IsSold:      p.IsSold,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		}
	}

	return productsResponse, nil
}

var ErrProductNotFound = errors.New("product not found")

func (ps *ProductsService) GetProductById(ctx context.Context, productId uuid.UUID) (pgstore.Product, error) {
	product, err := ps.queries.GetProductById(ctx, productId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Product{}, ErrProductNotFound
		}

		return pgstore.Product{}, err
	}

	return product, nil
}
