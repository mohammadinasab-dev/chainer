package config

import (
	"time"
)

var (
	GetDatabaseConnectionDB   = getDatabaseConnectionDB
	GetDatabaseConnectionPass = getDatabaseConnectionPass
	GetDatabaseConnectionUser = getDatabaseConnectionUser
	GetDatabaseConnectionUrl  = getDatabaseConnectionUrl
	// GetDatabaseMaxIdleConnections returns max idle connection from database section in toml file
	GetDatabaseMaxIdleConnections = getDatabaseMaxIdleConnections

	// GetDatabaseMaxOpenConnections returns max open connections from database section in toml file
	GetDatabaseMaxOpenConnections = getDatabaseMaxOpenConnections

	// GetDatabaseConnectionMaxLifetime returns connection max lifetime from database section in toml file
	GetDatabaseConnectionMaxLifetime = getDatabaseConnectionMaxLifetime
)

func getDatabaseConnectionDB() string {
	return getConfigString("database.connection_db")
}

func getDatabaseConnectionPass() string {
	return getConfigString("database.connection_pass")
}

func getDatabaseConnectionUser() string {
	return getConfigString("database.connection_user")
}

func getDatabaseConnectionUrl() string {
	return getConfigString("database.connection_url")
}

func getDatabaseMaxIdleConnections() int {
	return getConfigInt("database.max_idle_connections")
}

func getDatabaseMaxOpenConnections() int {
	return getConfigInt("database.max_open_connections")
}

func getDatabaseConnectionMaxLifetime() time.Duration {
	return getConfigDuration("database.connection_max_lifetime")
}
