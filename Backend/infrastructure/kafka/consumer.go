package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"log"
)

// ConsumerClient the client to receive message to kafka
type ConsumerClient struct {
	client    sarama.Consumer // the client of a sarama kafka connection
	connected bool            // the connected status
}

// Connect try to connect kafka cluster/node with configuration
func (c *ConsumerClient) Connect(cfg ClusterConfig) (err error) {
	c.client, err = sarama.NewConsumer(cfg.ServerList, nil)
	if err == nil {
		c.connected = true
	}

	return
}

// Listen listens kafka messages and press it to MessageChanel
func (c ConsumerClient) Listen(topic string, callBack func(string)) (err error) {
	if !c.connected {
		err = errors.New("client does not connected")
		log.Println(err)
		return
	}

	partitionList, err := c.client.Partitions(topic)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, partition := range partitionList {
		partitionConsumer, pErr := c.client.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if pErr != nil {
			return pErr
		} else {
			defer partitionConsumer.AsyncClose()
			go func(consumer sarama.PartitionConsumer) {
				for msg := range consumer.Messages() {
					callBack(string(msg.Value))
				}
			}(partitionConsumer)
		}
	}

	select {}

	return nil
}
