package authdto

type RequestRegister struct {
	Email    string `json:"email" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
}

type RequestLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
	Token    string `json:"token"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
	Address  string `json:"address"`
}
