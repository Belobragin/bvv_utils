package util

const (
	TimeFormatContext  = "2006-01-02T15:04:05.000"
	TimeFormatContext1 = "2006-01-02"
	OutputLimitPerPage = 30
)

const (
	KeyUserUUID             = "user_uuid"
	FilterKeyCanUpdate      = "can_update"
	FilterKeyID             = "id"
	FilterKeyActive         = "active"
	FilterKeyDateRegister   = "date_register"
	FilterKeyLastLogin      = "last_login"
	FilterKeyLoginAttempts  = "register_attempts"
	FilterKeyRole           = "role"
	FilterKeyDistributionID = "distribution_id"
	FilterKeyUserID         = "user_id"
	// general keys for filtering:
	FilterKeyUserUUID         = "filter_user_uuid"
	FilterKeyName             = "name"
	FilterKeyEmail            = "email"
	FilterKeyCreatedTimeStamp = "created_at"
	FilterKeyUpdatedTimeStamp = "updated_at"
	FilterKeyPage             = "page"
	FilterSortOrder           = "sort_order"
	FilterSortBy              = "sort_by"
	// role keys for filtering
	FilterKeyRoleID = "role_id"
	// session keys for filtering:
	FilterKeySessionUUID = "session_uuid"
)

const (
	EmailRe = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
	PhoneRe = `(9[0-9]{9})`
)

var (
	PermittedSortOrder = []string{"asc", "desc"}
)
