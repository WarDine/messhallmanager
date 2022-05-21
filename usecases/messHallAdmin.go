package usecases

// MessHallAdmin
type MessHallAdmin struct {
	MessHallAdminUID string `db:"messhalls_admins_uid" json:"messHallAdminUID"`
	Nickname         string `db:"nickname" json:"nickname"`
	MessHallUID      string `db:"messhall_uid" json:"messHallUID"`
}
