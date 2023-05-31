package consts

import "os"

var (
	APIPort      = os.Getenv("API_PORT")
	Env          = os.Getenv("GO_ENV")
	DateFormat24 = "2006-01-02 15:04:05"
)

func IsTest() bool {
	return Env == "test"
}
