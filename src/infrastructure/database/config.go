package database

type Config struct {
	DSN string `env:"CIDER_SERVER_DSN,default=root:cider@tcp4(mysql:3306)/cider"`
}