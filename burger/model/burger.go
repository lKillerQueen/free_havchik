package model

type BurgerCupoons struct {
	Response Response `json:"response"`
}

type Response struct {
	Dishes []Dishe `json:"dishes"`
}

type Dishe struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Image       string `json:"image"`
}
