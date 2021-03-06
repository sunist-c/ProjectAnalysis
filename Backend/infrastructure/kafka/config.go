package kafka

import (
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

// TypePartitioner the enums of Kafka partitioner implement
type TypePartitioner string

const (
	RandomPartitioner        TypePartitioner = "RandomPartitioner"
	HashPartitioner          TypePartitioner = "HashPartitioner"
	ManualPartitioner        TypePartitioner = "ManualPartitioner"
	ReferenceHashPartitioner TypePartitioner = "ReferenceHashPartitioner"
	RoundRobinPartitioner    TypePartitioner = "RoundRobinPartitioner"
)

func (t TypePartitioner) toString() string {
	return string(t)
}

// TypeRequireAck the enums of Kafka require_ack options
type TypeRequireAck string

const (
	NoResponse   TypeRequireAck = "NoResponse"
	WaitForLocal TypeRequireAck = "WaitForLocal"
	WaitForAll   TypeRequireAck = "WaitForAll"
)

func (t TypeRequireAck) toString() string {
	return string(t)
}

// ClusterConfig the configuration structure of kafka cluster
type ClusterConfig struct {
	ServerList  []string `json:"server_list" yaml:"server_list"`
	RequireAck  string   `json:"require_ack" yaml:"require_ack"`
	Timeout     string   `json:"timeout" yaml:"timeout"`
	Partitioner string   `json:"partitioner" yaml:"partitioner"`
	EnableSasl  string   `json:"enable_sasl" yaml:"enable_sasl"`
	Username    string   `json:"username" yaml:"username"`
	Password    string   `json:"password" yaml:"password"`
}

// toPartitioner exchange the config field to partition implement
func (c ClusterConfig) toPartitioner() func(topic string) sarama.Partitioner {
	switch c.Partitioner {
	case RandomPartitioner.toString():
		return sarama.NewRandomPartitioner
	case HashPartitioner.toString():
		return sarama.NewHashPartitioner
	case ManualPartitioner.toString():
		return sarama.NewManualPartitioner
	case ReferenceHashPartitioner.toString():
		return sarama.NewReferenceHashPartitioner
	case RoundRobinPartitioner.toString():
		return sarama.NewRandomPartitioner
	default:
		return sarama.NewRandomPartitioner
	}
}

// toRequireAck exchange the config field to require_ack option
func (c ClusterConfig) toRequireAck() sarama.RequiredAcks {
	switch c.RequireAck {
	case NoResponse.toString():
		return sarama.NoResponse
	case WaitForLocal.toString():
		return sarama.WaitForLocal
	case WaitForAll.toString():
		return sarama.WaitForAll
	default:
		return sarama.NoResponse
	}
}

// toTimeout exchange the config field to kafka admin timeout
func (c ClusterConfig) toTimeout() time.Duration {
	timeout, err := strconv.Atoi(c.Timeout)
	if err != nil {
		timeout = 60
	}

	return time.Second * time.Duration(timeout)
}

// toEnableSasl exchange the config field to enable_sasl option
func (c ClusterConfig) toEnableSasl() bool {
	if c.EnableSasl == "true" || c.EnableSasl == "TRUE" {
		return true
	} else {
		return false
	}
}
