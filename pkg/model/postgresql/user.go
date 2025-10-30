package postgresql

import (
	pg "github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// UserData defines a PostgreSQL user.
type UserData struct {
	// User is the PostgreSQL role.
	User *pg.Role
	// Password is the PostgreSQL user's password.
	Password pulumi.StringOutput
}
