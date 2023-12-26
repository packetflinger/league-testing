package main

import (
	"fmt"
	"os"

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
