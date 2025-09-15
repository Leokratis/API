package main

import (
	"net/http"

	// v0 "api/api/v0"
	v1 "api/api/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/", homePage)

	// // v0 routes
	// v0Group := router.Group("/api/v0")
	// {
	// 	v0.Food_Schedule_Routes(v0Group)
	// }

	// v1 routes
	v1Group := router.Group("/api/v1")
	{
		v1.RegisterRoutes(v1Group)
	}

	router.Run(":9237")
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is the Home Page for the OpenSourceDUTH API, welcome! \nGo to https://opensource.cs.duth.gr/docs/getting-started/ for API documentation.\n\nAvailable API versions:\n- v0: /api/v0/\n- v1: /api/v1/ (Food Schedule API)")
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
