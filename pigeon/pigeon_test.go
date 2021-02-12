package pigeon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	testCases := []struct {
		desc           string
		testPigeon     Pigeon
		expectedResult string
	}{
		{
			desc:           "returns current location of Chorlton",
			testPigeon:     CreateNewPigeon("Bob", "Brown", "Chorlton"),
			expectedResult: "Chorlton",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expectedResult, tC.testPigeon.GetLocation())
		})
	}
}

func TestFlyToLocation(t *testing.T) {
	testCases := []struct {
		desc           string
		testPigeon     Pigeon
		newLocation    string
		expectedResult string
	}{
		{
			desc:           "Current location is correctly updated with new location",
			testPigeon:     CreateNewPigeon("Bob", "Brown", "Chorlton"),
			newLocation:    "Macclesfield",
			expectedResult: "Macclesfield",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.testPigeon.FlyToLocation(tC.newLocation)
			assert.Equal(t, tC.expectedResult, tC.testPigeon.currentLocation)

		})
	}
}

func TestReturnHome(t *testing.T) {
	testCases := []struct {
		desc       string
		testPigeon Pigeon
	}{
		{
			desc:       "returns to home location",
			testPigeon: CreateNewPigeon("Bob", "Brown", "Chorlton"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.testPigeon.ReturnHome()
			assert.Equal(t, tC.testPigeon.currentLocation, tC.testPigeon.HomeLocation)

		})
	}
}

func TestCreateNewPigeon(t *testing.T) {
	testCases := []struct {
		desc                    string
		name                    string
		colour                  string
		homelocation            string
		expectedCurrentLocation string
	}{
		{
			desc:                    "Creates a new pigeon with the provided parameters",
			name:                    "Bob",
			colour:                  "Green",
			homelocation:            "Wilmslow",
			expectedCurrentLocation: "Wilmslow",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			np := CreateNewPigeon(tC.name, tC.colour, tC.homelocation)
			assert.Equal(t, tC.expectedCurrentLocation, np.currentLocation)
		})
	}
}
