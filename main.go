package main

import (
	"net/http"
	"os"
	"io"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	url := os.Getenv("URL")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/*", func(c echo.Context) error {
		if url == "" {
			return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
		}else{
			res, err := http.Get(url)
			if err != nil {
				fmt.Printf("error making http request: %s\n", err)
				os.Exit(1)
			}
		
			fmt.Printf("client: got response!\n")
			fmt.Printf("client: status code: %d\n", res.StatusCode)

			defer res.Body.Close()

			b, err := io.ReadAll(res.Body)
			// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier

			var responseStr = "we got a response from " + url + " they said " + string(b)
			return c.JSON(http.StatusOK, responseStr)
		}
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
