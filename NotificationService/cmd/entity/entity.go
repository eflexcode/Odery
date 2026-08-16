package entity

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Payload struct {
	Title   string `bson:"title"`
	Message string `bson:"message"`
	Intent  string `bson:"intent"`
	ImgUrl  string `bson:"imgurl"`
}

// add  Notification type
type Notification struct {
	Id        string  `bson:"id"`
	UserId    string  `bson:"userid"`
	Payload   Payload `bson:"payload"`
	Type      string  `bson:"type"` //order,payment
	CreatedAt string  `bson:"createdat"`
}

type NotificationDB struct {
	UserId    string  `json:"user_id"`
	Payload   Payload `json:"payload"`
	Type      string  `json:"type"` //order,payment
	CreatedAt string  `json:"created_at"`
}

type NotificationResult struct {
	Notifications []Notification `json:"notifications"`
	Total         int64          `json:"total"`
	Page          int64          `json:"page"`
	Limit         int64          `json:"limit"`
}

type Order struct {
	Id              string `json:"id"`
	ProductId       string `json:"product_id"`
	ProductName     string `json:"product_name"`
	ProductCurrency string `json:"product_currency"`
	ProductImgUrl   string `json:"product_img_url"`
	Count           int64  `json:"count"`
	Amount          int64  `json:"amount"`
	UserId          string `json:"user_id"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type StandardResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type Payment struct {
	Id          string  `json:"id"`
	UserId      string  `json:"user_id"`
	CardId      string  `json:"card_id"`
	Amount      float64 `json:"amount"`
	ProductId   string  `json:"product_id"`
	OrderId     string  `json:"order_id"`
	Status      string  `json:"status"` //done, processing, submitted, failed
	Reason      string  `json:"reason"` //eg: insufficient funds, card error, network error,order canceled
	Type        string  `json:"type"`   //refund,paid,-
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

