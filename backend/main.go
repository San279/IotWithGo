package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v4"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var jwtSecret string = "iotcharansanitwong"

type HaierSignal struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	IrId     int    `json:"irId"`
	Air      string `json:"air"`
	Fan      string `json:"fan"`
	Temp     int    `json:"temp"`
	Signal   string `json:"signal"`
}

type SensorData struct {
	Id          int    `json:"id"`
	Humidity    string `json:"humidity"`
	Temperature string `json:"temperature"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/iot")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var PreviousMsgCount int = 0
	fmt.Println("succesfully connect to dabase")
	app := fiber.New()

	authorizeSite := jwtware.New(jwtware.Config{

		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return fiber.ErrUnauthorized
		},
	})
	app.Use("/remotecontrol", authorizeSite)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowMethods: "POST, GET, DELETE, UPDATE",
		AllowHeaders: "*",
	}))

	app.Get("/getRecentCommand", func(c *fiber.Ctx) error {

		var currentMsgCount int

		err := db.QueryRow("SELECT COUNT(*) FROM HaierOneDump;").Scan(&currentMsgCount)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if PreviousMsgCount == currentMsgCount {
			return c.JSON("No new message")
		}
		PreviousMsgCount = currentMsgCount
		//fmt.Println(PreviousMsgCount)
		//fmt.Println(currentMsgCount)
		response := HaierSignal{}
		QueryErr := db.QueryRow("SELECT * FROM HaierOneDump ORDER BY id DESC LIMIT 1").Scan(&response.Id, &response.Username, &response.IrId,
			&response.Air, &response.Fan, &response.Temp, &response.Signal)
		if QueryErr != nil {
			fmt.Println(err)
			return err
		}

		return c.JSON(response)
	})

	app.Post("/postHumidTemp", func(c *fiber.Ctx) error {

		request := SensorData{}
		ParseErr := c.BodyParser(&request)
		if ParseErr != nil {
			return fiber.ErrBadRequest
		}

		query := "insert SensorData (id, humidity, temperature) values (0,?,?)"
		insert, err := db.Exec(query, request.Humidity, request.Temperature)

		if err != nil {
			fmt.Println("fail to insert to db")
			return err
		}

		id, err := insert.LastInsertId()
		if err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity)
		}
		request.Id = int(id)
		fmt.Println(request)
		return c.Status(fiber.StatusCreated).JSON(request)

	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		request := User{}

		ParseErr := c.BodyParser(&request)
		fmt.Println(request)
		if ParseErr != nil {
			return fiber.ErrBadRequest
		}

		if request.Username == "" || request.Password == "" {
			return fiber.ErrUnprocessableEntity
		}

		password, cryptErr := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
		if cryptErr != nil {
			return fiber.ErrExpectationFailed
		}

		request.Password = string(password)
		query := "INSERT Users VALUES(?,?,?)"
		insert, inErr := db.Exec(query, 0, request.Username, request.Password)

		if inErr != nil {
			return inErr
		}

		inId, err := insert.LastInsertId()
		if err != nil {
			return err
		}

		request.Id = int(inId)
		return c.Status(fiber.StatusCreated).JSON(request)

	})

	app.Post("/login", func(c *fiber.Ctx) error {
		request := User{}

		parErr := c.BodyParser(&request)

		if parErr != nil {
			return fiber.ErrBadRequest
		}

		if request.Username == "" || request.Password == "" {
			return fiber.ErrBadRequest
		}

		user := User{}
		query := "SELECT id, username, password from Users WHERE username = ?"
		err := db.QueryRow(query, request.Username).Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			fmt.Println("cant find in database")
			return err
		}

		check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

		if check != nil {
			return check
		}

		claims := jwt.MapClaims{
			"id":  strconv.Itoa(user.Id),
			"exp": time.Now().Add(time.Hour * 72).Unix(),
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		token, err := jwtToken.SignedString([]byte(jwtSecret))
		if err != nil {
			return fiber.ErrInternalServerError
		}
		fmt.Println(token)
		return c.JSON(fiber.Map{ //return json in map format
			"user":  user.Username,
			"token": token,
		})
	})

	app.Get("/remotecontrol", func(c *fiber.Ctx) error {
		return c.SendString("okidoki")
	})

	app.Listen(":5001")
	fmt.Println("server established")
}
