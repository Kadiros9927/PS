package auth

type LoginRequest struct { // LoginRequest represents the request payload for login
	Email    string `json:"email" validate:"required,email"` // Email is the user's email address
	Password string `json:"password" validate:"required"`    // Password is the user's password
}

type RegisterRequest struct { // LoginRequest represents the request payload for login
	Email    string `json:"email" validate:"required,email"` // Email is the user's email address
	Username string `json:"username" validate:"required"`    // Username is the user's username
	Password string `json:"password" validate:"required"`    // Password is the user's password
}

type LoginPayload struct { // LoginPayload represents the payload for login requests
	Token string `json:"token"`
}

type RegisterPayload struct { // LoginPayload represents the payload for login requests
	Token string `json:"token"`
}
