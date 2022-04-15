package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

//variables
var data string = `
			{
				"success": true,
				"data": {
					"PHP":    "Lovely",
					"NodeJs": "Cool",
					"Python": [
						"Simple",
						"Shorter",
						"Quick",
						"Lazy"
					],
					"GO":"Love for fast & simplicity"
				},
				"hello": "We are learning Go JSON Decoding from JSON String!"
			}
		`

//Main Func
func main() {
	//instantiate app
	app := fiber.New()

	//routes
	app.Get("/users", func(c *fiber.Ctx) error {
		//define variable where decoded json will store
		var result map[string]interface{}

		//unmarshal/decode, &result is to pointing the address of result variable
		json.Unmarshal([]byte(data), &result)

		//return json
		return c.JSON(result)
	})

	//listen
	app.Listen(":8080")
}
