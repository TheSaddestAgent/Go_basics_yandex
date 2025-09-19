package main

import (
	"cmp"
	"math"
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

var playersCollection = make([]Player, 0)

func NewPlayer(name string, goals, misses, assists int) Player {
	var rating float64 = calculateRating(goals, misses, assists)

	p := Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: rating}
	AddPlayer(p)
	return p
}
func calculateRating(goals, misses, assists int) float64 {
	var res float64 = (float64(goals) + float64(assists)/2) / (math.Max(1, float64(misses)))
	return res
}
func AddPlayer(player Player) {

	playersCollection = append(playersCollection, player)
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a Player, b Player) int {
		if cmpResult := cmp.Compare(b.Goals, a.Goals); cmpResult != 0 {
			return cmpResult
		}
		return cmp.Compare(a.Name, b.Name)
	})
	return players
}
func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a Player, b Player) int {
		if cmpResult := cmp.Compare(b.Rating, a.Rating); cmpResult != 0 {
			return cmpResult
		}
		return cmp.Compare(a.Name, b.Name)
	})
	return players
}
func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a Player, b Player) int {
		if cmpResult := cmp.Compare(float64(b.Goals)/math.Max(1, float64(b.Misses)), float64(a.Goals)/math.Max(1, float64(a.Misses))); cmpResult != 0 {
			return cmpResult
		}
		return cmp.Compare(a.Name, b.Name)
	})
	return players
}
