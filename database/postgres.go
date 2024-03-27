package database

import (
	"banky/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
		cfg.Db.SSLMode,
		cfg.Db.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{
		Db: db,
	}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}

func CreateUserIDTrigger(db *gorm.DB) error {
	// Trigger function SQL
	triggerSQL := `
		CREATE OR REPLACE FUNCTION generate_user_id()
		RETURNS TRIGGER AS $$
		DECLARE
			max_id BIGINT;
			new_id TEXT;
		BEGIN
			SELECT COALESCE(MAX(SUBSTRING(id, 2)::BIGINT), 0) INTO max_id FROM users;
			new_id := 'U' || LPAD(CAST((max_id + 1) AS TEXT), 6, '0');
			NEW.id := new_id;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER user_id_trigger
		BEFORE INSERT ON users
		FOR EACH ROW
		EXECUTE PROCEDURE generate_user_id();
	`

	// Run the trigger SQL
	return db.Exec(triggerSQL).Error
}
