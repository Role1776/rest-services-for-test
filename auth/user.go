package app

type User struct {
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binding:"required"`
}

type UserTable struct {
	UserId      int    `json:"-"`
	ListId      int    `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
