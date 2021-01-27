package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Weather struct {
	City     string
	Temp     string
	Wind     string
	Data     string
	AirLevel string
	Imgurl   string
	Humidity string
}

const (
	URL = "https://www.heweather.com/"
)

func removeBlank(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}
func GetWeather(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(doc, err)
	now := doc.Find("div.now").Eq(0)

	// row0 := now.Find(".row").Eq(0)
	// row1 := now.Find(".row").Eq(1)
	row2 := now.Find(".row").Eq(2)

	city := now.Find(".name").Text()
	city = removeBlank(city)
	temp := now.Find(".tmp").Text()
	temp = removeBlank(temp)
	wind := row2.Find("div").Eq(0).Text()
	wind = removeBlank(wind)
	data := now.Find(".txt").Text()
	data = removeBlank(data)
	air := now.Find(".air").Text()
	air = removeBlank(air)
	imgurl, _ := now.Find("img").Attr("src")
	imgurl = removeBlank(imgurl)
	humidity := row2.Find("div").Eq(2).Text()
	humidity = removeBlank(humidity)
	fmt.Println(city, temp, wind, data, air, imgurl, humidity)

}

func main() {
	GetWeather(URL)
}
