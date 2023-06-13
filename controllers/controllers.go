package controllers

import (
	"leylahosseini/my-url-shorter/db"
	"leylahosseini/my-url-shorter/db/models"
	"leylahosseini/my-url-shorter/utils"

	"github.com/labstack/echo"
)

var db2 = db.ConnectToMysql()

//var err2 error

func Shorter(c echo.Context) error {

	u := new(models.Url)

	if err := c.Bind(u); err != nil {
		return err
	}
	code := utils.GenerateCode()
	//u.code = code
	//c.Request().ParseForm()
	URL2 := c.FormValue("url")
	//url := c.Request().FormValue("url")
	//echo.Logger.Debug(url)
	//c.middleware.Logger()
	//Url2 := c.Param("url")
	//Url2 := c.QueryParam("url")

	db2.Create(&models.Url{URL: URL2, Code: code})

	code2 := "http://localhost:8080" + "/" + code
	return c.JSON(201, code2)

}

func Code(c echo.Context) error {
	code := c.Param("code")
	var shortURL models.Url
	if result := db2.Where("code = ?", code).First(shortURL.Code); result.Error != nil {
		return echo.NewHTTPError(404, "URL not found")
	}
	return c.Redirect(301, shortURL.URL)
}
