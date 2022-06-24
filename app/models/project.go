package models

type Project struct {
	Id          string  `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string  `bson:"name" json:"name,omitempty"`
	Description string  `bson:"description,omitempty" json:"description,omitempty"`
	Members     []*User `bson:"members" json:"members,omitempty"`
	Owner       []*User `bson:"owner" json:"owner,omitempty"`
}
