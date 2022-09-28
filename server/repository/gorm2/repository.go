package gorm2

import (
	"database/sql"
	"fmt"

	"github.com/digicon-hack-07/ez-toon/server/config"
	"github.com/digicon-hack-07/ez-toon/server/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Config struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

func GetGorm2Config(c *config.Config) *Config {
	return &Config{
		Hostname: c.MariaDBHostname,
		Port:     c.MariaDBPort,
		Database: c.MariaDBDatabase,
		Username: c.MariaDBUsername,
		Password: c.MariaDBPassword,
	}
}

func NewGorm2Repository(c *Config) (*Repository, error) {
	db, err := newDBConnection(c)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db :%w", err)
	}

	db.AutoMigrate(
		repository.Project{},
		repository.Page{},
		repository.Line{},
		repository.Dialogue{},
	)

	return &Repository{
		db: db,
	}, nil
}

func newDBConnection(c *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Hostname, c.Port, c.Database) + "?parseTime=true&loc=Local&charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB : %w", err)
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	return db, nil
}

func GetSQLDb(repo *Repository) (*sql.DB, error) {
	db, err := repo.db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db : %w", err)
	}

	return db, nil
}
