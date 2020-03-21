package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

//GetNow exact time
func GetNow() time.Time {
	return time.Now().UTC()
}

//Format given Time to layout
func Format(date time.Time, layout string) string {
	return date.Format(layout)
}

//GetNowAsString With layout
func GetNowAsString() string {
	return Format(GetNow(), apiDateLayout)
}

//GetNowAsDataBaseString With layout
func GetNowAsDataBaseString() string {
	return Format(GetNow(), apiDbLayout)
}
