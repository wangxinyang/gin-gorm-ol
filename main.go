package main

import "gin-gorm-ol/router"

func main() {
	r := router.Router()

	r.Run("127.0.0.1:3000")
}
