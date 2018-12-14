package run

type AthleteCredential struct {
	ID          uint32
	AccessToken string
}

type Athlete struct {
	ID                   uint32 `json:"id"`
	UserName             string `json:"username"`
	FirstName            string `json:"firstname"`
	LastName             string `json:"lastname"`
	City                 string `json:"city"`
	State                string `json:"state"`
	Country              string `json:"country"`
	Gender               string `json:"sex"`
	ProfilePicture       string `json:"profile"`
	ProfilePictureMedium string `json:"profile_medium"`
	Email                string `json:"email"`
	Stats                `json:"stats"`
}

// Stats is running stats
type Stats struct {
	RecentRun          RecentStats `json:"recent"`
	RecentRunTotals    RunStats    `json:"recent_run_totals"`
	AllRunTotals       RunStats    `json:"all_run_totals"`
	ThisMonthRunTotals RunStats    `json:"this_month_run_totals"`
	ThisYearRunTotals  RunStats    `json:"this_year_run_totals"`
}

// RecentStats is stats of recent run
type RecentStats struct {
	Distance       float64 `json:"distance"`
	ElapsedTime    uint32  `json:"elapsed_time"`
	MovingTime     uint32  `json:"moving_time"`
	Title          string  `json:"title"`
	StartDate      string  `json:"start_date"`
	TimeZoneOffset float64 `json:"utc_offset"`
}

// RunStats is detailed of stats
type RunStats struct {
	Count         uint32  `json:"count"`
	Distance      float64 `json:"distance"`
	MovingTime    uint32  `json:"moving_time"`
	ElapsedTime   uint32  `json:"elapsed_time"`
	ElevationGain float64 `json:"elevation_gain"`
}

// Activity is activity
type Activity struct {
	Athlete        StravaAthlete `json:"athlete"`
	Distance       float64       `json:"distance"`
	MovingTime     uint32        `json:"moving_time"`
	ElapsedTime    uint32        `json:"elapsed_time"`
	ElevationGain  float64       `json:"total_elevation_gain"`
	Type           string        `json:"type"`
	StartDate      string        `json:"start_date"`
	TimeZoneOffset float64       `json:"utc_offset"`
	Title          string        `json:"name"`
}

type StravaAthlete struct {
	ID uint32 `json:"id"`
}
