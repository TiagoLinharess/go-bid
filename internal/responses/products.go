package responses

import (
	"time"

	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
	IsSold      bool      `json:"is_sold"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
