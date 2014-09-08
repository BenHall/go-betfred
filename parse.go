package betfred

import (
    "fmt"
	"strings"
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "sort"
)
 
type Query struct {
	FixtureList []Fixture `xml:"event"`
}
 
type Fixture struct {
	Id    string `xml:"eventid,attr" json:"id"`
	Title   string `xml:"name,attr" json:"title"`
	Home 	string `json:"home"`
	Away 	string `json:"away"`
	Date	string `xml:"date,attr" json:"date"`
	Time	string `xml:"time,attr" json:"time"`
	Bets     []Bet `xml:"bettype" json:"bets"`
}

type Bet struct {
    Id    string `xml:"bettypeid,attr" json:"id"`
	Name    string `xml:"name,attr" json:"name"`
	Odds	[]Odd `xml:"bet" json:"odds"`
}

type Odd struct {
	Id    string `xml:"id,attr" json:"id"`
	Value    string `xml:"had-value,attr" json:"value"`
	Name    string `xml:"name,attr" json:"name"`
	Price    string `xml:"price,attr" json:"price"`
}

type byDate []Fixture
func (v byDate) Len() int { return len(v) }
func (v byDate) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v byDate) Less(i, j int) bool { return v[i].Date < v[j].Date } 

func Parse(contents []byte) []Fixture {
	var q Query
	xml.Unmarshal(contents, &q)

	transformed_array := transform(q.FixtureList)
    sort.Sort(byDate(transformed_array))
    return transformed_array
}

func transform(s []Fixture) []Fixture {
    var p []Fixture

    for _, v := range s {
    	split := strings.Split(v.Title, " v ")

    	if(len(split) == 2) {
    		v.Home = split[0];
    		v.Away = split[1];
            p = append(p, v)

    	}
	}

	return p;
}

func RequestPremiership() []byte {
    url := "http://xml.betfred.com/Football-Premiership.xml"
    return request(url)
}

func request(url string) []byte {
    response, err := http.Get(url)
    if err != nil {
        fmt.Printf("%s", err)
        return []byte{}     
    }
    if response != nil {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
    	if err != nil {
        	fmt.Printf("%s", err)
            return []byte{} 
    	}

        return contents
    }

    return []byte{}	
}
