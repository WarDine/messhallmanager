package usecases

type MessHall struct {
	MessHallUID      int    `db:"mass_halls_uid" json:"messHallUID"`
	Street           string `db:"street" json:"street"`
	City             string `db:"city" json:"city"`
	Country          string `db:"county" json:"county"`
	MenuUID          int    `db:"menu_uid" json:"menuUID"`
	AttendanceNumber int    `db:"attendance_number" json:"attendanceNumber"`
}

// tx.NamedExec("INSERT INTO mass_hall (mass_halls_uid, street, city, county, menu_uid, attendance_number) VALUES (:ingredient_name, :recipe_uid, :amount)", &secondIngredient
//  psql -U postgres -d wardine
