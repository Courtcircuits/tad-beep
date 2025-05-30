package sql

import (
	"fmt"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/types"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Sql interface {
	Setup() error
	Shutdown() error
	CreateMessage(message *types.Message) (*types.Message, error)
	GetMessages(channelID string) ([]types.Message, error)
}

type sql struct {
	db         *sqlx.DB
	migrations string
	username   string
	password   string
	host       string
	port       int
	dbname     string
}

func NewSql(config Config) (Sql, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.Dbname))
	if err != nil {
		return nil, err
	}
	return sql{
		db:         db,
		migrations: config.Migrations,
		username:   config.Username,
		password:   config.Password,
		host:       config.Host,
		port:       config.Port,
		dbname:     config.Dbname,
	}, nil
}

func (m sql) Setup() error {
	migrations, err := migrate.New(fmt.Sprintf("file://%s", m.migrations), fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", m.username, m.password, m.host, m.port, m.dbname))

	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil {
		return err
	}
	fmt.Println("Migrations up")

	return nil
}

func (m sql) Shutdown() error {
	return m.db.Close()
}

func (m sql) CreateMessage(message *types.Message) (*types.Message, error) {
	var new_message types.Message
	err := m.db.Get(&new_message, `INSERT INTO messages (content, channel_id, owner_id, created_at) VALUES ($1, $2, $3, $4) RETURNING *`, message.Content, message.Channel, message.Owner, message.CreatedAt)
	if err != nil {
		return &types.Message{}, err
	}
	return &new_message, nil
}

func (m sql) GetMessages(channelID string) ([]types.Message, error) {
	var messages []types.Message

	err := m.db.Select(&messages, `SELECT * FROM messages WHERE channel_id = $1`, channelID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
