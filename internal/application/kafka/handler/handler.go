package handler

import (
	"mailer_ms/internal/application/kafka/guard"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

type Handler struct {
	SAY_HELLO string
}

var (
	MESSAGE_KIND = &Handler{
		SAY_HELLO: "hello",
	}
)

func HandleMessage(logger *zap.Logger, msg *kafka.Message) {
	if err := guard.ApiKey(msg); err != nil {
		return
	}

	// switch on msg.key and send the request to the correct handler
	if string(msg.Key) == MESSAGE_KIND.SAY_HELLO {
		// handle msg
	}
	// else, do nothing
	logger.Sugar().Infof("[HANDLER] A message has not been handle: %v", string(msg.Value))
}
