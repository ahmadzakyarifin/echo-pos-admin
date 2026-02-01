package domain

import "database/sql"

type User struct {
	ID       int           `db:"id"`
	OwnerID  sql.NullInt64 `db:"owner_id"`
	TenantID sql.NullInt64 `db:"tenant_id"`
	Name     string        `db:"name"`
	Username string        `db:"username"`
	Email    string        `db:"email"`
	Password string        `db:"password"`
	Role     string        `db:"role"`
}
