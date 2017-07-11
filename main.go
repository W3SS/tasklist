package main

import (
	"database/sql"

	"./handlers"
	"github.com/labstack/echo"
	 _ "github.com/go-sql-driver/mysql"
)

func main() {
	db := initDb()
	migrate(db)

	e := echo.New()
	e.Static("/", "public/dist")
	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/task", handlers.PostTask(db))
	e.PUT("/task", handlers.PutTask(db))
	e.DELETE("/task/:id", handlers.DeleteTask(db))
	e.Start(":8080")
}

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@/govue_service")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(20) NOT NULL,
			done INT NOT NULL
		);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
