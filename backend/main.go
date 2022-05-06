package main

import (
	"demo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	addressDb := "postgres://user:password@db:5432/data?sslmode=disable"
	dbClient := db.NewPostgresClient(addressDb)
	h := &Handler{ dbClient: dbClient }
	
	e := echo.New()
	e.GET("/accounts/:id", h.Home)
	e.Logger.Fatal(e.Start(":1323"))
}

type Handler struct {
	dbClient        db.DbClient
}

func (h *Handler) Home(c echo.Context) error {
	accountID := c.Param("id")
	account, err := h.dbClient.InquiryAccount(c.Request().Context(), accountID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Account not found")
	}
	return c.JSON(http.StatusOK, account)
}