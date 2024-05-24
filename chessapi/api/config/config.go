package config

import "github.com/vivilCODE/chess/chessapi/dbhandler"

type Config struct {
	DBHandler        *dbhandler.DBHandler
	GapiClientID     string
	GapiClientSecret string
}
