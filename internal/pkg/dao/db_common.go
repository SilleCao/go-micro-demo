package dao

func DbQuery() *Query {
	return Use(Db())
}
