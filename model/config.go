package model

type Config struct {
	AuthToken string  `json:"auth_token"`
	Addr      string  `json:"addr"`
	Wallet    string  `json:"wallet"`
	FilValue  float64 `json:"fil"`
}
