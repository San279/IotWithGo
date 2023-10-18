package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var db *sql.DB

type HaierSignal struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	IrId     int    `json:"irId"`
	Air      string `json:"air"`
	Fan      string `json:"fan"`
	Temp     int    `json:"temp"`
	Signal   string `json:"signal"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/iot")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("succesfully connect to dabase")
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	app.Get("/getRecentCommand", func(c *fiber.Ctx) error {

		get, err := db.Query("SELECT * FROM HaierOneDump ORDER BY id DESC LIMIT 1")
		if err != nil {
			fmt.Println(err)
			return err
		}

		response := HaierSignal{}
		for get.Next() {
			parseErr := get.Scan(&response.Id, &response.Username, &response.IrId,
				&response.Air, &response.Fan, &response.Temp, &response.Signal)
			if parseErr != nil {
				fmt.Println(err)
				return parseErr
			}
		}
		return c.JSON(response)
	})

	app.Listen(":6000")
	fmt.Println("server established")
}
