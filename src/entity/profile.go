package entity

type Profile struct {
	Name   string `form:"name" bson:"name" binding:"required"`
	Email  string `form:"email" bson:"email" binding:"required, email"`
	Detail string `form:"detail" bson:"detail" binding:"required"`
}
