package mysql

import (
	"reflect"
	"testing"
	"xorm.io/xorm"
)

func TestClient_Connect(t *testing.T) {
	type fields struct {
		connection *xorm.Engine
		connected  bool
	}
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				connection: tt.fields.connection,
				connected:  tt.fields.connected,
			}
			if err := c.Connect(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Create(t *testing.T) {
	type fields struct {
		connection *xorm.Engine
		connected  bool
	}
	type args struct {
		opts []interface{}
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantEffected []int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				connection: tt.fields.connection,
				connected:  tt.fields.connected,
			}
			gotEffected, err := c.Create(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEffected, tt.wantEffected) {
				t.Errorf("Create() gotEffected = %v, want %v", gotEffected, tt.wantEffected)
			}
		})
	}
}

func TestClient_Delete(t *testing.T) {
	type fields struct {
		connection *xorm.Engine
		connected  bool
	}
	type args struct {
		condition interface{}
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSuccess bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				connection: tt.fields.connection,
				connected:  tt.fields.connected,
			}
			gotSuccess, err := c.Delete(tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("Delete() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}

func TestClient_Retrieve(t *testing.T) {
	type fields struct {
		connection *xorm.Engine
		connected  bool
	}
	type args struct {
		condition interface{}
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSuccess bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				connection: tt.fields.connection,
				connected:  tt.fields.connected,
			}
			gotSuccess, err := c.Retrieve(tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("Retrieve() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}

func TestClient_Sync(t *testing.T) {
	type fields struct {
		connection *xorm.Engine
		connected  bool
	}
	type args struct {
		opts []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				connection: tt.fields.connection,
				connected:  tt.fields.connected,
			}
			if err := c.Sync(tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Update(t *testing.T) {
	type fields struct {
		connection *xorm.Engine
		connected  bool
	}
	type args struct {
		newOpt interface{}
		oldOpt interface{}
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantSuccess bool
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				connection: tt.fields.connection,
				connected:  tt.fields.connected,
			}
			gotSuccess, err := c.Update(tt.args.newOpt, tt.args.oldOpt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("Update() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}
