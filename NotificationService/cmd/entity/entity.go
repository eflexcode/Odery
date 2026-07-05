package entity

type Payload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Intent  string `json:"intent"`
	ImgUrl  string `json:"imgUrl"`
}

type Notification struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Total     int64  `json:"total"`
	CreatedAt string `json:"created_at"`
}
