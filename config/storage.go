package config

import "encoding/json"

var Storage struct {
	Port          int    `json:"port"`
	AdminPassword string `json:"admin_password"`
	DBName        string `json:"db_name"`
}

func Load(file []byte) error {
	return json.Unmarshal(file, &Storage)
}
