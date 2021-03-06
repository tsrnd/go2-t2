package config

import "os"

//SetEnv set up env of db
func SetEnv() {
	os.Setenv("DB_CONNECTION", "mysql")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_DATABASE", "blog_golang")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "1")
	os.Setenv("DB_CHARSET", "utf8")
	os.Setenv("DB_PARSETIME", "True")
	os.Setenv("SSLMODE", "disable")
	os.Setenv("SERVER_PORT", ":6969")

}
