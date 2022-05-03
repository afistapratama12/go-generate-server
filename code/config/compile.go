package config

import "fmt"

func compileConfig(driverName string) string {
	switch driverName {
	case "mysql":
		var (
			configInitStr = fmt.Sprintf(configInit, DriverMysql, configMysql)
		)
		return configInitStr
	case "postgres":
		var (
			configInitStr = fmt.Sprintf(configInit, DriverPostgres, configPostgres)
		)
		return configInitStr
	case "sqlite":
		var (
			configInitStr = fmt.Sprintf(configInit, DriverSqlite, configSqlite)
		)
		return configInitStr
	default:
		return ""
	}
}

func AddInitConfig(driverName string) (string, error) {
	configInitStr := compileConfig(driverName)

	if configInitStr == "" {
		return "", fmt.Errorf("driver name not found")
	} else {
		return configInitStr, nil
	}
}
