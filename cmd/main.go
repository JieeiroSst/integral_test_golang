package main

import (
	"strconv"
	"database/sql"
	"fmt"
	"os"
	"bufio"
	"integral_test/database"
	"integral_test/service"
	
)

func main() {
	connString := "dbname=<your main db name> sslmode=disable"
	db, _ := sql.Open("postgres", connString)
	store := database.NewStore(db)
	service := &service.Service{Store: store}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		ID, _ := strconv.Atoi(scanner.Text())
		err := service.GetNumber(ID)
		if err != nil {
			fmt.Printf("result invalid: %v", err)
			continue
		}
		fmt.Println("result valid")
	}
}