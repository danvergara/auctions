package auctions

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		name        string
		input       []Bidder
		expected    *Winner
		expectError bool
	}{
		{
			name: "Auction #1",
			input: []Bidder{
				{
					Name:          "Sasha",
					InitialBid:    5000,
					MaxBid:        8000,
					AutoIncrement: 300,
				}, {
					Name:          "John",
					InitialBid:    6000,
					MaxBid:        8200,
					AutoIncrement: 200,
				}, {
					Name:          "Pat",
					InitialBid:    5500,
					MaxBid:        8500,
					AutoIncrement: 500,
				},
			},
			expected: &Winner{
				Name:   "Pat",
				MaxBid: 8500,
			},
		},
		{
			name: "Auction #2",
			input: []Bidder{
				{
					Name:          "Riley",
					InitialBid:    70000,
					MaxBid:        72500,
					AutoIncrement: 200,
				}, {
					Name:          "Morgan",
					InitialBid:    59900,
					MaxBid:        72500,
					AutoIncrement: 1500,
				}, {
					Name:          "Charlie",
					InitialBid:    62500,
					MaxBid:        72500,
					AutoIncrement: 800,
				},
			},
			expected: &Winner{
				Name:   "Riley",
				MaxBid: 72200,
			},
		},
		{
			name: "Auction #3",
			input: []Bidder{
				{
					Name:          "Alex",
					InitialBid:    250000,
					MaxBid:        300000,
					AutoIncrement: 50000,
				}, {
					Name:          "Jesse",
					InitialBid:    280000,
					MaxBid:        310000,
					AutoIncrement: 20100,
				}, {
					Name:          "Drew",
					InitialBid:    250100,
					MaxBid:        320000,
					AutoIncrement: 24700,
				},
			},
			expected: &Winner{
				Name:   "Jesse",
				MaxBid: 300100,
			},
		},
		{
			name: "Auction #4 with a tie",
			input: []Bidder{
				{
					Name:          "Sasha",
					InitialBid:    5000,
					MaxBid:        8000,
					AutoIncrement: 300,
				}, {
					Name:          "John",
					InitialBid:    6000,
					MaxBid:        8500,
					AutoIncrement: 500,
				}, {
					Name:          "Pat",
					InitialBid:    5500,
					MaxBid:        8500,
					AutoIncrement: 500,
				},
			},
			expected: &Winner{
				Name:   "John",
				MaxBid: 8500,
			},
		},
		{
			name: "Single bidder",
			input: []Bidder{
				{
					Name:          "Mario",
					InitialBid:    10000,
					MaxBid:        20000,
					AutoIncrement: 500,
				},
			},
			expected: &Winner{
				Name:   "Mario",
				MaxBid: 10000,
			},
			expectError: false,
		},
		{
			name: "Error case: Zero auto-increment",
			input: []Bidder{
				{Name: "Valid", InitialBid: 10, MaxBid: 100, AutoIncrement: 10},
				{Name: "Invalid", InitialBid: 10, MaxBid: 100, AutoIncrement: 0},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Error case: Negative auto-increment",
			input: []Bidder{
				{Name: "Invalid", InitialBid: 10, MaxBid: 100, AutoIncrement: -10},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Error case: Negative initial bid",
			input: []Bidder{
				{Name: "Invalid", InitialBid: -10, MaxBid: 100, AutoIncrement: 10},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Error case: Max bid less than initial bid",
			input: []Bidder{
				{Name: "Invalid", InitialBid: 100, MaxBid: 90, AutoIncrement: 10},
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := New(test.input...)
			got, err := a.Run()

			if test.expectError {
				if err == nil {
					t.Fatal("expected an error, but got none")
				}
				return
			}

			if test.expected == nil || got.Name != test.expected.Name ||
				got.MaxBid != test.expected.MaxBid {
				t.Fatalf("Run returned winner %v; expected %v", got, test.expected)
			}
		})
	}
}

func TestCalculateMaxBidPossible(t *testing.T) {

	tests := []struct {
		name          string
		initialBid    int64
		maxBid        int64
		autoIncrement int64
		expected      struct {
			autoIncrementCount int64
			maxPossibleBid     int64
		}
	}{
		{
			name:          "Sasha",
			initialBid:    5000,
			maxBid:        8000,
			autoIncrement: 300,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 10,
				maxPossibleBid:     8000,
			},
		},
		{
			name:          "John",
			initialBid:    6000,
			maxBid:        8200,
			autoIncrement: 200,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 11,
				maxPossibleBid:     8200,
			},
		},
		{
			name:          "Pat",
			initialBid:    5500,
			maxBid:        8500,
			autoIncrement: 500,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 6,
				maxPossibleBid:     8500,
			},
		},
		{
			name:          "Riley",
			initialBid:    70000,
			maxBid:        72500,
			autoIncrement: 200,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 12,
				maxPossibleBid:     72400,
			},
		},
		{
			name:          "Morgan",
			initialBid:    59900,
			maxBid:        72500,
			autoIncrement: 1500,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 8,
				maxPossibleBid:     71900,
			},
		},
		{
			name:          "Charlie",
			initialBid:    62500,
			maxBid:        72500,
			autoIncrement: 800,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 12,
				maxPossibleBid:     72100,
			},
		},
		{
			name:          "Alex",
			initialBid:    250000,
			maxBid:        300000,
			autoIncrement: 50000,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 1,
				maxPossibleBid:     300000,
			},
		},
		{
			name:          "Jesse",
			initialBid:    280000,
			maxBid:        310000,
			autoIncrement: 20100,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 1,
				maxPossibleBid:     300100,
			},
		},
		{
			name:          "Drew",
			initialBid:    250100,
			maxBid:        320000,
			autoIncrement: 24700,
			expected: struct {
				autoIncrementCount int64
				maxPossibleBid     int64
			}{
				autoIncrementCount: 2,
				maxPossibleBid:     299500,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			maxPossibleBid := calculateMaxBidPossible(
				test.initialBid,
				test.maxBid,
				test.autoIncrement,
			)

			if maxPossibleBid != test.expected.maxPossibleBid {
				t.Fatalf(
					"calculateMaxBidPossible returned %d as max bid possible; expected %d as max bid possible",
					maxPossibleBid,
					test.expected.maxPossibleBid,
				)
			}
		})
	}
}
