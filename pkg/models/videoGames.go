package models

// file used to contain structure that will be used througout entire project

// Struct representing a new PostgreSQL object
// look up what bson stands for
type VideoGame struct {
	Title    string  `json:"title" bson:"title,omitempty"`
	Platform string  `json:"platform" bson:"platform,omitempty"`
	Genre    string  `json:"genre" bson:"genre,omitempty"`
	Price    float32 `json:"price" bson:"price,omitempty"` // look into the difference between float32 and 64
}
