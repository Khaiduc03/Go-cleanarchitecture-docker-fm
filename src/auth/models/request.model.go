package models

type SignInWithGoogleModel struct {
	idToken string `json:"idToken"`
}

type Payload struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}
