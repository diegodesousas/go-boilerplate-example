package order

type Order struct {
	TotalInCents int    `json:"total_in_cents"`
	Buyer        Buyer  `json:"buyer"`
	Items        []Item `json:"items"`
}

type Buyer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	Description string `json:"description"`
	PriceInCent int    `json:"price_in_cent"`
}
