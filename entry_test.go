package schedule

import "testing"

func TestEveryMinuteReturnsOneStarFollowing(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.EveryMinute()

	exp := defaultExp
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestEveryFiveMinutesReturnsStarSlashFiveOnPositionOne(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.EveryFiveMinutes()

	exp := "*/5 * * * * *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestEveryTenMinutesReturnsStarSlashTenOnPositionOne(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.EveryTenMinutes()

	exp := "*/10 * * * * *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestEveryFifteenMinutesReturnsStarSlashFifteenOnPositionOne(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.EveryFifteenMinutes()

	exp := "*/15 * * * * *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestEveryThirtyMinutesReturnsZeroCommaThirtyOnPositionOne(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.EveryThirtyMinutes()

	exp := "0,30 * * * * *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}

func TestHourlyReturnsZeroInPositionOne(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Hourly()

	exp := "0 * * * * *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}

func TestDailyReturnsZeroInPositionOneAndTwo(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Daily()

	exp := "0 0 * * * *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}

func TestDailyAtReturnsAppropriateExpression(t *testing.T) {
	testCases := []struct {
		desc      string
		givenSpec string
		expected  string
	}{
		{
			desc:      "15:00 returns '00 15 * * * *",
			givenSpec: "15:00",
			expected:  "00 15 * * * *",
		},
		{
			desc:      "01:23 returns '23 01 * * * *",
			givenSpec: "01:23",
			expected:  "23 01 * * * *",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			e := &Entry{expression: defaultExp}
			e.DailyAt(tC.givenSpec)

			if e.expression != tC.expected {
				t.Errorf("expected '%s', got '%s'", tC.expected, e.expression)
			}
		})
	}
}

func TestWeeklyReturnsZeroInPositionOneAndTwoAndFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Weekly()

	exp := "0 0 * * 0 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestWeekdaysReturnsCommaJoinedOneToFiveOnPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Weekdays()

	exp := "* * * * 1,2,3,4,5 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}

func TestDaysReturnsCommaJoinedDaysOnPositionFive(t *testing.T) {
	testCases := []struct {
		desc     string
		days     []cronDay
		expected string
	}{
		{
			desc:     "Monday(1), Tuesday(2), Wednesday(3) return * * * * 1,2,3 *",
			days:     []cronDay{Monday, Tuesday, Wednesday},
			expected: "* * * * 1,2,3 *",
		},
		{
			desc:     "Tuesday(2), Thursday(4), Sunday(0) return * * * * 2,3,0 *",
			days:     []cronDay{Tuesday, Thursday, Sunday},
			expected: "* * * * 2,4,0 *",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			e := &Entry{expression: defaultExp}
			e.Days(tC.days...)
			if e.expression != tC.expected {
				t.Errorf("expected '%s', got '%s'", tC.expected, e.expression)
			}
		})
	}
}

func TestMondaysReturnsOneInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Mondays()
	exp := "* * * * 1 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestTuesdaysReturnsTwoInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Tuesdays()
	exp := "* * * * 2 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestWednesdaysReturnsThreeInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Wednesdays()
	exp := "* * * * 3 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestThurdaysReturnsFourInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Thursdays()
	exp := "* * * * 4 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestFridaysReturnsFiveInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Fridays()
	exp := "* * * * 5 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestSaturdaysReturnsSixInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Saturdays()
	exp := "* * * * 6 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}
func TestSundaysReturnsZeroInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Sundays()
	exp := "* * * * 0 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}

func TestMondaysHourlyReturnsZeroInPositionOneAndOneInPositionFive(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.Mondays().Hourly()

	exp := "0 * * * 1 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}

func TestCombinedSpecsDontOverwriteEachOther(t *testing.T) {
	e := &Entry{expression: defaultExp}
	e.EveryFiveMinutes()
	e.Mondays()

	exp := "*/5 * * * 1 *"
	if e.expression != exp {
		t.Errorf("expected '%s', got '%s'", exp, e.expression)
	}
}

func TestStringer(t *testing.T) {
	e := &Entry{expression: defaultExp}

	if e.String() != defaultExp {
		t.Errorf("expexted '%s', got '%s", defaultExp, e.String())
	}
}

func TestDailyAtPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("DailyAt should panic if invalid segment is given")
		}
	}()
	e := &Entry{expression: defaultExp}
	e.DailyAt("foooo")
}
