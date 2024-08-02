package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
Mutexes can be used to protect this shared data from simultaneous access
and possible corruption, while a WaitGroup can ensure that the server
doesn’t shut down until all requests have been fully processed.
https://medium.com/@andrekardec/mutexes-and-waitgroups-better-performance-with-golangs-sync-package-9c77c5125bf5
*/

type (
	response struct {
		Status string `json:"status"`
		Data   []user `json:"data"`
	}

	user struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Cards []card `json:"cards"`
	}
	card struct {
		ID         int     `json:"id"`
		CardType   string  `json:"cardType"`
		CardNumber string  `json:"cardNumber"`
		Balance    float64 `json:"balance"`
	}
)

var (
	users = []user{}
	seq   = 1
	lock  = sync.Mutex{}
)

func createUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	r := new(response)
	u := new(user)
	u.ID = seq

	if err := c.Bind(u); err != nil { //Mapping JSON data into "u" variable
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	users = append(users, *u)
	r.Status = http.StatusText(http.StatusCreated)
	r.Data = users
	seq++
	return c.JSON(http.StatusCreated, r)
}

func getUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	index := findIndexById(users, id)
	if index == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, users[index])
}

func updateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	u := new(user)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	u.ID = id
	index := findIndexById(users, id)
	if index == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	users[index] = *u
	return c.JSON(http.StatusOK, users[index])
}

/* func deleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	removeAtID(users, id)
	return c.NoContent(http.StatusNoContent)
} */

func deleteCard(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	cardId, _ := strconv.Atoi(c.QueryParam("cardID"))
	usrIndex := findIndexById(users, id)

	if usrIndex == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}

	cards := users[usrIndex].Cards
	cardIndex := findCardById(cards, cardId)

	if cardIndex == -1 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Card not found"})
	}
	cards = append(cards[:cardIndex], cards[cardIndex+1:]...)
	users[usrIndex].Cards = cards
	return c.JSON(http.StatusOK, users)
}

func getAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	r := new(response)
	r.Status = http.StatusText(http.StatusOK)
	r.Data = users
	return c.JSON(http.StatusOK, r)
}

/* func removeAtID(s []user, id int) []user {
	index := findIndexById(s, id)
	return append(s[:index], s[index+1:]...)
} */

func findIndexById(s []user, id int) int {
	index := -1
	for i, data := range s {
		if data.ID == id {
			index = i
			break
		}
	}
	return index
}

func findCardById(cards []card, id int) int {
	cardIndex := -1
	for i, v := range cards {
		if v.ID == id {
			cardIndex = i
			break
		}
	}
	return cardIndex
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteCard)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
