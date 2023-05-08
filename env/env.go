package env

import "time"

const (
	SERVER_PORT             = "8081"
	EXTERNAL_API_TIMEOUT    = time.Millisecond * 200
	DATABASE_UPDATE_TIMEOUT = time.Millisecond * 10
	CLIENT_REQUEST_TIMEOUT  = time.Millisecond * 300
	LOG_FILE_NAME           = "cotacao.txt"
)
