package handler

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	Utils "gitlab.com/tesseract/backend/utils"
)

type Defect struct {
	Week       int     `json:"week"`
	Year       int     `json:"year"`
	PPM        float64 `json:"ppm"`
	PPMPercent float64 `json:"ppmPercent"`
	Total      int     `json:"total"`
	Feature    string  `json:"feature"`
}

// GetDefectsByWeek returns the defects for a given week
func GetDefectsByWeek(c *gin.Context) {
	// end date from the get request
	endDate := c.Query("endDate")
	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}

	// parse the end date
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		// send an error if the date is not valid
		c.JSON(400, gin.H{
			"error": "Invalid date",
		})
	}

	// generate an array of iso weeks to the end date
	isoWeeks := Utils.GenerateCalendarWeeks(end)
	// make a copy of the complaints from fake data
	var complaints = []Complaint{}
	complaints = append(complaints, ComplaintData...)
	// filter by iso weeks
	complaints = filterComplaintsByIsoWeek(complaints, isoWeeks, end.Year())

	// prepare data
	var defects = []Defect{}

	// group the complaints by their Reason

	groupedComplaints := GroupComplaintsByKey(complaints, "Reason")
	// loop through the groups
	for key, value := range groupedComplaints {
		// loop through the iso weeks
		for _, isoWeek := range isoWeeks {
			// get the total number of complaints for the week
			complaintsByIsoWeek := getComplaintsByIsoWeek(value, isoWeek, end.Year())
			total := len(complaintsByIsoWeek)
			// calculate the ppm
			ppm := float64(total) / float64(1000) * float64(1000000)
			// calculate the ppm percent
			ppmPercent := ppm / float64(1000000) * float64(100)
			// add the data to the defects array
			defects = append(defects, Defect{
				Week:       isoWeek,
				Year:       end.Year(),
				PPM:        ppm,
				PPMPercent: ppmPercent,
				Total:      total,
				Feature:    key,
			})

		}

	}

	c.JSON(200, defects)
}

// FilterComplaintsByIsoWeek returns the filtered complaint data by a provided array of iso weeks
func filterComplaintsByIsoWeek(complaintData []Complaint, isoWeeks []int, year int) []Complaint {
	// create a new array to store the filtered data
	var filteredData = []Complaint{}

	for _, complaint := range complaintData {
		// parse the date of the complaint
		date, err := time.Parse("2006-01-02", complaint.Date)
		if err != nil {
			return []Complaint{}
		}

		// get the iso week of the complaint
		_, week := date.ISOWeek()

		// if the week of the complaint is in the array of iso weeks, append it to the filtered data. Also check if the year is the same
		for _, isoWeek := range isoWeeks {
			if week == isoWeek && date.Year() == year {
				filteredData = append(filteredData, complaint)
			}
		}
	}

	return filteredData
}

// getComplaintsByIsoWeek returns the number of complaints for a given iso week
func getComplaintsByIsoWeek(complaints []Complaint, isoWeek int, year int) []Complaint {
	// create a new array to store the filtered data
	var filteredData = []Complaint{}

	for _, complaint := range complaints {
		// parse the date of the complaint
		date, err := time.Parse("2006-01-02", complaint.Date)
		if err != nil {
			return []Complaint{}
		}

		// get the iso week of the complaint
		_, week := date.ISOWeek()

		// if the week of the complaint is in the array of iso weeks, append it to the filtered data. Also check if the year is the same
		if week == isoWeek && date.Year() == year {
			filteredData = append(filteredData, complaint)
		}
	}

	return filteredData
}

// GeneratePPMByWeek returns a ppm value for a given week
func GeneratePPMByWeek(complaints []Complaint, week int, year int) float64 {
	// get the number of defects for the week
	defects := len(filterComplaintsByIsoWeek(complaints, []int{week}, year))

	// calculate the ppm
	ppm := float64(defects) / float64(100000)

	return ppm
}

// GroupComplaintsByKey returns a map of grouped complaints by a given struct field
func GroupComplaintsByKey(complaints []Complaint, key string) map[string][]Complaint {
	// create a new map
	groupedComplaints := make(map[string][]Complaint)

	// loop through the complaints
	for _, complaint := range complaints {
		// get the value of the key
		value := reflect.ValueOf(complaint).FieldByName(key).String()

		// if the value is not empty, append the complaint to the map
		if value != "" {
			groupedComplaints[value] = append(groupedComplaints[value], complaint)
		}
	}

	return groupedComplaints
}
