package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	ans := []BrainrotMeme{}

	for _, i := range memes {
		if i.Views > minViews {
			ans = append(ans, i)
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].TrendLevel > ans[j].TrendLevel
	})

	return ans
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	m := make(map[string]float64)

	for _, meme := range memes {
		m[meme.Category] += meme.Views
	}

	return m
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var ans []string

	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			ans = append(ans, meme.Name)
		}
	}

	return ans
}

func main() {
	memes := []BrainrotMeme{
		{
			Name:       "Skibidi Toilet 76",
			TrendLevel: 9,
			Category:   "Skibidi",
			Views:      245.3,
		},
		{
			Name:       "Sigma Male Grindset",
			TrendLevel: 8,
			Category:   "Sigma",
			Views:      178.9,
		},
		{
			Name:       "Mewing Challenge",
			TrendLevel: 7,
			Category:   "Mewing",
			Views:      132.5,
		},
		{
			Name:       "Subo Bratik Compilation",
			TrendLevel: 10,
			Category:   "Subo Bratik",
			Views:      310.7,
		},
		{
			Name:       "TUNTUNTUN Morning Routine",
			TrendLevel: 6,
			Category:   "TUNTUNTUNSAHUR",
			Views:      89.4,
		},
		{
			Name:       "Gigachad Omega",
			TrendLevel: 8,
			Category:   "Other",
			Views:      156.2,
		},
		{
			Name:       "Fanum Taxing",
			TrendLevel: 9,
			Category:   "Other",
			Views:      221.8,
		},
	}

	fmt.Println("ТОП-ТРЕНДЫ (просмотры > 150 млн)")
	fmt.Print(FindTopTrending(memes, 150.0))
	fmt.Println()
	fmt.Println()

	fmt.Println("ВЛИЯНИЕ КАТЕГОРИЙ (суммарные просмотры)")
	fmt.Print(CalculateCategoryImpact(memes))
	fmt.Println()
	fmt.Println()

	fmt.Println("МЕМЫ ПО СЛОЖНОМУ УСЛОВИЮ")
	fmt.Print(FilterByComplexCondition(memes))
	fmt.Println()
	fmt.Println()

}
