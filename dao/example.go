package dao

import (
	//"database/sql"
	"goclean/entity"
	"database/sql"
	"time"
)

type exampleDao struct {
	dbSession        *sql.DB
}

func InitiateExampleDao(dbSession *sql.DB) *exampleDao {
	return &exampleDao{
		dbSession:        dbSession,
	}
}

func (u *exampleDao) GetExampleData(key string) entity.ExampleEntity {
	return entity.ExampleEntity{
		Key:          key,
		Name:         "Example product",
		Color:        "Red",
		Remain:       213,
		HasPromotion: false,
		CreatedDate:  time.Now(),
		Size:         8,
		Type:         2,
	}
}
