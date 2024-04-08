package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create the HTML template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	// Setup session middleware
	store := session.New(session.Config{
		Expiration: 4 * time.Hour,
	})

	app.Static("/", "./public")

	// Route for serving the login page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	// Route for handling the login logic
	app.Post("/login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if user := getUser(username); user.Password == password {
			sess, err := store.Get(c)
			if err != nil {
				return err
			}
			sess.Set("username", username)
			if err := sess.Save(); err != nil {
				return err
			}

			return c.Redirect("/dashboard", fiber.StatusFound)
		}
		return c.SendString("Invalid login credentials.")
	})

	// Route for serving the dashboard page
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}

		if username := sess.Get("username"); username != nil {
			return c.Render("dashboard", fiber.Map{"Username": username})
		}
		// Redirect to login if not authenticated
		return c.Redirect("/", fiber.StatusFound)
	})

	log.Fatal(app.Listen(":30000"))
}

func genRandomUsers() {
	for i := 0; i < 10; i++ {
		username := genRandomString()
		password := genRandomString()
		users[username] = password
	}

	for k, v := range users {
		log.Println(k, v)
	}
}

var users = map[string]string{
	"admin": "password_1234",
	"john":  "s1z7jWa3eCulBhRn",
	"jane":  "E6tnC4r671iyFNC3",
	"alice": "xL90lV1RTiZ269Z7",
	"bob":   "mNHLGiwpb996L5hf",
	"smith": "qoZUjp7N3CPyu2u4",
	"doe":   "iqckyK7VcmK2eprL",
	"joe":   "iAs901Ty65B3EFey",
	"jim":   "dgcetXfvCh0yJA2w",
	"jill":  "RIKPwXvUPVivBdgG",
}

func genRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func getUser(username string) User {
	if password, ok := users[username]; ok {
		log.Println(username, password)
	}
	return fakeUsers(username, users[username])

}

type User struct {
	Username string
	Password string
}

func fakeUsers(username string, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}
