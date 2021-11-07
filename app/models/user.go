package models

type User struct {
	Name      string `bson:"name" json:"name"`
	Key       string `bson:"key" json:"key"`
	Role      string `bson:"role" json:"role"`
	Email     string `bson:"email" json:"email"`
	CreatedAt string `bson:"createdAt" json:"createdAt,omitempty"`
}
