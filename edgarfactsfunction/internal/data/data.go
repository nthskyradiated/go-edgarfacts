package data

type FactsLoaderHandlerInput struct {
	Cik   string `json:"cik"`
	Org   string `json:"org"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
