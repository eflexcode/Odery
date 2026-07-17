package entity

type Payload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Intent  string `json:"intent"`
	ImgUrl  string `json:"imgUrl"`
}
// add  Notification type
type Notification struct {
	Id        string  `json:"id"`
	UserId    string  `json:"user_id"`
	Payload   Payload `json:"payload"`
	CreatedAt string  `json:"created_at"`
}

type NotificationResult struct {
	Notifications []Notification `json:"notifications"`
	Total         int64          `json:"total"`
	Page          int64          `json:"page"`
	Limit         int64          `json:"limit"`
}

type StandardResponse struct{
	Message string  `json:"message"`
	StatusCode int `json:"status_code"`
}
