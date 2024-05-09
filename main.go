package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkExperience struct {
	Place      string   `json:"place"`
	Title      string   `json:"title"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	JobDetails []string `json:"jobDetails"`
	Technology []string `json:"technology"`
}

// this is temporary until I hookup DB code.
var jobs = []WorkExperience{
	{
		Place:      "ActivTrak",
		Title:      "Senior Software Engineer",
		StartDate:  "June, 2022",
		EndDate:    "Feb, 2022",
		JobDetails: []string{"Shopping Cart", "Tray", "Data Shuttling"},
		Technology: []string{"Python", "React", "GCP", "C#", "CircleCi"},
	},
	{
		Place:      "Helix Technologies",
		Title:      "Senior Software Developer",
		StartDate:  "Jan, 2019",
		EndDate:    "Feb, 2020",
		JobDetails: []string{"Dashboard Metrics", "Excel Creation and SFTP", "Linux Server Admin"},
		Technology: []string{"Python", "Angular", "AWS", "PHP", "Bash"},
	},
}

func main() {
	router := gin.Default()

	// Get Routes
	router.GET("/", getWorkHistory)
	router.GET("/getWorkHistoryByPlace", getWorkHistoryByPlace)
	router.GET("/getWorkHistoryByTitle", getWorkHistoryByTitle)
	router.GET("/getWorkHistoryByTechnology", getWorkHistoryByTechnology)

	// Post Routes
	router.POST("/postWorkHistory", postWorkHistory)

	router.Run("localhost:8080")
}

// getWorkHistory responds with job entries of previous work.
func getWorkHistory(c *gin.Context) {
	// TODO: Need a db fetch instead of this line.
	c.IndentedJSON(http.StatusOK, jobs)
}

func getWorkHistoryByPlace(c *gin.Context) {
	//pass
}

func getWorkHistoryByTitle(c *gin.Context) {
	//pass
}

func getWorkHistoryByTechnology(c *gin.Context) {
	//pass
}

// Post new experience
// TODO: Add authentication to this
func postWorkHistory(c *gin.Context) {
	var newWork WorkExperience

	// TODO: add better validation/error handling here.
	// Call BindJSON to bind the received JSON to newWork
	if err := c.BindJSON(&newWork); err != nil {
		return
	}

	// TODO: Need to convert this to database insertion.
	jobs = append(jobs, newWork)
	c.IndentedJSON(http.StatusCreated, newWork)
}
