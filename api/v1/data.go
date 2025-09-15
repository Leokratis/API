package v1

import "time"

// Sample data based on the provided monthly schedule
var SampleFoodSchedule = FoodScheduleSemester{
	ID:           1,
	StartingDate: time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
	EndingDate:   time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
	IsCurrent:    true,
	CreatedAt:    time.Now(),
	UpdatedAt:    time.Now(),
	Weeks: map[string]FoodScheduleWeek{
		"week1": {
			WeekNumber: 1,
			Days: map[string]FoodScheduleDay{
				"monday": {
					Date: "2024-09-01",
					Lunch: FoodScheduleMeal{
						Option1: "Λαδερό",
						Option2: "Ψάρι",
						Option3: "Λευκό Κρέας",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Μακαρόνια Σάλτσα Ντομάτας",
						Option2: "Μοσχαράκι Μαγειρεμένο Με Αρακά",
						Option3: "Μπάμιες",
					},
				},
				"tuesday": {
					Date: "2024-09-02",
					Lunch: FoodScheduleMeal{
						Option1: "Μοσχαράκι Κοκκινιστό Πιπεριά,Μανιτάρια,Καρότο",
						Option2: "Χοιρινή Μπριζόλα",
						Option3: "Μακαρόνια Σάλτσα Ντομάτας Και Μανιτάρια",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Κρέπα Με Τυρί Γκούντα,Μπέικον,Μανιτάρια",
						Option2: "Ομελέτα Μανιτάρια Πιπεριές",
						Option3: "Πατάτα Χασάπα",
					},
				},
				"wednesday": {
					Date: "2024-09-03",
					Lunch: FoodScheduleMeal{
						Option1: "Παστίτσιο",
						Option2: "Όσπρια",
						Option3: "Κοτόπουλο Σούπα Αυγολέμονο",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Ψάρι",
						Option2: "Μακαρόνια",
						Option3: "Αρακάς",
					},
				},
				"thursday": {
					Date: "2024-09-04",
					Lunch: FoodScheduleMeal{
						Option1: "Λευκό Κρέας Ψητό",
						Option2: "Χοιρινή Τηγανιά Φούρνου",
						Option3: "Λαδερό Με Πατάτες",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Ομελέτα Τυρί Πατάτες Τηγανιτές",
						Option2: "Μελιτζάνες Ραγού",
						Option3: "Κοτομπουκιές",
					},
				},
				"friday": {
					Date: "2024-09-05",
					Lunch: FoodScheduleMeal{
						Option1: "Όσπρια",
						Option2: "Μακαρόνια Με Κιμά",
						Option3: "Γεμιστά Με Κιμά Ή Λαδερό",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Σουτζουκάκια",
						Option2: "Λαδερό",
						Option3: "Λουκάνικα Με Κοφτό Μακαρονάκι",
					},
				},
				"saturday": {
					Date: "2024-09-06",
					Lunch: FoodScheduleMeal{
						Option1: "Μοσχαράκι",
						Option2: "Καλαμαράκια",
						Option3: "Χοιρινό",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Πίτσα",
						Option2: "Όσπριο",
						Option3: "Σουπιές Σπανάκι",
					},
				},
				"sunday": {
					Date: "2024-09-07",
					Lunch: FoodScheduleMeal{
						Option1: "Λευκό Κρέας Ψητό",
						Option2: "Μουσακάς Μελιτζάνας",
						Option3: "Μπριάμ",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Μπιφτέκι",
						Option2: "Σπανακόρυζο",
						Option3: "Λουκάνικο",
					},
				},
			},
		},
		"week2": {
			WeekNumber: 2,
			Days: map[string]FoodScheduleDay{
				"monday": {
					Date: "2024-09-08",
					Lunch: FoodScheduleMeal{
						Option1: "Μοσχαράκι Κοκκινιστό Κολοκυθάκια",
						Option2: "Λαδερό",
						Option3: "Ψάρι",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Λευκό Κρέας Ογκρατέν",
						Option2: "Όσπριο",
						Option3: "Λουκάνικο Ταβέρνας",
					},
				},
				"tuesday": {
					Date: "2024-09-09",
					Lunch: FoodScheduleMeal{
						Option1: "Σουφλέ Λαζάνια",
						Option2: "Όσπρια",
						Option3: "Κολοκυθιά Γεμιστά Αυγολέμονο",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Ψάρι",
						Option2: "Λαδερό",
						Option3: "Τορτίγια Αλλαντικών",
					},
				},
				"wednesday": {
					Date: "2024-09-10",
					Lunch: FoodScheduleMeal{
						Option1: "Όσπρια",
						Option2: "Μακαρόνια Με Κιμά",
						Option3: "Λευκό Κρέας Ψητό",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Κοτομπουκιές",
						Option2: "Καλαμαράκι Φρικασέ",
						Option3: "Λαδερό",
					},
				},
				"thursday": {
					Date: "2024-09-11",
					Lunch: FoodScheduleMeal{
						Option1: "Χοιρινό Λεμονάτο",
						Option2: "Ριζότο Θαλασσινών",
						Option3: "Όσπρια",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Λαδερό",
						Option2: "Λευκό Κρέας Σουφλέ",
						Option3: "Τριβελάκι",
					},
				},
				"friday": {
					Date: "2024-09-12",
					Lunch: FoodScheduleMeal{
						Option1: "Χοιρινό Ψητό",
						Option2: "Ριζότο Με Μύδια",
						Option3: "Αρακάς",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Πίτσα Ζαμπόν, Μανιτάρια, Τυρί",
						Option2: "Μακαρόνια Σάλτσα Τομάτας",
						Option3: "Κοτόπουλο Πανέ",
					},
				},
				"saturday": {
					Date: "2024-09-13",
					Lunch: FoodScheduleMeal{
						Option1: "Κοτόπουλο Λεμονάτο",
						Option2: "Μοσχάρι Στιφάδο",
						Option3: "Χορτόσουπα",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Μπιφτέκια",
						Option2: "Όσπρια",
						Option3: "Μπάμιες",
					},
				},
				"sunday": {
					Date: "2024-09-14",
					Lunch: FoodScheduleMeal{
						Option1: "Πεταλούδες Με Κιμά",
						Option2: "Χοιρινή Τηγανιά",
						Option3: "Μελιτζάνα Ιμάμ",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Σουτζουκάκια",
						Option2: "Καλαμαράκι Κοκκινιστό",
						Option3: "Πίτσα",
					},
				},
			},
		},
		"week3": {
			WeekNumber: 3,
			Days: map[string]FoodScheduleDay{
				"monday": {
					Date: "2024-09-15",
					Lunch: FoodScheduleMeal{
						Option1: "Μοσχάρι",
						Option2: "Όσπριο",
						Option3: "Κοτόπουλο",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Λουκάνικα",
						Option2: "Λαδερό",
						Option3: "Καρμπονάρα",
					},
				},
				"tuesday": {
					Date: "2024-09-16",
					Lunch: FoodScheduleMeal{
						Option1: "Παστίτσιο",
						Option2: "Αγκινάρες Αλά Πολίτα",
						Option3: "Χοιρινό Ψητό",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Μπριάμ",
						Option2: "Ομελέτα",
						Option3: "Μοσχαράκι Σούπα Αυγολέμονο",
					},
				},
				"wednesday": {
					Date: "2024-09-17",
					Lunch: FoodScheduleMeal{
						Option1: "Ψάρι",
						Option2: "Χοιρινό Κοκκινιστό",
						Option3: "Λαδερό",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Μακαρόνια Ογκρατέν",
						Option2: "Καλαμαράκια Ρύζι",
						Option3: "Πατάτες Με Τυρί Και Λαχανικά",
					},
				},
				"thursday": {
					Date: "2024-09-18",
					Lunch: FoodScheduleMeal{
						Option1: "Γιουβαρλάκια",
						Option2: "Λαδερό",
						Option3: "Σπετσοφάι",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Ομελέτα Με Πατάτες",
						Option2: "Κοτόσουπα",
						Option3: "Κοτομπουκιές",
					},
				},
				"friday": {
					Date: "2024-09-19",
					Lunch: FoodScheduleMeal{
						Option1: "Όσπρια",
						Option2: "Κοτόπουλο Χυλοπίτες",
						Option3: "Μοσχαράκι Βίδες Λαχανικά",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Μακαρόνια Αματριτσιάνα",
						Option2: "Τυρόπιτα",
						Option3: "Φασολάκια",
					},
				},
				"saturday": {
					Date: "2024-09-20",
					Lunch: FoodScheduleMeal{
						Option1: "Μπριζόλα",
						Option2: "Ψάρι",
						Option3: "Αρακάς",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Λαχανόρυζο",
						Option2: "Πανσέτα Χοιρινή",
						Option3: "Καλαμάρια Με Ρύζι",
					},
				},
				"sunday": {
					Date: "2024-09-21",
					Lunch: FoodScheduleMeal{
						Option1: "Μπιφτέκι",
						Option2: "Λευκό Κρέας",
						Option3: "Μπριάμ",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Όσπριο",
						Option2: "Σουτζουκάκια",
						Option3: "Ομελέτα",
					},
				},
			},
		},
		"week4": {
			WeekNumber: 4,
			Days: map[string]FoodScheduleDay{
				"monday": {
					Date: "2024-09-22",
					Lunch: FoodScheduleMeal{
						Option1: "Μοσχάρι Κοκκινιστό",
						Option2: "Λουκάνικο Χωριάτικο",
						Option3: "Μακαρόνια Σάλτσα Τομάτα Και Μανιτάρια",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Ψαροκροκέτα",
						Option2: "Λαδερό",
						Option3: "Πατάτα Χασάπα",
					},
				},
				"tuesday": {
					Date: "2024-09-23",
					Lunch: FoodScheduleMeal{
						Option1: "Κοτόπουλο Λεμονάτο",
						Option2: "Όσπρια",
						Option3: "Κοτομπουκιές",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Σπετσοφάι",
						Option2: "Σουφλέ Κοτόπουλο",
						Option3: "Μπάμιες",
					},
				},
				"wednesday": {
					Date: "2024-09-24",
					Lunch: FoodScheduleMeal{
						Option1: "Λαδερό",
						Option2: "Σουτζουκάκια",
						Option3: "Σουφλέ Τυριών",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Ψάρι",
						Option2: "Πίτσα Μαργαρίτα",
						Option3: "Αρακάς",
					},
				},
				"thursday": {
					Date: "2024-09-25",
					Lunch: FoodScheduleMeal{
						Option1: "Ψάρι",
						Option2: "Λευκό Κρέας Ψητό",
						Option3: "Λαδερό Με Πατάτες",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Όσπρια",
						Option2: "Μακαρόνια Αλά Κρεμ Κοτόπουλο",
						Option3: "Ομελέτα",
					},
				},
				"friday": {
					Date: "2024-09-26",
					Lunch: FoodScheduleMeal{
						Option1: "Φασολάδα",
						Option2: "Ψαροκροκέτα",
						Option3: "Γιουβαρλάκια",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Μακαρόνια Με Κιμά",
						Option2: "Λουκάνικα",
						Option3: "Λαδερό",
					},
				},
				"saturday": {
					Date: "2024-09-27",
					Lunch: FoodScheduleMeal{
						Option1: "Κοτόπουλο Κοκκινιστό",
						Option2: "Μοσχαράκι Κοκκινιστό",
						Option3: "Κριθαρότο Μανιτάρια",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Χοιρινό Λεμονάτο",
						Option2: "Μπριάμ",
						Option3: "Λαχανόρυζο",
					},
				},
				"sunday": {
					Date: "2024-09-28",
					Lunch: FoodScheduleMeal{
						Option1: "Σουφλέ Λαζάνια",
						Option2: "Μελιτζάνες Ραγού",
						Option3: "Χοιρινό Ψητό",
					},
					Dinner: FoodScheduleMeal{
						Option1: "Ψάρι",
						Option2: "Λαδερό",
						Option3: "Κοτόπουλο Λεμονάτο",
					},
				},
			},
		},
	},
}

// Sample status data
var SampleServiceStatus = FoodServiceStatus{
	ID:        1,
	Status:    "Open",
	UpdatedAt: time.Now(),
	Reason:    "",
}

// Sample announcements
var SampleAnnouncements = []Announcement{
	{
		ID:        1,
		Title:     "Καλώς ήρθατε στη νέα περίοδο",
		Content:   "Καλωσορίζουμε όλους τους φοιτητές στην καινούρια περίοδο. Το εστιατόριο θα λειτουργεί κανονικά από Δευτέρα.",
		StartDate: time.Date(2024, 9, 1, 9, 0, 0, 0, time.UTC),
		EndDate:   nil,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		Title:     "Αλλαγές στο πρόγραμμα",
		Content:   "Τη Δευτέρα 18/10 το εστιατόριο θα κλείσει στις 20:00 λόγω συντήρησης.",
		StartDate: time.Date(2024, 10, 15, 10, 0, 0, 0, time.UTC),
		EndDate:   &[]time.Time{time.Date(2024, 10, 18, 22, 0, 0, 0, time.UTC)}[0],
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
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
