package daointerface

import "goclean/entity"

type ExampleDaoInterface interface {
	GetExampleData(key string) entity.ExampleEntity
}
