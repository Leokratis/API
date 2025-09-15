package v1

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Helper function to create standard API response
func createAPIResponse(data interface{}, errors []string, requestID string) APIResponse {
	return APIResponse{
		Data:   data,
		Errors: errors,
		Metadata: Metadata{
			Timestamp: time.Now(),
			Version:   "v1",
			RequestID: requestID,
		},
	}
}

// Helper function to parse date in DDMMYYYY format
func parseDate(dateStr string) (time.Time, error) {
	if len(dateStr) != 8 {
		return time.Time{}, fmt.Errorf(`Invalid date format; Expected DDMMYYYY but recieved str[%v]`, len(dateStr))
	}

	day, err := strconv.Atoi(dateStr[:2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day")
	}

	month, err := strconv.Atoi(dateStr[2:4])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month")
	}

	year, err := strconv.Atoi(dateStr[4:])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year")
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

// Helper function to check if date is in valid range
func isValidDateRange(requestDate, semesterStart, semesterEnd time.Time) bool {
	return (requestDate.Equal(semesterStart) || requestDate.After(semesterStart)) &&
		(requestDate.Equal(semesterEnd) || requestDate.Before(semesterEnd))
}

// GET /api/v1/schedule - Get current semester schedule or schedule for specific date
func GetSchedule(c *gin.Context) {
	requestID := uuid.New().String()

	dateParam := c.Query("date")
	allParam := c.Query("all")

	// Handle "all" parameter - return all data as ZIP
	if allParam != "" {
		handleAllDataRequest(c, requestID)
		return
	}

	// Handle specific date request
	if dateParam != "" {
		requestDate, err := parseDate(dateParam)
		fmt.Print(err)
		if err != nil {
			response := createAPIResponse(nil, []string{"Invalid date format. Use DDMMYYYY (Date/Month/Year)"}, requestID)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// Check if date is within valid range
		if !isValidDateRange(requestDate, SampleFoodSchedule.StartingDate, SampleFoodSchedule.EndingDate) {
			var errorMsg string
			if requestDate.Before(SampleFoodSchedule.StartingDate) {
				errorMsg = "ND (ERROR) - Date is before available data range"
			} else {
				errorMsg = "ND (ERROR) - Date is after available data range"
			}
			response := createAPIResponse(nil, []string{errorMsg}, requestID)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// Find the specific date in the schedule
		dateStr := requestDate.Format("2006-01-02")
		for _, week := range SampleFoodSchedule.Weeks {
			for _, day := range week.Days {
				if day.Date == dateStr {
					response := createAPIResponse(day, []string{}, requestID)
					c.JSON(http.StatusOK, response)
					return
				}
			}
		}

		response := createAPIResponse(nil, []string{"No data found for the specified date"}, requestID)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// Return current semester schedule (default behavior)
	response := createAPIResponse(SampleFoodSchedule, []string{}, requestID)
	c.JSON(http.StatusOK, response)
}

// Handle ?all parameter - return all data as ZIP
func handleAllDataRequest(c *gin.Context, requestID string) {
	// Create a buffer to write our archive to
	buf := new(bytes.Buffer)

	// Create a new zip archive
	w := zip.NewWriter(buf)

	// Add schedule data to zip
	scheduleData, err := json.MarshalIndent(SampleFoodSchedule, "", "  ")
	if err != nil {
		response := createAPIResponse(nil, []string{"Error creating archive"}, requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	f1, err := w.Create("food_schedule.json")
	if err != nil {
		response := createAPIResponse(nil, []string{"Error creating archive"}, requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	_, err = f1.Write(scheduleData)
	if err != nil {
		response := createAPIResponse(nil, []string{"Error writing to archive"}, requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Add announcements to zip
	announcementData, err := json.MarshalIndent(SampleAnnouncements, "", "  ")
	if err != nil {
		response := createAPIResponse(nil, []string{"Error creating archive"}, requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	f2, err := w.Create("announcements.json")
	if err != nil {
		response := createAPIResponse(nil, []string{"Error creating archive"}, requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	_, err = f2.Write(announcementData)
	if err != nil {
		response := createAPIResponse(nil, []string{"Error writing to archive"}, requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Close the zip writer
	err = w.Close()
	if err != nil {
		response := createAPIResponse(nil, []string{"Error finalizing archive"}, requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Set headers for file download
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"food_schedule_all_data_%s.zip\"", time.Now().Format("20060102")))
	c.Header("X-Request-ID", requestID)

	// Send the zip file
	c.Data(http.StatusOK, "application/zip", buf.Bytes())
}

// GET /api/v1/status - Get service status (Open | Closed)
func GetStatus(c *gin.Context) {
	requestID := uuid.New().String()

	response := createAPIResponse(SampleServiceStatus, []string{}, requestID)
	c.JSON(http.StatusOK, response)
}

// GET /api/v1/announcement - Get active announcements
func GetAnnouncement(c *gin.Context) {
	requestID := uuid.New().String()

	// Filter only active announcements
	var activeAnnouncements []Announcement
	currentTime := time.Now()

	for _, announcement := range SampleAnnouncements {
		if announcement.IsActive &&
			currentTime.After(announcement.StartDate) &&
			(announcement.EndDate == nil || currentTime.Before(*announcement.EndDate)) {
			activeAnnouncements = append(activeAnnouncements, announcement)
		}
	}

	response := createAPIResponse(activeAnnouncements, []string{}, requestID)
	c.JSON(http.StatusOK, response)
}

// GET /api/v1/announcement/history - Get all announcements including historical
func GetAnnouncementHistory(c *gin.Context) {
	requestID := uuid.New().String()

	response := createAPIResponse(SampleAnnouncements, []string{}, requestID)
	c.JSON(http.StatusOK, response)
}

// POST /api/v1/status - Update service status (for admin use)
func UpdateStatus(c *gin.Context) {
	requestID := uuid.New().String()

	var newStatus struct {
		Status string `json:"status" binding:"required"`
		Reason string `json:"reason,omitempty"`
	}

	if err := c.ShouldBindJSON(&newStatus); err != nil {
		response := createAPIResponse(nil, []string{"Invalid request body"}, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate status
	if !strings.EqualFold(newStatus.Status, "Open") && !strings.EqualFold(newStatus.Status, "Closed") {
		response := createAPIResponse(nil, []string{"Status must be 'Open' or 'Closed'"}, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Update the sample status (in a real app, this would update the database)
	SampleServiceStatus.Status = strings.Title(strings.ToLower(newStatus.Status))
	SampleServiceStatus.Reason = newStatus.Reason
	SampleServiceStatus.UpdatedAt = time.Now()

	response := createAPIResponse(SampleServiceStatus, []string{}, requestID)
	c.JSON(http.StatusOK, response)
}

// POST /api/v1/announcement - Create new announcement (for admin use)
func CreateAnnouncement(c *gin.Context) {
	requestID := uuid.New().String()

	var newAnnouncement struct {
		Title     string     `json:"title" binding:"required"`
		Content   string     `json:"content" binding:"required"`
		StartDate time.Time  `json:"start_date" binding:"required"`
		EndDate   *time.Time `json:"end_date,omitempty"`
	}

	if err := c.ShouldBindJSON(&newAnnouncement); err != nil {
		response := createAPIResponse(nil, []string{"Invalid request body"}, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create new announcement
	announcement := Announcement{
		ID:        len(SampleAnnouncements) + 1,
		Title:     newAnnouncement.Title,
		Content:   newAnnouncement.Content,
		StartDate: newAnnouncement.StartDate,
		EndDate:   newAnnouncement.EndDate,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Add to sample data (in a real app, this would save to database)
	SampleAnnouncements = append(SampleAnnouncements, announcement)

	response := createAPIResponse(announcement, []string{}, requestID)
	c.JSON(http.StatusCreated, response)
}

// GET /api/v1/schedule/semesters - Get all semester periods for database seeding
func GetSemesters(c *gin.Context) {
	requestID := uuid.New().String()

	semesters := []struct {
		ID           int       `json:"id"`
		Name         string    `json:"name"`
		StartingDate time.Time `json:"starting_date"`
		EndingDate   time.Time `json:"ending_date"`
		IsCurrent    bool      `json:"is_current"`
	}{
		{
			ID:           1,
			Name:         "Fall 2024",
			StartingDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
			EndingDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			IsCurrent:    true,
		},
		{
			ID:           2,
			Name:         "Spring 2025",
			StartingDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			EndingDate:   time.Date(2025, 6, 30, 0, 0, 0, 0, time.UTC),
			IsCurrent:    false,
		},
		{
			ID:           3,
			Name:         "Fall 2025",
			StartingDate: time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC),
			EndingDate:   time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
			IsCurrent:    false,
		},
	}

	response := createAPIResponse(semesters, []string{}, requestID)
	c.JSON(http.StatusOK, response)
}

/*
This project is the monolithic backend API for the OpenSourceDUTH team. Access to open data compiled and provided by the OpenSourceDUTH University Team.
API Copyright (C) 2025 OpenSourceDUTH
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
