package visualization

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestOverviewDataHandler(t *testing.T) {
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
			OverviewDataHandler(tt.args.ctx)
		})
	}
}
