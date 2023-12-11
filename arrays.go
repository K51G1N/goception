package main

import "fmt"

func main() {

	// var team_avatar_ages [3]int = [3]int{10, 12, 14}
	var teamAvatarAges = [3]int{10, 12, 14} //shortened
	names := [4]string{"Aang", "Katara", "Sokka", "Toph"}
	names[2] = "Zuko"

	fmt.Println(teamAvatarAges, len(teamAvatarAges) )
	fmt.Println(names, len(names))

	//slices use array. can be manipulated

	var scores = []int{20, 30, 40}
	scores[2] = 25
	scores = append(scores, 60)

	fmt.Println(scores, len(scores))

	// slice ranges

	nonAvatarMembers := names[1:4]
	fmt.Println(nonAvatarMembers)

	earthAndFireBenders := names[2:]

	excludeToph := names[:3]

	fmt.Println(earthAndFireBenders)

	fmt.Println(excludeToph)
}