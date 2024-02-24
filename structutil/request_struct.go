package structutil

type RequestBody struct {
	Username     string `bson:"username"`
	Email        string `bson:"email"`
	ReferralCode string `bson:"referral_code"`
	Password     string `bson:"password"`
}
