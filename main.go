package main

import (
	"os"

	"github.com/MuhAndriJP/dealls.git/repo/mysql"
	"github.com/MuhAndriJP/dealls.git/routes"
)

func main() {
	mysql.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
