package handler

import (
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	Utils "gitlab.com/tesseract/backend/utils"
)

type Complaint struct {
	ID       int     `json:"id"`
	Date     string  `json:"date"`
	Order    string  `json:"order"`
	Machine  string  `json:"machine"`
	Quantity int     `json:"quantity"`
	Cost     float32 `json:"cost"`
	Intern   bool    `json:"intern"`
	Reason   string  `json:"reason"`
	Material string  `json:"material"`
}

type Material struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Reason struct {
	ID      int    `json:"id"`
	Feature string `json:"feature"`
	Reason  string `json:"reason"`
}

var material = []Material{
	{ID: Utils.RandomStringNumber(4, 5), Name: "Kolben"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Kolbenring"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Kolbenbolzen"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Flansch"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Kolbenträger"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "ZSB Gehäuse PL73"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "ZSB Gehäuse PL74"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "ZSB Gehäuse PL75"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Geteilter Ring"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Gehäuse"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Endscheibe"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Gehäuse"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Buchse"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Buchse"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Lagerbuchse"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Ringkolben"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Kolben B"},
	{ID: Utils.RandomStringNumber(4, 5), Name: "Stauscheibe"},
}

var reason = []Reason{
	{ID: 1, Feature: "Durchmesser", Reason: "zu groß"},
	{ID: 2, Feature: "Durchmesser", Reason: "zu klein"},
	{ID: 3, Feature: "Länge", Reason: "zu kurz"},
	{ID: 4, Feature: "Länge", Reason: "zu lang"},
	{ID: 5, Feature: "Position", Reason: "zu groß"},
	{ID: 6, Feature: "Position", Reason: "zu klein"},
	{ID: 7, Feature: "Konzentrizität", Reason: "versetzt"},
	{ID: 8, Feature: "Messprüfung", Reason: "zu groß"},
	{ID: 9, Feature: "Lunker", Reason: "zu rauh"},
	{ID: 10, Feature: "Planlauf", Reason: "zu groß"},
	{ID: 11, Feature: "Fläche", Reason: "beschädigt"},
}

// function to mock the data for the complaints. It generates between 500 to 1000 rows in a timeframe from a given datetime to a given end datetime
func generateComplaintData() []Complaint {
	var complaints = []Complaint{}

	// start date time
	var start = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	// end datetime is today
	var end = time.Now()

	// generate a random number of rows between 10.000 and 25.000
	var rows = rand.Intn(10000) + 25000

	// generate rows of orders and machines as random strings
	var randomOrders = []string{}
	var randomMachines = []string{}
	var randomCosts = []float32{}
	var randomQuantities = []int{}

	for i := 0; i < 5; i++ {
		// generate a random 6 digit number string
		randomSixDigitNumber := Utils.RandomStringNumber(6, 6)
		// generate a random number string between 2 and 3 digits
		randomTwoOrThreeDigitNumber := Utils.RandomStringNumber(2, 3)

		// if randomTwoOrThreeDigitNumber is less than 100, add a 0 to the front
		if len(randomTwoOrThreeDigitNumber) < 3 {
			randomTwoOrThreeDigitNumber = "0" + randomTwoOrThreeDigitNumber
		}

		randomOrders = append(randomOrders, "PA0"+randomSixDigitNumber)
		randomMachines = append(randomMachines, "MA"+randomTwoOrThreeDigitNumber)
	}

	for i := 0; i < rows; i++ {
		randomCosts = append(randomCosts, rand.Float32()*100)
		randomQuantities = append(randomQuantities, rand.Intn(100))
	}

	// generate rows of quantities and costs as random numbers
	// generate the random data
	for i := 0; i < rows; i++ {
		// pick a random order, machine, quantity and cost
		var order = randomOrders[rand.Intn(len(randomOrders))]
		var machine = randomMachines[rand.Intn(len(randomMachines))]
		var quantity = randomQuantities[rand.Intn(len(randomQuantities))]
		var cost = randomCosts[rand.Intn(len(randomCosts))]
		// generate a random date between the start and end date
		var date = start.Add(time.Duration(rand.Int63n(end.UnixNano()-start.UnixNano())) * time.Nanosecond)

		// generate random true or false
		var intern = rand.Intn(2) == 1

		// pick a random reason
		var reason = reason[rand.Intn(len(reason))]

		// pick a random material
		var material = material[rand.Intn(len(material))]

		if len(material.ID) < 5 {
			material.ID = "0" + material.ID
		}

		// append the data to the array
		complaints = append(complaints, Complaint{
			ID:       i,
			Date:     date.Format("2006-01-02"), // format the date to yyyy-mm-dd
			Order:    order,
			Machine:  machine,
			Quantity: quantity,
			Cost:     cost,
			Intern:   intern,
			Reason:   reason.Feature + " - " + reason.Reason,
			Material: material.ID + " - " + material.Name,
		})
	}
	// sort the data by date
	sort.Slice(complaints, func(i, j int) bool {
		return complaints[i].Date < complaints[j].Date
	})
	return complaints
}

var ComplaintData = generateComplaintData()

// API Handler to get all complaints. Can also take a query parameter to get the complaints by start and end date
func GetAllComplaints(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	// make a copy of the complaints
	var complaints = []Complaint{}
	complaints = append(complaints, ComplaintData...)

	// if there is a start date and end date, filter the data
	if startDate != "" && endDate != "" {
		// parse the dates
		start, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		end, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// filter the data by date range
		complaints = filterComplaintsByDateRange(complaints, start, end)
	}

	c.JSON(http.StatusOK, complaints)
}

// API Handler to get a single complaint by id.
func GetComplaint(c *gin.Context) {
	id := c.Param("id")

	// if no id is in the context, return a error
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No id provided"})
		return
	}
	// id is a string, so we need to convert it to an int
	idInt, _ := strconv.Atoi(id)

	for _, complaint := range ComplaintData {
		if complaint.ID == idInt {
			c.JSON(http.StatusOK, complaint)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Complaint not found"})
}

// function to filter the complaint data between a given date range
func filterComplaintsByDateRange(complaintData []Complaint, startDate time.Time, endDate time.Time) []Complaint {
	// create a new array to store the filtered data
	var filteredData = []Complaint{}

	for _, complaint := range complaintData {
		// parse the date of the complaint
		date, err := time.Parse("2006-01-02", complaint.Date)
		if err != nil {
			return []Complaint{}
		}

		// if the date is between the start and end date, append it to the filtered data
		if date.After(startDate) && date.Before(endDate) {
			filteredData = append(filteredData, complaint)
		}
	}

	return filteredData
}

// function to filter the complaint data by a given parameter
func FilterComplaintsByParameter(complaints []Complaint, parameter string, value string) []Complaint {
	var filteredComplaints = []Complaint{}
	for _, complaint := range complaints {
		// if field exists, check if field value is equal to the value
		complaintFieldValue := Utils.GetStructFieldValue(complaint, parameter)
		if complaintFieldValue != "" && complaintFieldValue == value {
			filteredComplaints = append(filteredComplaints, complaint)
		}
	}

	return filteredComplaints
}
