package tools

import "testing"

func TestCateRetionCode(t *testing.T) {
	type args struct {
		callsign string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t",
			args: args{
				callsign: "BG0CFD",
			},
			want: 65,
		},
		{
			name: "t",
			args: args{
				callsign: "BG0GM",
			},
			want: 54,
		},
		{
			name: "t",
			args: args{
				callsign: "BH3WNL",
			},
			want: 14,
		}, {
			name: "t",
			args: args{
				callsign: "BH4FCY",
			},
			want: 31,
		}, {
			name: "t",
			args: args{
				callsign: "BG5UWQ",
			},
			want: 35,
		}, {
			name: "t",
			args: args{
				callsign: "BG5CW",
			},
			want: 33,
		}, {
			name: "t",
			args: args{
				callsign: "BA7NQ",
			},
			want: 44,
		}, {
			name: "t",
			args: args{
				callsign: "BG8BXM",
			},
			want: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CateRetionCode(tt.args.callsign); got != tt.want {
				t.Errorf("CateRetionCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
