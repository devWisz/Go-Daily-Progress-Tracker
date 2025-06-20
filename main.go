package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type TopicDetail struct {
	Name   string `json:"name"`
	Action string `json:"what_did_you_do"`
}

type DebuggingDetail struct {
	FacedDebugging      string `json:"faced_debugging"`
	DebuggingExperience string `json:"debugging_experience,omitempty"`
}

type Goal struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type DSLog struct {
	Day           string          `json:"day"`
	DateTime      string          `json:"date_time"`
	Topics        []TopicDetail   `json:"topics"`
	Resources     []string        `json:"resources_used"`
	TimeSpent     string          `json:"time_spent"`
	ProductiveDay string          `json:"was_productive"`
	Rating        int             `json:"day_rating"`
	Experience    string          `json:"experience_summary"`
	Debugging     DebuggingDetail `json:"debugging"`
	Goal          *Goal           `json:"goal,omitempty"`
}

func readExistingLogs(filename string) []DSLog {
	var logs []DSLog
	data, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(data, &logs)
	}
	return logs
}

func calculateDayNumber(logs []DSLog) int {
	if len(logs) == 0 {
		return 1
	}
	layout := "January 2, 2006 | 3:04 PM"
	firstDate, _ := time.Parse(layout, logs[0].DateTime)
	return int(time.Since(firstDate).Hours()/24) + 1
}

func getListInput(prompt string) []string {
	fmt.Println(prompt + " (type 'done' to finish):")
	var list []string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("- ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "done" {
			break
		}
		if input != "" {
			list = append(list, input)
		}
	}
	return list
}

func getStringInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func getValidatedExperience(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println(" Input cannot be empty. Try again.")
			continue
		}

		// Check if input is purely numeric
		if _, err := strconv.Atoi(input); err == nil {
			fmt.Println("❌ Only numbers are not allowed. Write a meaningful experience.")
			continue
		}

		// Check if it has at least one letter
		hasLetter := false
		for _, ch := range input {
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				hasLetter = true
				break
			}
		}

		if !hasLetter {
			fmt.Println(" Input must contain at least some text.")
			continue
		}

		return input
	}
}

func getYesNoInput(prompt string) string {
	for {
		answer := strings.ToLower(getStringInput(prompt + " (yes/no): "))
		if answer == "yes" {
			return "Yes"
		} else if answer == "no" {
			return "No"
		}
		fmt.Println(" Please enter 'yes' or 'no'")
	}
}

func getIntInput(prompt string, min int, max int) int {
	for {
		fmt.Print(prompt)
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil || num < min || num > max {
			fmt.Printf(" Please enter a number between %d and %d.\n", min, max)
			continue
		}
		return num
	}
}

func formatTimeSpent(minutes int) string {
	if minutes < 60 {
		return fmt.Sprintf("%d mins", minutes)
	}
	hrs := minutes / 60
	mins := minutes % 60
	if mins == 0 {
		return fmt.Sprintf("%d hr(s)", hrs)
	}
	return fmt.Sprintf("%d hr(s) %d mins", hrs, mins)
}

func saveLog(entry DSLog, jsonFile, txtFile string) {
	// Save JSON
	logs := readExistingLogs(jsonFile)
	logs = append(logs, entry)
	data, err := json.MarshalIndent(logs, "", "  ")
	if err == nil {
		os.WriteFile(jsonFile, data, 0644)
	}

	// Save TXT
	f, err := os.OpenFile(txtFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString(fmt.Sprintf("=== %s (%s) ===\n", entry.Day, entry.DateTime))
		for _, t := range entry.Topics {
			f.WriteString(fmt.Sprintf("Topic: %s\nAction: %s\n", t.Name, t.Action))
		}
		f.WriteString(fmt.Sprintf("Resources: %s\n", strings.Join(entry.Resources, ", ")))
		f.WriteString(fmt.Sprintf("Time Spent: %s\n", entry.TimeSpent))
		f.WriteString(fmt.Sprintf("Productive: %s\n", entry.ProductiveDay))
		f.WriteString(fmt.Sprintf("Rating: %d/5\n", entry.Rating))
		f.WriteString(fmt.Sprintf("Experience: %s\n", entry.Experience))
		f.WriteString(fmt.Sprintf("Debugging: %s\n", entry.Debugging.FacedDebugging))
		if entry.Debugging.FacedDebugging == "Yes" {
			f.WriteString(fmt.Sprintf("Debugging Experience: %s\n", entry.Debugging.DebuggingExperience))
		}
		if entry.Goal != nil {
			f.WriteString(fmt.Sprintf("Goal Type: %s\nGoal: %s\nSet on: %s\n", entry.Goal.Type, entry.Goal.Description, entry.Goal.CreatedAt))
		}
		f.WriteString("\n")
	}

	fmt.Println("\n✅ Your progress has been saved successfully to both JSON and TXT!")
}

func askGoal() *Goal {
	setGoal := getYesNoInput(" Do you want to set a weekly or monthly goal?")
	if setGoal == "Yes" {
		for {
			typeChoice := strings.ToLower(getStringInput(" Type of goal (weekly/monthly): "))
			if typeChoice == "weekly" || typeChoice == "monthly" {
				goalDesc := getValidatedExperience("📝 Write your goal (e.g., Finish Go routines, Practice 5 problems/day): ")
				fmt.Println(" Goal created successfully!")
				return &Goal{
					Type:        strings.Title(typeChoice),
					Description: goalDesc,
					CreatedAt:   time.Now().Format("January 2, 2006 | 3:04 PM"),
				}
			} else {
				fmt.Println("❌ Please enter either 'weekly' or 'monthly'")
			}
		}
	}
	return nil
}

func main() {
	jsonFile := "golang_log.json"
	txtFile := "golang_log.txt"
	existingLogs := readExistingLogs(jsonFile)
	dayNumber := calculateDayNumber(existingLogs)
	dayLabel := fmt.Sprintf("Day %d", dayNumber)

	fmt.Println(" Golang Daily Tracker")
	fmt.Printf("Logging %s\n", dayLabel)

	topics := getListInput("What Golang topics did you learn today? (e.g., Struct, Maps, Pointers)")
	var topicDetails []TopicDetail
	for _, t := range topics {
		action := getStringInput(fmt.Sprintf(" What did you do in %s today?\n> ", t))
		topicDetails = append(topicDetails, TopicDetail{Name: t, Action: action})
	}

	resources := getListInput(" Which resources did you use today? (Docs, YouTube, Blogs, etc.)")
	timeSpent := getIntInput("  Time spent learning (in minutes): ", 1, 1440)
	productive := getYesNoInput(" Was your day productive?")
	rating := getIntInput(" Rate your day out of 5: ", 1, 5)
	experience := getValidatedExperience(" How was the overall experience today?\n> ")

	debug := getYesNoInput(" Did you get errors or spend time in debugging?")
	var debugExp string
	if debug == "Yes" {
		debugExp = getValidatedExperience(" Describe your debugging experience:\n> ")
	} else {
		fmt.Println(" Nice! You did a good job today!")
	}

	goal := askGoal()

	entry := DSLog{
		Day:           dayLabel,
		DateTime:      time.Now().Format("January 2, 2006 | 3:04 PM"),
		Topics:        topicDetails,
		Resources:     resources,
		TimeSpent:     formatTimeSpent(timeSpent),
		ProductiveDay: productive,
		Rating:        rating,
		Experience:    experience,
		Debugging: DebuggingDetail{
			FacedDebugging:      debug,
			DebuggingExperience: debugExp,
		},
		Goal: goal,
	}

	confirm := getYesNoInput(" Do you want to save this entry?")
	if confirm == "Yes" {
		saveLog(entry, jsonFile, txtFile)
	} else {
		fmt.Println("❌ Entry was not saved.")
	}

	fmt.Println("\n Developed by Sarjak Khanal — Thanks for using Golang Daily Tracker!")
}
