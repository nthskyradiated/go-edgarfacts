package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nthskyradiated/go-edgarfacts/internal/facts"
	"github.com/nthskyradiated/go-edgarfacts/internal/storage"
)

func main() {
	// cik := ""
	// name := ""
	// org := ""
	// email := ""

	// facts, err := facts.LoadFacts(cik, name, org, email)
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(facts))
	// bucketName := "go-edgarfacts-google-storage"
	// filepath := fmt.Sprintf("sec/edgar/facts/stage/%s.json", cik)
	// err = storage.UploadBytes(facts, bucketName, filepath)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Uploaded %s\n", cik)

	var cik string
	var org string
	var name string
	var email string

	flag.StringVar(&cik, "cik", "", "CIK number")
	flag.StringVar(&org, "org", "", "Your Org name")
	flag.StringVar(&name, "name", "", "Your name")
	flag.StringVar(&email, "email", "", "Your email address")

	flag.Parse()

	if len(cik) != 10 {
		panic("cik must be 10 characters long")
	}

	if org == "" {
		panic("please provide your Organization's name")
	}

	if name == "" {
		panic("please provide your name")
	}

	if email == "" {
		panic("email address required")
	}

	bucketName := os.Getenv("BUCKET_NAME")
	folderPath := os.Getenv("STAGE")

	if bucketName == "" || folderPath == "" {
		panic("error reading environment variables")
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	logger.Printf("loading facts for %s\n", cik)
	facts, err := facts.LoadFacts(cik, org, name, email)
	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s.json", cik)
	filePath := filepath.Join(folderPath, fileName)

	logger.Printf("uploading facts to %s on bucket %s\n", fileName, bucketName)

	err = storage.UploadBytes(facts, bucketName, filePath)
	if err != nil {
		panic(err)
	}
}
