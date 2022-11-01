package config

type Settings struct {
	BotToken      string `json:"botToken"`
	BurgerCommand string `json:"burgerCommand"`
	WkusnoCommand string `json:"wkusnoCommand"`
	HelloText     string `json:"helloText"`
	ApiWkusno     string `json:"apiWkusno"`
	ApiBurger     string `json:"apiBurger"`
}
