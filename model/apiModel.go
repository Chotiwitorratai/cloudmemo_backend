package model

// SignUp struct to describe register a new user.
type SignUp struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	Username string `json:"username" validate:"required,lte=25"`
	Image *string `json:"image"`
}

// SignIn struct to describe login user.
type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

type UserUpdateRequest struct {
		Email    string `json: "email" validate:"email"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
}

type CreateMemo struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"desc"`
		Body        string `json:"body"`
		Weather     string `json:"weather"`
		MusicUrl    *string `json:"music_url"`
}

type UpdateMemo struct {
	ID          uint   `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"desc"`
	Body        string `json:"body"`
	Weather     string `json:"weather"`
	MusicUrl    *string `json:"music_url"`
}
type Renew struct {
	RefreshToken string `json:"refresh_token"`
}

type PublishMemo struct {
	ID uint `json:"id" validate:"required"`
}