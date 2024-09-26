package dbservice

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/vivilCODE/chess/db/log"
	"github.com/vivilCODE/chess/db/model"
)

var (
	connectionString = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	ErrorNoUserFound = errors.New("no user found with matching id")
)

// DBService holds the connection to the database.
// It has methods to query different data and parse them into models defined in the models package.
type DBService struct {
	c      Config
	conn *sql.DB
}

// Config contains all the necessary parameters to form a database connection.
type Config struct {
	DBPort   int
	DBName   string
	Host     string
	User     string
	Password string
}

// Func New creates a new instance of *DBController with the given configuration.
func New(c Config) (*DBService, error) {
	dbc := &DBService{
		c: c,
	}

	return dbc, nil
}

// Func Connect opens a connection to the postgres database.
func (s *DBService) Connect() error {
	connectionString := fmt.Sprintf(connectionString,
		s.c.Host, s.c.DBPort, s.c.User, s.c.Password, s.c.DBName)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("unable to open db connection: %v", err)
	}

	if err := conn.Ping(); err != nil {
		return fmt.Errorf("unable to contact database: %v", err)
	}

	s.conn = conn

	log.Logger.Info("successfully contected to", "database", s.c.DBName)

	return nil
}

// Func Disconnect closes the connection to the postgres database.
func (s *DBService) Disconnect() error {
	if err := s.conn.Close(); err != nil {
		return fmt.Errorf("unable to close connection to databse: %v", err)
	}
	return nil
}


func (s *DBService) InsertUser(user model.User) (error) {
	insertUserStatement := `
    INSERT INTO "user" (id, name, email, signup_timestamp)
    VALUES ($1, $2, $3, $4)`

	_, err :=s.conn.Exec(insertUserStatement, user.Id, user.Name, user.Email, user.SignedUp)
		if err != nil {
			return fmt.Errorf("unable to insert user: %v", err)
		}

	return nil
}

func (s *DBService) GetUser(id string) (model.User, error) {
	getUserStatement := `SELECT id, name, email, signup_timestamp FROM "user" WHERE id = $1;`
	
	row :=s.conn.QueryRow(getUserStatement, id)
		
	var user = model.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.SignedUp)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, ErrorNoUserFound
		}

		return model.User{}, fmt.Errorf("unexpected error scanning rows for user with id %s, err: %v", id, err)
	}

	return user, nil
}
