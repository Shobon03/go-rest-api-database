package helpers

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func HandleDBError(dbErr error) bool {
	if dbErr != nil {
		date := strings.Split(time.Now().Format(time.RFC3339), "T")
		f, err := os.OpenFile(
			date[0] + " error.log", 
			os.O_RDWR | os.O_CREATE | os.O_APPEND, 
			0766,
		)

		f.Write([]byte(dbErr.Error()))
		
		if (err != nil)  {
			log.Fatal("Could not write due to permissions")
		}

		defer f.Close()
		return true
	}

	return false
}

func LoadEnvironmentVariables() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error loading .env file")
  }
}