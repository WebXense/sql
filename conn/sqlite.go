package conn

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite connects to SQLite database
func SQLite(path string, silent bool) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(path), getGORMConfig(silent))
}

// SQLiteInMemory connects to in-memory SQLite database
func SQLiteInMemory(silent bool) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(sqliteInMemoryBuildDNS()), getGORMConfig(silent))
}
