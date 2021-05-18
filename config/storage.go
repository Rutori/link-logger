package config

import "encoding/json"

var Storage struct {
	Port          int    `json:"port"`
	TLSPort       int    `json:"tls_port"`
	Host          string `json:"host"`
	AdminPassword string `json:"admin_password"`
	DBName        string `json:"db_name"`
	CertPath      string `json:"cert"`
	KeyPath       string `json:"key"`
}

func Load(file []byte) error {
	return json.Unmarshal(file, &Storage)
}
