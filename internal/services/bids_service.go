package services

import (
	"context"
	"errors"
	"gobid/internal/store/pgstore"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BidsService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewBidsService(pool *pgxpool.Pool) BidsService {
	return BidsService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var ErrBidIsTooLow = errors.New("the bid value is too low")

func (bs *BidsService) PlaceBid(
	ctx context.Context,
	product_id uuid.UUID,
	bidder_id uuid.UUID,
	amaount float64,
) (pgstore.Bid, error) {
	product, err := bs.queries.GetProductById(ctx, product_id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Bid{}, err
		}
	}

	highestBid, err := bs.queries.GetHighestBidByProductId(ctx, product_id)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Bid{}, err
		}
	}

	if product.Baseprice >= amaount || highestBid.BidAmount >= amaount {
		return pgstore.Bid{}, ErrBidIsTooLow
	}

	highestBid, err = bs.queries.CreateBid(ctx, pgstore.CreateBidParams{
		ProductID: product_id,
		BidderID:  bidder_id,
		BidAmount: amaount,
	})

	if err != nil {
		return pgstore.Bid{}, err
	}

	return highestBid, nil
}
