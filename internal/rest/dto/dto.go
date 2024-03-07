package dto

type LoginForm struct {
	Username string `json:"user"`
	Password string `json:"pwd"`
}

type UpdateForm struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type PlaylistForm struct {
	Title string `json:"title"`
}
