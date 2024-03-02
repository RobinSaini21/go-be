package structutil

type RequestBody struct {
	Username     string `bson:"username"`
	Email        string `bson:"email"`
	ReferralCode string `bson:"referral_code"`
	Password     string `bson:"password"`
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}
