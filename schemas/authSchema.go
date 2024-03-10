package schemas

type LoginSchema struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}

type SignUpSchema struct {
	LoginSchema
	Email string `json:"email" binding:"required" form:"email"`
}
