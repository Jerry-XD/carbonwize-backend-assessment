package calculateEmission

import (
	"testing"
)

type TestCase struct {
	name           string
	activityType   string
	inputValue     float64
	emissionFactor float64
	expected       float64
}

func TestCalculateEmission(t *testing.T) {
	var testCases = []*TestCase{
		{
			name:           "Success",
			activityType:   "transportation",
			inputValue:     100,
			emissionFactor: 0.254,
			expected:       25.4,
		},
		{
			name:           "Zero value input",
			activityType:   "transportation",
			inputValue:     0,
			emissionFactor: 1.5,
			expected:       0,
		},
		{
			name:           "Zero emission factor",
			activityType:   "walking",
			inputValue:     50,
			emissionFactor: 0,
			expected:       0,
		},
		{
			name:           "High value input",
			activityType:   "energy_consumption",
			inputValue:     10000,
			emissionFactor: 2.867,
			expected:       28670,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateEmission(tt.activityType, tt.inputValue, tt.emissionFactor)
			if result != tt.expected {
				t.Errorf("calculateEmission(%q, %f, %f) = %f; want %f", tt.activityType, tt.inputValue, tt.emissionFactor, result, tt.expected)
			}
		})
	}
}
