package conn

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getGORMConfig(silent bool) *gorm.Config {
	config := &gorm.Config{}
	if silent {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return config
}

func mySQLBuildDNS(host, port, username, password, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
}

func sqliteInMemoryBuildDNS() string {
	return "file::memory:?cache=shared"
}
