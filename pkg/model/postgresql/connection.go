package postgresql

// Connection holds the necessary information to connect to a PostgreSQL database.
type Connection struct {
	// Address is the hostname or IP address of the PostgreSQL server.
	Address string
	// Port is the port number on which the PostgreSQL server is listening.
	Port int
	// Username is the PostgreSQL user to connect as.
	Username string
	// Password is the PostgreSQL user's password.
	//nolint:gosec // This field is necessary for the connection and is not a hardcoded credential.
	Password string
}
