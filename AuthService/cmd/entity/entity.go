package entity

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Name      string `json:"Name"`
	ImgUrl    string `json:"img_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type StandardResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type TokenRes struct {
	Token string `json:"token"`
}
