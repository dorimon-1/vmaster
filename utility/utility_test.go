package utility

import "testing"

func Test_splitKeyValue(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name      string
		args      args
		wantKey   string
		wantValue string
	}{
		{
			name: "Test splitKeyValue",
			args: args{
				arg: "key=value",
			},
			wantKey:   "key",
			wantValue: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotValue := splitKeyValue(tt.args.arg)
			if gotKey != tt.wantKey {
				t.Errorf("splitKeyValue() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotValue != tt.wantValue {
				t.Errorf("splitKeyValue() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
