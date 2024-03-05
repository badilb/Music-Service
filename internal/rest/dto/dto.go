package dto

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateForm struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type PlaylistForm struct {
	Title string `json:"title"`
}
