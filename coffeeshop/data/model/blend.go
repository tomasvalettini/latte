package datamodel

type Blend struct {
	Id    int    `json:"id"`
	Title string `json:"text"`
	Drips []Drip `json:"drips"`
}
