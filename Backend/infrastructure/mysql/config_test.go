package mysql

import "testing"

func TestConfig_toMaxIdleConn(t *testing.T) {
	type fields struct {
		Address           string
		Username          string
		Password          string
		DatabaseName      string
		MaxOpenConnection string
		MaxIdleConnection string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test_case_success",
			fields: fields{
				Address:           "",
				Username:          "",
				Password:          "",
				DatabaseName:      "",
				MaxOpenConnection: "",
				MaxIdleConnection: "114514",
			},
			want: 114514,
		},
		{
			name: "test_case_failed",
			fields: fields{
				Address:           "",
				Username:          "",
				Password:          "",
				DatabaseName:      "",
				MaxOpenConnection: "",
				MaxIdleConnection: "iiolys",
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				Address:           tt.fields.Address,
				Username:          tt.fields.Username,
				Password:          tt.fields.Password,
				DatabaseName:      tt.fields.DatabaseName,
				MaxOpenConnection: tt.fields.MaxOpenConnection,
				MaxIdleConnection: tt.fields.MaxIdleConnection,
			}
			if got := c.toMaxIdleConn(); got != tt.want {
				t.Errorf("toMaxIdleConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_toMaxOpenConn(t *testing.T) {
	type fields struct {
		Address           string
		Username          string
		Password          string
		DatabaseName      string
		MaxOpenConnection string
		MaxIdleConnection string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				Address:           tt.fields.Address,
				Username:          tt.fields.Username,
				Password:          tt.fields.Password,
				DatabaseName:      tt.fields.DatabaseName,
				MaxOpenConnection: tt.fields.MaxOpenConnection,
				MaxIdleConnection: tt.fields.MaxIdleConnection,
			}
			if got := c.toMaxOpenConn(); got != tt.want {
				t.Errorf("toMaxOpenConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
