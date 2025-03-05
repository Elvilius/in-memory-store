package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCompute_Parse(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name    string
		args    args
		want    PreparedCommand
		wantErr bool
	}{
		{
			name: "Successful SET command ", args: args{cmd: "SET weather_2_pm cold_moscow_weather"}, want: PreparedCommand{
				Cmd:  Command("SET"),
				Key:  "weather_2_pm",
				Value: "cold_moscow_weather",
			},
			wantErr: false,
		},
		{
			name: "Successful GET command", args: args{cmd: "GET weather_2_pm"}, want: PreparedCommand{
				Cmd: Command("GET"),
				Key: "weather_2_pm",
			},
			wantErr: false,
		},
		{
			name: "Successful DEL command", args: args{cmd: "DEL weather_2_pm"}, want: PreparedCommand{
				Cmd: Command("DEL"),
				Key: "weather_2_pm",
			},
			wantErr: false,
		},
		{
			name: "Empty command", args: args{cmd: ""}, want: PreparedCommand{},
			wantErr: true,
		},
		{
			name: "Random string", args: args{cmd: "randomd random rand"}, want: PreparedCommand{},
			wantErr: true,
		},
	}

	logger := &zap.Logger{}
	c := New(logger)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Parse(tt.args.cmd)
			assert.Equal(t, got, tt.want)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
