package facts

import (
	"fmt"
	"io"
	"net/http"
)

func LoadFacts(cik, name, org, email string) ([]byte, error) {
	url := fmt.Sprintf("https://data.sec.gov/api/xbrl/companyfacts/CIK%s.json", cik)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	userAgent := fmt.Sprintf("%s %s %s", org, name, email)
	req.Header.Set("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err := fmt.Errorf("status Code != ok: %v", res.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
