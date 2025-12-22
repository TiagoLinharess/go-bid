package product

import (
	"context"
	"gobid/internal/validator"
	"time"

	"github.com/google/uuid"
)

type CreateProductReq struct {
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

const minAuctionDuration = 2 * time.Hour

func (req CreateProductReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.ProductName), "product_name", "this field cannot be blank")
	eval.CheckField(validator.NotBlank(req.Description), "product_description", "this field cannot be blank")
	eval.CheckField(validator.MinChars(req.Description, 10) && validator.MaxChars(req.Description, 255), "product_description", "this field must have a length between 10 and 255")
	eval.CheckField(req.Baseprice > 0, "baseprice", "this field must be greater than 0")
	eval.CheckField(req.AuctionEnd.Sub(time.Now()) >= minAuctionDuration, "auction_end", "must be at least 2 hours duration")

	return eval
}
