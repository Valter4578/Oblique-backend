package model

// Category ...
type Category struct {
	Title     string `json:"title,omitempty" bson:"title,omitempty"`
	ImageName string `json:"imageName,omitempty" bson:"imageName,omitempty"`
	Color     string `json:"color,omitempty" bson:"color,omitempty"`
}
