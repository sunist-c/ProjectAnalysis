package kafka

import (
	"ProjectAnalysis/infrastructure/common"
	"errors"
	"github.com/Shopify/sarama"
	"strconv"
)

// ProducerClient the client to produce message to kafka
type ProducerClient struct {
	client    sarama.SyncProducer // the client of a sarama kafka connection
	connected bool                // the connected status
}

// Connect try to connect kafka cluster/node with configuration
func (c *ProducerClient) Connect(cfg ClusterConfig) (err error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = cfg.toPartitioner()
	config.Producer.RequiredAcks = cfg.toRequireAck()
	config.Producer.Timeout = cfg.toTimeout()
	config.Net.SASL.Enable = cfg.toEnableSasl()
	config.Net.SASL.User = cfg.Username
	config.Net.SASL.Password = cfg.Password
	config.Producer.Return.Successes = true

	c.client, err = sarama.NewSyncProducer(cfg.ServerList, config)
	if err == nil {
		c.connected = true
	} else {
		c.connected = false
	}

	return
}

// Send sends message to kafka with topic
func (c ProducerClient) Send(topic, message string) (uuid string, err error) {
	uuid, err = "", nil
	if !c.connected {
		err = errors.New("client does not connected")
		return
	}

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(message)

	partition, offset, err := c.client.SendMessage(msg)
	if err == nil {
		uuid = common.GenerateMd5Len16(strconv.Itoa(int(partition)), strconv.FormatInt(offset, 10))
	}

	return
}
