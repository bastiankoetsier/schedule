package schedule

import (
	"strings"
)

// Cron expression positions (https://en.wikipedia.org/wiki/Cron#CRON_expression)
// 1 Minutes
// 2 Hours
// 3 Day of the month
// 4 Month
// 5 Day of the week
// 6 Year

const (
	// Sunday is representation of cron's 0 for Sundays
	Sunday = cronDay("0")
	// Monday is representation of cron's 1 for Mondays
	Monday = cronDay("1")
	// Tuesday is representation of cron's 2 for Tuesdays
	Tuesday = cronDay("2")
	// Wednesday is representation of cron's 3 for Wednesdays
	Wednesday = cronDay("3")
	// Thursday is representation of cron's 4 for Thursdays
	Thursday = cronDay("4")
	// Friday is representation of cron's 5 for Fridays
	Friday = cronDay("5")
	// Saturday is representation of cron's 6 for Saturdays
	Saturday = cronDay("6")
)

// cronDay is an unexported type to represent a cron's specification of a weekday
type cronDay string

// Entry is the holding type that has the crons specification for when the job has to run
// This is stored in the unexported field expression.
// Command keeps track of the actual func to run
type Entry struct {
	expression string
	Command    Runner
}

func (e *Entry) String() string {
	return e.expression
}

// EveryMinute specifies that this entry should be executed every single minute
func (e *Entry) EveryMinute() *Entry {
	return e.spliceIntoPosition(1, "*")
}

// EveryFiveMinutes specifies that this entry should be executed every five minutes
func (e *Entry) EveryFiveMinutes() *Entry {
	return e.spliceIntoPosition(1, "*/5")
}

// EveryTenMinutes specifies that this entry should be executed every ten minutes
func (e *Entry) EveryTenMinutes() *Entry {
	return e.spliceIntoPosition(1, "*/10")
}

// EveryFifteenMinutes specifies that this entry should be executed every fifteen minutes
func (e *Entry) EveryFifteenMinutes() *Entry {
	return e.spliceIntoPosition(1, "*/15")
}

// EveryThirtyMinutes specifies that this entry should be executed every thirty minutes
func (e *Entry) EveryThirtyMinutes() *Entry {
	return e.spliceIntoPosition(1, "0,30")
}

// Hourly specifies that this entry should be executed every hour at the minute 0
func (e *Entry) Hourly() *Entry {
	return e.spliceIntoPosition(1, "0")
}

// Daily specifies that this entry should be executed at hour 0 and minute 0
func (e *Entry) Daily() *Entry {
	return e.spliceIntoPosition(1, "0").spliceIntoPosition(2, "0")
}

// DailyAt gives you more flexibility in specifying exactly at what time the Runner should be executed
// You need to specify the spec like "15:23". This func panics if an invalid argument is given!
func (e *Entry) DailyAt(spec string) *Entry {
	segments := strings.Split(spec, ":")
	if len(segments) != 2 {
		panic("invalid specification of an entry")
	}
	return e.spliceIntoPosition(2, segments[0]).spliceIntoPosition(1, segments[1])
}

// Weekly runs every Friday at hour 0, minute 0 and day 5
func (e *Entry) Weekly() *Entry {
	return e.spliceIntoPosition(1, "0").spliceIntoPosition(2, "0").spliceIntoPosition(5, "0")
}

// Weekdays runs every MON-FRI at hour 0 and minute 0
func (e *Entry) Weekdays() *Entry {
	return e.Days(Monday, Tuesday, Wednesday, Thursday, Friday)
}

// Sundays runs every Sunday every minute (if not specified further)
func (e *Entry) Sundays() *Entry {
	return e.Days(Sunday)
}

// Mondays runs every Monday every minute (if not specified further)
func (e *Entry) Mondays() *Entry {
	return e.Days(Monday)
}

// Tuesdays runs every Tuesday every minute (if not specified further)
func (e *Entry) Tuesdays() *Entry {
	return e.Days(Tuesday)
}

// Wednesdays runs every Wednesday every minute (if not specified further)
func (e *Entry) Wednesdays() *Entry {
	return e.Days(Wednesday)
}

// Thursdays runs every Thursday every minute (if not specified further)
func (e *Entry) Thursdays() *Entry {
	return e.Days(Thursday)
}

// Fridays runs every Friday every minute (if not specified further)
func (e *Entry) Fridays() *Entry {
	return e.Days(Friday)
}

// Saturdays runs every Saturday every minute (if not specified further)
func (e *Entry) Saturdays() *Entry {
	return e.Days(Saturday)
}

// Days specifies the days the Runner should be executed on.
// This follows the specification for cron jobs, meaning it goes from 0–6 or SUN–SAT
func (e *Entry) Days(days ...cronDay) *Entry {
	s := []string{}
	for _, d := range days {
		s = append(s, string(d))
	}
	return e.spliceIntoPosition(5, strings.Join(s, ","))
}

func (e *Entry) spliceIntoPosition(position int, value string) *Entry {
	segments := strings.Fields(e.expression)
	segments[position-1] = value

	e.expression = strings.Join(segments, " ")
	return e
}
