package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	FirstName     string
	LastName      string
	Email         string
	Age           int
	MonthlySalary []MonthlySalary
}

type MonthlySalary struct {
	Basic int `json:"basic"`
	HRA   int `json:"hra"`
	TA    int `json:"ta"`
}

func main() {
	jsonFile, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Println("ERROR! Failed to read the file.")
	}

	var jsonData []Data
	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		log.Println("ERROR! Failed to unmarshal the file.")
	}

	r := gin.Default()

	r.GET("/allusers", func(c *gin.Context) {
		c.JSON(http.StatusOK, jsonData)
	})

	r.GET("/user", func(c *gin.Context) {
		start := c.Query("/letter")
		var filteredUsers []Data
		for _, user := range jsonData {
			if strings.HasPrefix(strings.ToLower(user.LastName), strings.ToLower(start)) {
				filteredUsers = append(filteredUsers, user)
			}
		}
		c.JSON(http.StatusOK, filteredUsers)
	})

	if err := r.Run(":90"); err != nil {
		log.Println("ERROR! Failed to load GIN server.")
	}
}
