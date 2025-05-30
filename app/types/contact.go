package types

type ContactEmail struct {
	Name    string `form:"name" binding:"required,min=3"`
	Email   string `form:"email" binding:"required,email"`
	Message string `form:"message" binding:"required,min=10,max=400"`

	Errors map[string]string
}
