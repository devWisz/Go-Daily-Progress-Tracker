# Go-Daily-Progress-Tracker
Go Daily Progress Tracker is a program that allows a user to track their daily progress in a interactive way in a offline mode without risking your personal info only useful features are included that helps you to keep your progress tracked

## Why I Built This

Most productivity apps are filled with distractions, unnecessary features, or too much fluff.

I wanted something simple.

This is not just a program.

This is my daily learning tracker. Built for real use. Built for myself.  
A focused, clutter-free way to journal what I actually did each day.  
And now, shared with the world.

---

## Features

- Track daily learning progress with topic details
- Automatically convert minutes to hours+minutes
- Reflect on your experience with meaningful input validation
- Log debugging efforts (if any)
- Rate your day out of 5
- Set weekly or monthly learning goals
- Saves logs in both JSON and TXT formats
- Auto-day tracking (Day 1, Day 2, Day 3, ...)

Golang Daily Tracker
Logging Day 2

What Golang topics did you learn today?

Pointers
Structs


What did you do in Pointers today?

Practiced examples, understood references

Which resources did you use today?

Go Docs
YouTube

Time spent learning (in minutes): 90

Was your day productive? (yes/no): yes

Rate your day out of 5: 5

How was the overall experience today?
Learned a lot today, debugging went smooth

Did you get errors or spend time in debugging? (yes/no): yes
Describe your debugging experience:
Faced nil pointer panic, fixed it using struct checks

Do you want to set a weekly or monthly goal? (yes/no): yes
Type of goal (weekly/monthly): weekly

Write your goal:

Complete Go routines and write 3 CLI apps
Goal created successfully!

Do you want to save this entry? (yes/no): yes
Your progress has been saved successfully.
Developed by Sarjak Khanal — Thanks for using Golang Daily Tracker!


## Output Format

### JSON (`golang_log.json`)


```json-----
{
  "day": "Day 2",
  "date_time": "June 20, 2025 | 3:04 PM",
  "topics": [
    {
      "name": "Pointers",
      "what_did_you_do": "Practiced examples, understood references"
    }
  ],
  "resources_used": ["Go Docs", "YouTube"],
  "time_spent": "1 hr(s) 30 mins",
  "was_productive": "Yes",
  "day_rating": 4,
  "experience_summary": "Learned a lot today, debugging went smooth",
  "debugging": {
    "faced_debugging": "Yes",
    "debugging_experience": "Faced nil pointer panic, fixed it using struct checks"
  },
  "goal": {
    "type": "Weekly",
    "description": "Complete Go routines and write 3 CLI apps",
    "created_at": "June 20, 2025 | 3:04 PM"
  }
}


TXT (golang_log.txt)

=== Day 2 (June 20, 2025 | 3:04 PM) ===
Topic: Pointers
Action: Practiced examples, understood references
Resources: Go Docs, YouTube
Time Spent: 1 hr(s) 30 mins
Productive: Yes
Rating: 4/5
Experience: Learned a lot today, debugging went smooth
Debugging: Yes
Debugging Experience: Faced nil pointer panic, fixed it using struct checks
Goal Type: Weekly
Goal: Complete Go routines and write 3 CLI apps
Set on: June 20, 2025 | 3:04 PM

Installation
Make sure you have Go installed: https://golang.org/dl/

Clone this repository and run the program:

bash
Copy
Edit
git clone https://github.com/your-username/golang-daily-tracker.git
cd golang-daily-tracker
go run main.go


Use Cases
Keep a clean log of your programming or learning journey
Track your debugging experiences for future reference
Stay accountable to weekly/monthly goals
Maintain discipline without distractions

Contributing

If you have ideas to improve it:
Star the repo
Fork it
Submit a pull request
Create an issue

Let’s make CLI productivity better, together.

Developer
Sarjak Khanal
Coder, focused learner, builder of personal tools.

LinkedIn:https://www.linkedin.com/in/sarjak-khanal/
GitHub: https:https://github.com/devWisz/

