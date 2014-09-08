package betfred

import (
	"testing"
	"github.com/stretchr/testify/assert"
    "io/ioutil"
)

func GetData() []Fixture {
	data, _ := ioutil.ReadFile("Football-Premiership.xml")
	fixtures := Parse(data)

	return fixtures;
}

func TestParse_Returns_Fixture_List(t *testing.T) {
	fixtures := GetData()

	assert.NotEqual(t, fixtures, "")
	assert.Equal(t, fixtures[0].Title, "Arsenal v Manchester City")
	assert.Equal(t, fixtures[0].Date, "20140913")
	assert.Equal(t, fixtures[0].Time, "1245")
}

func TestParse_Splits_Home_Away_Teams(t *testing.T) {
	fixtures := GetData()

	assert.NotEqual(t, fixtures, "")
	assert.Equal(t, fixtures[0].Home, "Arsenal")
	assert.Equal(t, fixtures[0].Away, "Manchester City")
}

func TestParse_Returns_Odds(t *testing.T) {
	fixtures := GetData()

	assert.NotEqual(t, fixtures, "")
	assert.Equal(t, fixtures[0].Bets[0].Name, "Match Result")
	assert.Equal(t, fixtures[0].Bets[0].Odds[0].Name, "Arsenal")
	assert.Equal(t, fixtures[0].Bets[0].Odds[0].Price, "7/4")
}