package repository

type Account struct {
	ApiKey interface{} `db:"api_key"`
}

type AccountEntity struct {
	ApiKey string
}
