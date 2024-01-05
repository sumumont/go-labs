package rabbit

import (
	"fmt"
	"github.com/apulisai/sdk/go-utils/broker"
	"github.com/apulisai/sdk/go-utils/broker/rabbitmq"
	"github.com/go-labs/internal/logging"
	"testing"
	"time"
)

func TestMQ(t *testing.T) {
	if err := startMQConnector(); err != nil {
		logging.Error(err).Send()
	}
	//defer stopMQConnector()
	for {
		time.Sleep(10 * time.Second)
		fmt.Println("wait msg")
	}
}

var mq_broker broker.Broker

func startMQConnector() error {

	addr := fmt.Sprintf("amqp://%s:%s@%s:%d", "admin", "aEsz6aN8cLEGp64A", "192.168.0.23", 30004)
	//@mark: no special orgnization rule now
	//if tenant.GetTenantId() > 1
	{
		addr += "/"
	}
	mq_broker = rabbitmq.NewBroker(
		broker.Addrs(addr),
		rabbitmq.ExchangeName("hysen"),
		rabbitmq.DurableExchange(),
	)
	if err := mq_broker.Connect(); err != nil {

		return err
	}

	//@mark: listen job monitor
	topic := fmt.Sprintf("%v", "topic_1")
	if _, err := mq_broker.Subscribe(topic,
		MonitorJobStatus, broker.Queue("test"), rabbitmq.DurableQueue(), broker.DisableAutoAck(), rabbitmq.AckOnSuccess()); err != nil {
		return err
	}

	return nil
}
func MonitorJobStatus(event broker.Event) error {
	logging.Debug().Interface("event", string(event.Message().Body)).Msg("get msg")
	return nil
}
func stopMQConnector() {
	if mq_broker != nil {
		mq_broker.Disconnect()
	}
}
