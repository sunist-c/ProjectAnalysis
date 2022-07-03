package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

func TestClient_Connect(t *testing.T) {
	type fields struct {
		client    sarama.SyncProducer
		connected bool
	}
	type args struct {
		cfg ClusterConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_case_success",
			fields: fields{
				client:    nil,
				connected: false,
			},
			args: args{cfg: ClusterConfig{
				ServerList:  []string{"ceobebot.tech:9092"},
				RequireAck:  "NoResponse",
				Timeout:     "60",
				Partitioner: "RandomPartitioner",
				EnableSasl:  "FALSE",
				Username:    "",
				Password:    "",
			}},
			wantErr: false,
		},
		{
			name: "test_case_failed",
			fields: fields{
				client:    nil,
				connected: false,
			},
			args: args{cfg: ClusterConfig{
				ServerList:  []string{"ceobebot.tech:9093"},
				RequireAck:  "NoResponse",
				Timeout:     "60",
				Partitioner: "RandomPartitioner",
				EnableSasl:  "FALSE",
				Username:    "",
				Password:    "",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ProducerClient{
				client:    tt.fields.client,
				connected: tt.fields.connected,
			}
			bytes, _ := json.Marshal(tt.args.cfg)
			fmt.Println(string(bytes))
			if err := c.Connect(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	type fields struct {
		client    sarama.SyncProducer
		connected bool
	}
	type args struct {
		topic   string
		message string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUuid string
		wantErr  bool
	}{
		{
			name: "test_case_send_1",
			fields: fields{
				client:    nil,
				connected: false,
			},
			args: args{
				topic:   "hive_process_request",
				message: `{"data": 7,"date": "2020-01-26","type": 2,"uuid": "edwi"}`,
			},
			wantUuid: "",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ProducerClient{
				client:    tt.fields.client,
				connected: tt.fields.connected,
			}
			c.Connect(ClusterConfig{
				ServerList:  []string{""},
				RequireAck:  "NoResponse",
				Timeout:     "60",
				Partitioner: "RandomPartitioner",
				EnableSasl:  "FALSE",
				Username:    "",
				Password:    "",
			})
			gotUuid, err := c.Send(tt.args.topic, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUuid == tt.wantUuid {
				t.Errorf("Send() gotUuid = %v, want %v", gotUuid, tt.wantUuid)
			}
		})
	}
}
