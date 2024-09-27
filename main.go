package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong2",
		})
	})
	r.GET("/foo", func(c *gin.Context) {
		stats := db.Stats()
		fmt.Printf("%+v\n", stats)
		rows, err := db.Query("SELECT * FROM people")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		ages := make([]int, 0)
		for rows.Next() {
			var age int
			if err := rows.Scan(&age); err != nil {
				log.Fatal(err)
			}
			ages = append(ages, age)
		}
		// Check for errors from iterating over rows.
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", ages)

		c.String(http.StatusOK, fmt.Sprintf("%+v", ages))
	})
	r.Run("0.0.0.0:1000")
}
