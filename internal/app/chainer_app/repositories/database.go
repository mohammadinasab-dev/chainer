package repositories

type DbReposiroty interface {
	GetAllModels() []interface{}
	CreateTables() error
}
