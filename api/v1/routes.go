package v1

import "github.com/gin-gonic/gin"

// RegisterRoutes registers all v1 routes
func RegisterRoutes(rg *gin.RouterGroup) {
	// Schedule routes
	scheduleGroup := rg.Group("/schedule")
	{
		scheduleGroup.GET("", GetSchedule)            // GET /api/v1/schedule
		scheduleGroup.GET("/semesters", GetSemesters) // GET /api/v1/schedule/semesters
	}

	// Status routes
	statusGroup := rg.Group("/status")
	{
		statusGroup.GET("", GetStatus)     // GET /api/v1/status
		statusGroup.POST("", UpdateStatus) // POST /api/v1/status (admin)
	}

	// Announcement routes
	announcementGroup := rg.Group("/announcement")
	{
		announcementGroup.GET("", GetAnnouncement)                // GET /api/v1/announcement
		announcementGroup.GET("/history", GetAnnouncementHistory) // GET /api/v1/announcement/history
		announcementGroup.POST("", CreateAnnouncement)            // POST /api/v1/announcement (admin)
	}
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
