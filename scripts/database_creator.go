package main

import (
	"fmt"
	"github.com/SebastienBoisard/newsletters/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {

	server.LoadConfig()

	db, err := server.InitDatabase()
	if err != nil {
		log.Fatalf("Can't initialized database [err: %v]", err)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalf("Can't close database [err: %v]", err)
		}
	}()

	fmt.Printf("clean database...")
	server.CleanDatabase(db)
	fmt.Printf(" database cleaned\n")

	fmt.Printf("create database...")
	server.CreateDatabase(db)
	fmt.Printf(" database created\n")

	fmt.Printf("fill database...")
	server.FillDatabase(db)
	fmt.Printf("database filled\n")

	fmt.Printf("\n=> database ready\n")
}
