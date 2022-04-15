package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

//structs
type LanguageState struct {
	Success  string                 `json:"success"`
	Data     map[string]interface{} `json:"data"`
	Hello    string                 `json:"hello"`
	GoStruct string                 `json:"go_struct"`
}

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
				"hello": "We are learning Go JSON Decoding from JSON String to specific Struct!",
				"go_struct":"Go's structs are typed collections of fields that awesome useful"
			}
		`

//Main Func
func main() {
	//instantiate app
	app := fiber.New()

	//routes
	app.Get("/users", func(c *fiber.Ctx) error {
		//define variable where decoded json will store, in this case our LanguageState Struct
		var result LanguageState

		//unmarshal/decode, &result is to pointing the address of result variable
		json.Unmarshal([]byte(data), &result)

		//return json
		return c.JSON(result)
	})

	//listen
	app.Listen(":8080")
}
