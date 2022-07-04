package visualization

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMapDataHandler(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapDataHandler(tt.args.ctx)
		})
	}
}
