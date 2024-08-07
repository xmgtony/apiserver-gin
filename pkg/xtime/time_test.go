package xtime

import (
	"testing"
	"time"
)

func TestTimeMarshalJSON(t *testing.T) {
	testCases := []struct {
		name  string
		input Time
		want  []byte
	}{
		{
			name:  "Case 1",
			input: Time(time.Date(2022, 6, 15, 12, 30, 0, 0, time.UTC)),
			want:  []byte(`"2022-06-15 12:30:00"`),
		},
		{
			name:  "Case 2",
			input: Time(time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)),
			want:  []byte(`"2023-12-31 23:59:59"`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.MarshalJSON()
			if err != nil {
				t.Errorf("MarshalJSON() returned an unexpected error: %v", err)
			}
			if string(got) != string(tc.want) {
				t.Errorf("MarshalJSON() = %s, want %s", string(got), string(tc.want))
			}
		})
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		wantError bool
	}{
		{
			name:      "Null data",
			data:      []byte("null"),
			wantError: false,
		},
		{
			name:      "Valid JSON string",
			data:      []byte(`"2018-11-25 20:04:51"`),
			wantError: false,
		},
		{
			name:      "Invalid JSON string",
			data:      []byte("2018-11-25 20:04:51"),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var time1 Time
			err := time1.UnmarshalJSON(tt.data)

			if tt.wantError && err == nil {
				t.Error("Expected error, but got nil")
			} else if !tt.wantError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestTime_Scan(t *testing.T) {
	tests := []struct {
		name    string
		time    Time
		value   interface{}
		wantErr bool
	}{
		{
			name:    "nil value",
			time:    Time{},
			value:   nil,
			wantErr: false,
		},
		{
			name:    "time.Time value",
			time:    Time{},
			value:   time.Now(),
			wantErr: false,
		},
		{
			name:    "string value with correct format",
			time:    Time{},
			value:   "2021-01-01 12:00:00",
			wantErr: false,
		},
		{
			name:    "string value with incorrect format",
			time:    Time{},
			value:   "invalid format",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.time.Scan(tt.value)

			if err != nil && !tt.wantErr {
				t.Errorf("Time.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
