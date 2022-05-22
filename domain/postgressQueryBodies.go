package domain

type AddMessHallInfoQuery struct {
	MessHallAdminNickname string
	Street                string
	City                  string
	Country               string
	Status                string
	AttendanceNumber      int
}

type GetMessHallInfoResponse struct {
	MessHallAdminNickname string
	Street                string
	City                  string
	Country               string
	Status                string
	MenuUID               string
	AttendanceNumber      int
}
