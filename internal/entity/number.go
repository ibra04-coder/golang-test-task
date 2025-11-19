package entity

type NumberInput struct {
	Number int `json:"number"`
}

type NumberResponse struct {
	Numbers []int `json:"numbers"`
}