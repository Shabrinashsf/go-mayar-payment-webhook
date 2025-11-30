package migrations

import (
	"go-mayar-payment-webhook/migrations/seeds"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListProductSeeder(db); err != nil {
		return err
	}
	return nil
}
