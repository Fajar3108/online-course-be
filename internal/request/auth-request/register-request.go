package authrequest

import "mime/multipart"

type RegisterRequest struct {
	Name     string                `json:"name" validate:"required"`
	Email    string                `json:"email" validate:"required,email"`
	Password string                `json:"password" validate:"required"`
	Avatar   *multipart.FileHeader `form:"avatar" validate:"omitempty,image"` // Use FileHeader for file uploads
}
