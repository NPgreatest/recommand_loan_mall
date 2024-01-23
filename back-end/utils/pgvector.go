package utils

import (
	"fmt"
	"github.com/pgvector/pgvector-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormItem struct {
	gorm.Model
	Embedding pgvector.Vector `gorm:"type:vector(3)"`
}

func PgvectorGetId(input []float32) (ids []uint, err error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=root dbname=vector_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"),
		&gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&GormItem{})
	if err != nil {
		return nil, err
	}
	var items []GormItem
	db.Raw("SELECT id FROM vectors ORDER BY (embedding <-> ?) LIMIT 5", pgvector.NewVector(input)).Scan(&items)
	for _, item := range items {
		fmt.Printf("ID: %d, Embedding: %v\n", item.ID, item.Embedding)
		ids = append(ids, item.ID)
	}
	return
}
