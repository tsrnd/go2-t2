package config

import "os"

//SetEnv set up env of db
func SetEnv() {
	os.Setenv("DB_CONNECTION", "postgres")
<<<<<<< HEAD
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_DATABASE", "blog_golang")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "root")
=======
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_DATABASE", "blog_golang")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "1")
	os.Setenv("DB_CHARSET", "utf8")
	os.Setenv("DB_PARSETIME", "True")
>>>>>>> 40c543874ac6f0596aba2e52915a84c7c26f0114
	os.Setenv("SSLMODE", "disable")
	os.Setenv("SERVER_PORT", ":6969")
}
