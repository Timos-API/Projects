package projects

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ProjectID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title         string             `json:"title" bson:"title" validate:"required"`
	Icon          string             `json:"icon" bson:"icon" validate:"required"`
	Description   string             `json:"description" bson:"description" validate:"required"`
	Hero          string             `json:"hero" bson:"hero" validate:"required,url"`
	DisplayOnHome bool               `json:"displayOnHome" bson:"displayOnHome" validate:"required"`
	Thumbnail     string             `json:"thumbnail" bson:"thumbnail" validate:"required,url"`
	Website       string             `json:"website" bson:"website" validate:"required,url"`
	Visible       bool               `json:"visible,omitempty" bson:"visible,omitempty"`
	GitHub        string             `json:"github" bson:"github" validate:"required,url"`
	DesignTools   []string           `json:"designTools" bson:"designTools"`
	Development   []string           `json:"development" bson:"development"`
	Frameworks    []string           `json:"frameworks" bson:"frameworks"`
	Sections      []ProjectSection   `json:"sections" bson:"sections"`
}

type ProjectSection struct {
	Title    string               `json:"title" bson:"title" validate:"required"`
	Subtitle string               `json:"subtitle" bson:"subtitle"`
	Items    []ProjectSectionItem `json:"items" bson:"items"`
}

type ProjectSectionItem struct {
	Title       string `json:"title" bson:"title" validate:"required"`
	Description string `json:"description" bson:"description"`
	Background  string `json:"background" bson:"background"`
	Asset       string `json:"asset" bson:"asset"`
	Size        string `json:"size" bson:"size"`
}

type Exception struct {
	Message string `json:"message,omitempty"`
}
