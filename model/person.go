package model

type Person struct {
	ID        int    `json:"id" form:"id" gorm:"primary_key"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
}

type CreatePerson struct {
	FirstName string `form:"firstName" binding:"required"`
	LastName  string `form:"lastName" binding:"required"`
}

type UpdatePerson struct {
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
}
