package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/nthskyradiated/go-edgarfacts/edgarfactsfunction/internal/data"
	"github.com/nthskyradiated/go-edgarfacts/edgarfactsfunction/internal/facts"
	"github.com/nthskyradiated/go-edgarfacts/edgarfactsfunction/internal/storage"
)

func init() {
	functions.HTTP("LoadFacts", LoadFactHandler)
}

func LoadFactHandler(w http.ResponseWriter, r *http.Request) {
	var inputUser data.FactsLoaderHandlerInput
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "error: could not parse user input.")
		return
	}

	if len(inputUser.Cik) != 10 {
		w.WriteHeader(400)
		fmt.Fprint(w, "cik must be 10 characters long")
		return
	}

	if inputUser.Org == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "please provide your Organization's name")
		return
	}

	if inputUser.Name == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "please provide your name")
		return
	}

	if inputUser.Email == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "email address required")
	}

	bucketName := os.Getenv("BUCKET_NAME")
	folderPath := os.Getenv("STAGE")

	if bucketName == "" || folderPath == "" {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal Error")
		return
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	logger.Printf("loading facts for %s\n", inputUser.Cik)
	facts, err := facts.LoadFacts(inputUser.Cik, inputUser.Org, inputUser.Name, inputUser.Email)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal Error")
		return
	}

	fileName := fmt.Sprintf("%s.json", inputUser.Cik)
	filePath := filepath.Join(folderPath, fileName)

	logger.Printf("uploading facts to %s on bucket %s\n", fileName, bucketName)

	err = storage.UploadBytes(facts, bucketName, filePath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal Error")
	}

	fmt.Fprint(w, "success")

}
