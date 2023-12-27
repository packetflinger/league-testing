package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	lpb "github.com/packetflinger/league-testing/proto"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	leagues, err := parseLeague()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Leagues:")
	for _, l := range leagues.GetLeague() {
		fmt.Println("  ", l.GetName())
	}

	league, err := fetchLeague("73f1fc07-fb0d-4741-9a9d-93df8aafc484", leagues)
	if err != nil {
		fmt.Println(err)
		return
	}
	matches, err := createGroupMatches(league)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, m := range matches {
		fmt.Println(m.GetTeamHome().GetName(), "vs", m.GetTeamAway().GetName())
	}
}

func parseLeague() (*lpb.Leagues, error) {
	filename := "league1.textpb"
	l := lpb.Leagues{}
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = prototext.Unmarshal(contents, &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

// Return a specific league proto from a UUID
func fetchLeague(uuid string, all *lpb.Leagues) (*lpb.League, error) {
	if all == nil {
		return nil, errors.New("input collection is nil")
	}
	for _, l := range all.GetLeague() {
		if l.GetUuid() == uuid {
			return l, nil
		}
	}
	return nil, errors.New("not found")
}

func createGroupMatches(league *lpb.League) ([]*lpb.Match, error) {
	matches := []*lpb.Match{}
	teams := league.GetTeam()
	for i := range teams {
		for j := range teams {
			var m *lpb.Match
			if teams[i].GetUuid() == teams[j].GetUuid() {
				continue
			}
			if matchContains(matches, teams[i], teams[j]) {
				continue
			}
			if i%2 == 0 {
				m = createGame(teams[i], teams[j])
			} else {
				m = createGame(teams[j], teams[i])
			}
			matches = append(matches, m)
		}
	}
	return matches, nil
}

func matchContains(collection []*lpb.Match, team1 *lpb.Team, team2 *lpb.Team) bool {
	for i := range collection {
		match := collection[i]
		if match.GetTeamHome().GetUuid() == team1.GetUuid() && match.GetTeamAway().GetUuid() == team2.GetUuid() {
			return true
		}
		if match.GetTeamHome().GetUuid() == team2.GetUuid() && match.GetTeamAway().GetUuid() == team1.GetUuid() {
			return true
		}
	}
	return false
}

func createGame(team1 *lpb.Team, team2 *lpb.Team) *lpb.Match {
	match := lpb.Match{
		TeamHome:    team1,
		TeamAway:    team2,
		TimeCreated: time.Now().Unix(),
	}
	return &match
}

/*
func shuffleTeams(t []*lpb.Team) []*lpb.Team {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(t), func(i, j int) { t[i], t[j] = t[j], t[i] })
	return t
}
*/
