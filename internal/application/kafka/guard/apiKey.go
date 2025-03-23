package guard

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	API_KEY = os.Getenv("API_KEY")
)

const (
	HEADER_API_KEY_KEY         = "x-api-key"
	ERR_API_KEY_NOT_SET        = "[Critical issue] API_KEY is not set"
	ERR_API_KEY_INVALID        = "API_KEY is invalid"
	ERR_API_KEY_HEADER_MISSING = "API_KEY header is not set"
)

// ApiKeyMiddleware checks if the API key in the message headers is valid
func ApiKey(msg *kafka.Message) error {
	if API_KEY == "" {
		log.Fatal(ERR_API_KEY_NOT_SET)
	}
	headers := msg.Headers

	for _, header := range headers {
		if header.Key == HEADER_API_KEY_KEY {
			if bytes.Equal(header.Value, []byte(API_KEY)) {
				return nil
			} else {
				return errors.New(ERR_API_KEY_INVALID)
			}
		}
	}

	return errors.New(ERR_API_KEY_HEADER_MISSING)
}
