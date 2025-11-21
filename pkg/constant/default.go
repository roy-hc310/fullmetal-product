package constant

const (
	DefaultLimit  int    = 10
	DefaultSortBy string = "id"
	DefaultOrder  string = "desc"
)

type key string

const (
	UserID key = "user_id"
	Role   key = "role"
)
