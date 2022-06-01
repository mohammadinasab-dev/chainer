package postgres

import (
	"chainer/internal/app/chainer_app/models"
	repo "chainer/internal/app/chainer_app/repositories"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type postgres struct {
	db *pg.DB
}

func NewDBDataSource(db *pg.DB) repo.DbReposiroty {
	return &postgres{
		db: db,
	}
}

func (p *postgres) GetAllModels() []interface{} {
	return []interface{}{
		(*models.Address)(nil),
	}
}

func (p *postgres) CreateTables() error {
	for _, model := range p.GetAllModels() {
		err := p.db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
