package stocks

import "time"

func StockMarketIsOpen() bool {
	now := time.Now()

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		loc = time.UTC
	}
	estNow := now.In(loc)

	if estNow.Weekday() == time.Saturday || estNow.Weekday() == time.Sunday {
		return false
	}

	marketOpen := time.Date(estNow.Year(), estNow.Month(), estNow.Day(), 9, 30, 0, 0, loc)
	marketClose := time.Date(estNow.Year(), estNow.Month(), estNow.Day(), 16, 0, 0, 0, loc)
	if estNow.Before(marketOpen) || estNow.After(marketClose) {
		return false
	}

	if isHoliday(estNow) {
		return false
	}

	return true
}

func isHoliday(date time.Time) bool {
	year := date.Year()
	holidays := []time.Time{
		newYorkTime(year, time.January, 1),   // New Year's Day
		martinLutherKingDay(year),            // Martin Luther King Jr. Day
		presidentsDay(year),                  // Presidents Day
		newYorkTime(year, time.May, 31),      // Memorial Day (last Monday in May, approximation)
		newYorkTime(year, time.July, 4),      // Independence Day
		laborDay(year),                       // Labor Day
		thanksgivingDay(year),                // Thanksgiving Day
		newYorkTime(year, time.December, 25), // Christmas Day
	}

	for _, holiday := range holidays {
		if date.Month() == holiday.Month() && date.Day() == holiday.Day() {
			return true
		}
	}

	return false
}

func martinLutherKingDay(year int) time.Time {
	return nthWeekday(year, time.January, time.Monday, 3)
}

func presidentsDay(year int) time.Time {
	return nthWeekday(year, time.February, time.Monday, 3)
}

func laborDay(year int) time.Time {
	return nthWeekday(year, time.September, time.Monday, 1)
}

func thanksgivingDay(year int) time.Time {
	return nthWeekday(year, time.November, time.Thursday, 4)
}

func nthWeekday(year int, month time.Month, weekday time.Weekday, n int) time.Time {
	t := newYorkTime(year, month, 1)
	for t.Weekday() != weekday {
		t = t.AddDate(0, 0, 1)
	}
	return t.AddDate(0, 0, (n-1)*7)
}

func newYorkTime(year int, month time.Month, day int) time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		loc = time.UTC
	}
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}
