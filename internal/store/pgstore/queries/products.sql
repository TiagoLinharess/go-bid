-- name: CreateProduct :one
INSERT INTO
    products (
        seller_id,
        product_name,
        description,
        baseprice,
        auction_end
    )
VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1;

-- name: GetProductsBySellerId :many
SELECT * FROM products
WHERE seller_id = $1
ORDER BY auction_end ASC;