package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jorgechavezrnd/go-hexagonal_http_api/internal/creating"
	"github.com/jorgechavezrnd/go-hexagonal_http_api/internal/platform/server"
	"github.com/jorgechavezrnd/go-hexagonal_http_api/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseSerivce(courseRepository)

	srv := server.New(host, port, creatingCourseService)
	return srv.Run()
}
