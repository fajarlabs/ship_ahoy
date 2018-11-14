package main

import (
	"testing"
)

func TestToInt(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected int
	}{
		{int(9), 9},
		{int64(121), 121},
		{string("23"), 23},
		{float64(99.4), 99},
	}

	for _, testCase := range testCases {
		answer := toInt(testCase.value)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.value, testCase.expected, answer)
		}
	}
}

func TestToString(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected string
	}{
		{int(9), "9"},
		{int64(121), "121"},
		{string("23"), "23"},
		{float64(99.4), "99.4"},
	}

	for _, testCase := range testCases {
		answer := toString(testCase.value)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.value, testCase.expected, answer)
		}
	}
}

func TestToFloat64(t *testing.T) {
	testCases := []struct {
		value    interface{}
		expected float64
	}{
		{int(9), 9},
		{int64(121), 121},
		{string("23"), 23},
		{float64(99.4), 99.4},
	}

	for _, testCase := range testCases {
		answer := toFloat64(testCase.value)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.value, testCase.expected, answer)
		}
	}
}

// TODO:
// Because this hardcodes the same values that are hardcoded into
// visibleFromApt() it makes it feel like a change detector test.
// What about factoring the hardcoded rectangle boundaries into
// a global area that this function can pull from?
func TestVisibleFromApt(t *testing.T) {
	testCases := []struct {
		lat      float64
		lon      float64
		expected bool
	}{
		{1.1, 2.2, false},
		{37.8052, -122.48, true},   // bottom left corner
		{37.8613, -122.48, true},   // top left corner
		{37.8052, -122.4092, true}, // bottom right corner
		{37.82, -122.46, true},     // mid triangle
		{37.805, -122.49, false},   // outside bottom left corner
		{37.87, -122.49, false},    // outside top left corner
		{37.805, -122.4, false},    // outside bottom right corner
	}

	for _, testCase := range testCases {
		answer := visibleFromApt(testCase.lat, testCase.lon)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v, %v expected %v, got %v", testCase.lat, testCase.lon, testCase.expected, answer)
		}
	}
}

func TestBox(t *testing.T) {
	testCases := []struct {
		lat          float64
		lon          float64
		nmiles       float64
		expectedLatA float64
		expectedLonA float64
		expectedLatB float64
		expectedLonB float64
	}{
		{1, 2, 0, 1, 2, 1, 2},
	}

	for _, testCase := range testCases {
		latA, lonA, latB, lonB := box(testCase.lat, testCase.lon, testCase.nmiles)
		if latA != testCase.expectedLatA || lonA != testCase.expectedLonA || latB != testCase.expectedLatB || lonB != testCase.expectedLonB {
			t.Errorf("ERROR: For %v, %v, %v expected %v, %v, %v, %v, got %v, %v, %v, %v", testCase.lat, testCase.lon, testCase.nmiles, testCase.expectedLatA, testCase.expectedLonA, testCase.expectedLatB, testCase.expectedLonB, latA, lonA, latB, lonB)
		}
	}
}
