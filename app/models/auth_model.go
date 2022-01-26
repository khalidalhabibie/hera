package models

// SignUp struct to describe register a new user.
type SignUp struct {
	Password   string `json:"password" validate:"required,lte=255" `
	Repassword string `json:"re_password" validate:"required,lte=255"`
	Username   string `json:"username" validate:"required,lte=255" groups:"author"`
}

// SignIn struct to describe login user.
type SignIn struct {
	Username string `json:"username" validate:"required,lte=255" groups:"author"`
	Password string `json:"password" validate:"required,lte=255"`
}

const (
	AuthAuthor = "author"
	AuthPublic = "public"
)
