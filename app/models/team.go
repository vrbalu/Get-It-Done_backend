package models

type Team struct {
	Name    string `bson:"name" json:"name"`
	Members []User `bson:"members" json:"members,omitempty"`
	Owner   User   `bson:"owner" json:"owner,omitempty"`
}
