package transactiondto

type RequestTransaction struct {
	UserID     int    `json:"user_id"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Limit      int    `json:"limit"`
	StatusUser string `json:"status_user"`
}

type ResponseTransaction struct {
	UserID     int    `json:"user_id"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Limit      int    `json:"limit"`
	StatusUser string `json:"status_user"`
}
