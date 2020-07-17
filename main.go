package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type date struct {
	Day   int `json:"date" binding:"required"`
	Month int `json:"month" binding:"required"`
}

type trivia struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
	Year   int    `json:"year"`
}

type message struct {
	Title        string
	Date         int
	Month        int
	TriviaNumber int
	Trivia       string
}

func triviaGet(c *gin.Context) {
	ts := time.Now()
	date := date{
		Day:   ts.Day(),
		Month: int(ts.Month()),
	}
	url := fmt.Sprintf("http://numbersapi.com/%s/%s/date?json", strconv.Itoa(date.Month), strconv.Itoa(date.Day))
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// validate response
	triviaRes := trivia{}
	jsonErr := json.Unmarshal(body, &triviaRes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	msg := message{
		Title:        "Trivia number by current date",
		Date:         date.Day,
		Month:        date.Month,
		TriviaNumber: triviaRes.Number,
		Trivia:       triviaRes.Text,
	}
	c.JSON(200, msg)
}

func main() {
	r := gin.Default()
	r.GET("/trivia", triviaGet)

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
