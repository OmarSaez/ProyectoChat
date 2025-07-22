package dto

type LoginRequest struct {
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
}
