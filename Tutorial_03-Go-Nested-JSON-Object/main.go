package main

import (
	"github.com/gofiber/fiber/v2"
)

//Main Func
func main() {
	//instantiate app
	app := fiber.New()

	//routes
	app.Get("/users", func(c *fiber.Ctx) error {
		//define variable where decoded json will store, in this case our LanguageState Struct
		result := map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"PHP":    "Lovely",
				"NodeJs": "Cool",
				"Python": [4]string{
					"Simple",
					"Shorter",
					"Quick",
					"Lazy",
				},
				"GO": [3]int{
					1,
					2,
					3,
				},
				"JavaScript": map[string]string{
					"web":      "Every web dev should know",
					"type":     "Dynamically typed language",
					"reaction": "We love JS but confused in so many Libraries",
				},
				"rust": map[string]interface{}{
					"reaction":    "Most Loved Language by Developer",
					"star":        65.5,
					"description": "This map containing string + float value where we set empty interface",
				},
				"mix_values": [4]interface{}{
					10,
					"Love Programming",
					14.3,
					"This array containing mixed values like string, int, float etc. No types specified as we defined empty interface",
				},
				"undefined_length": []string{
					"Here is no limit of items...",
				},
			},
		}

		//return json
		return c.JSON(result)
	})

	//listen
	app.Listen(":8080")
}
