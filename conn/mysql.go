package conn

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL connects to the MySQL database
func MySQL(host, port, username, password, database string, silent bool) (*gorm.DB, error) {
	dsn := mySQLBuildDNS(host, port, username, password, database)
	return gorm.Open(mysql.Open(dsn), getGORMConfig(silent))
}
