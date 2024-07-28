package main

import (
	"fmt"
	"log"

	"github.com/Gierdiaz/GoStock/api/internal/database"
	"github.com/Gierdiaz/GoStock/api/internal/router"
)

func main() {

	database.ConnectionDatabase()

	r := router.SetupRouter()
	fmt.Println("Server running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}