package types

type Student struct {
	Id    int64    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   uint8    `json:"age" validate:"required"`
	Class string `json:"class" validate:"required"`
	Email string `json:"email" validate:"required"`
}
