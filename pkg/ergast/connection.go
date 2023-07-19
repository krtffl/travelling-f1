package ergast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func seasonUrl(year int) string {
	return fmt.Sprintf("%s/%d.json", baseUrl, year)
}

func GetSeason(year int) RaceTable {
	resp, err := http.Get(seasonUrl(year))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}
	return response.MRData.RaceTable
}
