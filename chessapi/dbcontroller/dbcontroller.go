package dbcontroller

import (
	"database/sql"
	"fmt"
)

var (
	connectionString = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

// DBController holds the connection to the database.
// It has methods to query different data and parse them into models defined in the models package.
type DBController struct {
	c      Config
	dbConn *sql.DB
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
func New(c Config) (*DBController, error) {
	dbc := &DBController{
		c: c,
	}

	return dbc, nil
}

// Func Connect opens a connection to the postgres database.
func (dbc *DBController) Connect() error {
	connectionString := fmt.Sprintf(connectionString,
		dbc.c.Host, dbc.c.DBPort, dbc.c.User, dbc.c.Password, dbc.c.DBName)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("unable to open db connection: %v", err)
	}

	if err := conn.Ping(); err != nil {
		return fmt.Errorf("unable to contact database: %v", err)
	}

	dbc.dbConn = conn

	fmt.Printf("successfully connected to %s\n", dbc.c.DBName)

	return nil
}

// Func Disconnect closes the connection to the postgres database.
func (dbc *DBController) Disconnect() error {
	if err := dbc.dbConn.Close(); err != nil {
		return fmt.Errorf("unable to close connection to databse: %v", err)
	}
	return nil
}

// Func QueryChapterInfo takes a chapter_id and uses it to fetch all the information needed to create
// a *models.ChapterInfo{} struct.
// func (dbc *DBController) QueryChapterInfo(chapterID int64) (*models.ChapterInfo, error) {
// 	queryString := fmt.Sprintf(chapterInfoQuery, chapterID)

// 	rows, err := dbc.dbConn.Query(queryString)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to query chapter info, err: %v", err)
// 	}
// 	defer rows.Close()

// 	var c = &models.ChapterInfo{}

// 	foundChapter := false
// 	for rows.Next() {
// 		foundChapter = true
// 		rows.Scan(&c.Company, &c.Project, &c.Chapter)
// 	}

// 	if !foundChapter {
// 		return nil, ErrorNoDataFound
// 	}

// 	return c, nil
// }

// Func QueryVersionInfo takes a chapter_id and uses it to fetch all the information needed to create an array
// of  *models.Versioninfo{}.
// func (dbc *DBController) QueryVersionInfo(chapterID int64) ([]*models.VersionInfo, error) {
// 	queryString := fmt.Sprintf(versionInfoQuery, chapterID)

// 	rows, err := dbc.dbConn.Query(queryString)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to query version info, err: %v", err)
// 	}
// 	defer rows.Close()

// 	var versionInfos = []*models.VersionInfo{}

// 	for rows.Next() {
// 		var v = models.VersionInfo{}
// 		rows.Scan(&v.CreatedBy, &v.ChapterVersionID, &v.VersionNumber, &v.Created, &v.AppVersion)
// 		versionInfos = append(versionInfos, &v)
// 	}

// 	return versionInfos, nil
// }
