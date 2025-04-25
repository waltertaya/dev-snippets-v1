package main

import "fmt"

type Movie struct {
	Id     string `json:"id"`
	Title  string
	Year   int `json:"released"`
	Actors []string
}

func main() {

	movies := []Movie{
		{Id: "tyunmode64ndh7ww", Title: "Casablanca", Year: 2002, Actors: []string{"Denise Huskins", "Brett Cooper"}},
		{Id: "gshw8js90qnsj12n", Title: "Bullitt", Year: 2001, Actors: []string{"Matthew Muller", "Jacqueline Bisset", "Steve McQueen"}},
		{Id: "tyunmode64ndh7ww", Title: "Casablanca", Year: 2002, Actors: []string{"Denise Huskins", "Brett Cooper"}},
		{Id: "gshw8js90qnsj12n", Title: "Bullitt", Year: 2001, Actors: []string{"Matthew Muller", "Jacqueline Bisset", "Steve McQueen"}},
		{Id: "tyunmode64ndh7ww", Title: "Casablanca", Year: 2002, Actors: []string{"Denise Huskins", "Brett Cooper"}},
		{Id: "gshw8js90qnsj12n", Title: "Bullitt", Year: 2001, Actors: []string{"Matthew Muller", "Jacqueline Bisset", "Steve McQueen"}},
		{Id: "tyunmode64ndh7ww", Title: "Casablanca", Year: 2002, Actors: []string{"Denise Huskins", "Brett Cooper"}},
		{Id: "gshw8js90qnsj12n", Title: "Bullitt", Year: 2001, Actors: []string{"Matthew Muller", "Jacqueline Bisset", "Steve McQueen"}},
	}

	fmt.Println(movies)
	fmt.Println(len(movies))
}
