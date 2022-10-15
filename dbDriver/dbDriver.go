package dbDriver

type DBDriver interface {
	CloseDB()
}

func init() {
	NewMysqlDriverWithoutInterface()
}
