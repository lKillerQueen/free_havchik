package model

type WkusnoCupoons struct {
	Items []Items `json:"items"`
}
type Items struct {
	Name string `json:"name"`
	Pic  string `json:"picture"`
}
