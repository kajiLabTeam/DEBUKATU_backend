package main

import (
	"kajiLabTeam/DEBUKATU/lib"
	"kajiLabTeam/DEBUKATU/router"
)

func main() {
	lib.SqlConnect()
	router.Router()
}
