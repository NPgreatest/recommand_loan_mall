package utils

import (
	"fmt"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
	"main.go/global"
)

type GormItem struct {
	gorm.Model
	Embedding pgvector.Vector `gorm:"type:vector(3)"`
}

func PgvectorGetId(input []float32) (ids []uint, err error) {
	err = global.GVA_Postgres.AutoMigrate(&GormItem{})
	if err != nil {
		return nil, err
	}
	var items []GormItem
	global.GVA_Postgres.Raw("SELECT id FROM vectors ORDER BY (embedding <-> ?) LIMIT 5", pgvector.NewVector(input)).Scan(&items)
	for _, item := range items {
		fmt.Printf("ID: %d, Embedding: %v\n", item.ID, item.Embedding)
		ids = append(ids, item.ID)
	}
	return
}
