package domain

type AddMessHallInfoQuery struct {
	MessHallAdminUID      string
	MessHallAdminNickname string
	MessHallUID           string
	Street                string
	City                  string
	Country               string
	MenuUID               string
	Status                string
	AttendanceNumber      int
}
