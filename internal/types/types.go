package types

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required"`
	Class string `json:"class" validate:"required"`
	Email string `json:"email" validate:"required"`
}
