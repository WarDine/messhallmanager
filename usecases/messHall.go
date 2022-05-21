package usecases

// MessHall
type MessHall struct {
	MessHallUID      string `db:"messhalls_uid" json:"messHallUID"`
	Street           string `db:"street" json:"street"`
	City             string `db:"city" json:"city"`
	Country          string `db:"county" json:"county"`
	MenuUID          string `db:"menu_uid" json:"menuUID"`
	Status           string `db:"status" json:"status"`
	AttendanceNumber int    `db:"attendance_number" json:"attendanceNumber"`
}
