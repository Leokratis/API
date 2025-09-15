package v1

import "time"

// Standard API Response Format
type APIResponse struct {
	Data     interface{} `json:"data"`
	Errors   []string    `json:"errors"`
	Metadata Metadata    `json:"metadata"`
}

// Metadata for API responses
type Metadata struct {
	Timestamp   time.Time `json:"timestamp"`
	Version     string    `json:"version"`
	RequestID   string    `json:"request_id,omitempty"`
	TotalCount  int       `json:"total_count,omitempty"`
	CurrentPage int       `json:"current_page,omitempty"`
	PerPage     int       `json:"per_page,omitempty"`
}

// Food Schedule Models
type FoodScheduleSemester struct {
	ID           int                         `json:"id" gorm:"primary_key"`
	StartingDate time.Time                   `json:"starting_date" gorm:"not null"`
	EndingDate   time.Time                   `json:"ending_date" gorm:"not null"`
	IsCurrent    bool                        `json:"is_current" gorm:"default:false"`
	CreatedAt    time.Time                   `json:"created_at"`
	UpdatedAt    time.Time                   `json:"updated_at"`
	Weeks        map[string]FoodScheduleWeek `json:"weeks" gorm:"-"`
}

type FoodScheduleWeek struct {
	WeekNumber int                        `json:"week_number"`
	Days       map[string]FoodScheduleDay `json:"days"`
}

type FoodScheduleDay struct {
	Date   string           `json:"date"`
	Lunch  FoodScheduleMeal `json:"lunch"`
	Dinner FoodScheduleMeal `json:"dinner"`
}

type FoodScheduleMeal struct {
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
}

// Status Models
type FoodServiceStatus struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Status    string    `json:"status"` // "Open" or "Closed"
	UpdatedAt time.Time `json:"updated_at"`
	Reason    string    `json:"reason,omitempty"`
}

// Announcement Models
type Announcement struct {
	ID        int        `json:"id" gorm:"primary_key"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Historical Data Request Model
type HistoricalDataRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Format    string    `json:"format"` // "json" or "zip"
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
