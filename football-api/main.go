package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
)

/*
 * Complete the 'eliteClubs' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING nation
 *  2. INTEGER minValuation
 *  3. INTEGER minTitlesWon
 * API URL: https://jsonmock.hackerrank.com/api/football_teams?nation=<nation>
 */

var client = &http.Client{}

type FootballTeam struct {
	Name                    string `json:"name"`
	EstimatedValueNumeric   int64  `json:"estimated_value_numeric"`
	NumberOfLeagueTitlesWon int64  `json:"number_of_league_titles_won"`
}

type Response struct {
	Page       int32          `json:"page"`
	TotalPages int32          `json:"total_pages"`
	Total      int32          `json:"total"`
	Data       []FootballTeam `json:"data"`
}

type TeamsPage struct {
	Teams      []FootballTeam
	Page       int32
	IsLastPage bool
}

func getPage(nation string, page int32) TeamsPage {
	p := strconv.Itoa(int(page))
	url := "https://jsonmock.hackerrank.com/api/football_teams?nation=" + nation + "&page=" + p
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var pageResp Response
	err = json.NewDecoder(resp.Body).Decode(&pageResp)
	if err != nil {
		panic(err)
	}

	return TeamsPage{
		Teams:      pageResp.Data,
		Page:       pageResp.Page,
		IsLastPage: pageResp.Page > pageResp.TotalPages,
	}
}

func eliteClubs(nation string, minValuation int32, minTitlesWon int32) []string {
	var eliteClubs []FootballTeam
	page := getPage(nation, 1)
	for {
		for _, team := range page.Teams {
			if team.EstimatedValueNumeric >= int64(minValuation) && team.NumberOfLeagueTitlesWon >= int64(minTitlesWon) {
				eliteClubs = append(eliteClubs, team)
			}
		}
		if page.IsLastPage {
			break
		}
		page = getPage(nation, page.Page+1)
	}
	// sort by valuation then ascending name
	sort.SliceStable(eliteClubs, func(i, j int) bool {
		return eliteClubs[i].Name < eliteClubs[j].Name
	})
	sort.SliceStable(eliteClubs, func(i, j int) bool {
		return eliteClubs[i].EstimatedValueNumeric > eliteClubs[j].EstimatedValueNumeric
	})

	var eliteClubsNames []string
	for _, team := range eliteClubs {
		eliteClubsNames = append(eliteClubsNames, team.Name)
	}
	return eliteClubsNames
}

func main() {
	// Call the function with the required inputs
	eliteClubs("england", 823472, 2)
}
