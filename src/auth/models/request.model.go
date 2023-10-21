package models

type SignInWithGoogleModles struct {
	IDToken  string `json:"idToken"`
	Position string `json:"position"`
}

type Payload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
	Position string `json:"position"`
}
