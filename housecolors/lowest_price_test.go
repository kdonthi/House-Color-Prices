package housecolors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestTotalPrice(t *testing.T) {
	testTable := []struct{
		name string
		houseColorPrices []HouseColorPrices
		expectedResult int
	}{
		{
			houseColorPrices: []HouseColorPrices{
				HouseColorPrices{
					red: 1,
					green: 2,
					blue: 3,
				},
				HouseColorPrices{
					red: 17,
					green: 64,
					blue: 3,
				},
			},
			expectedResult: 4,
		},
		{
			houseColorPrices: []HouseColorPrices{
				HouseColorPrices{
					red: 1,
					green: 2,
					blue: 3,
				},
				HouseColorPrices{
					red: 5,
					green: 2,
					blue: 1,
				},
				HouseColorPrices{
					red: 17,
					green: 18,
					blue:4,
				},
			},
			expectedResult: 7,
		},
		{
			houseColorPrices: []HouseColorPrices{
				HouseColorPrices{
					red: 1,
					green: 2,
					blue: 3,
				},
				HouseColorPrices{
					red: 46,
					green: 21,
					blue: 86,
				},
				HouseColorPrices{
					red: 5,
					green: 2,
					blue: 1,
				},
				HouseColorPrices{
					red: 17,
					green: 18,
					blue:4,
				},
			},
			expectedResult: 31,
		},
	}

	for i, testCase := range testTable {
		if i > -1 {
			t.Run(testCase.name, func(t *testing.T) {
			result := lowestTotalPrice(testCase.houseColorPrices)
			assert.Equal(t, testCase.expectedResult, result)
		})	
		}
	}
}