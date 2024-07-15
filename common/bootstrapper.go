package common

func init() {
	initConfig()
	initKeys()
	setLogLevel(Level(AppConfig.LogLevel))
	createDbSession()
	addIndexes()
}