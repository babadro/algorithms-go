package _2115_find_all_possible_recipes_from_given_suppliers

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAllRecipes(t *testing.T) {
	tests := []struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
		want        []string
	}{
		{
			recipes: []string{"bread"}, ingredients: [][]string{{"yeast", "flour"}}, supplies: []string{"yeast", "flour", "corn"},
			want: []string{"bread"},
		},
		{
			recipes: []string{"bread", "sandwich"}, ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}},
			supplies: []string{"yeast", "flour", "meat"},
			want:     []string{"bread", "sandwich"},
		},
		{
			recipes:     []string{"bread", "sandwich", "burger"},
			ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
			supplies:    []string{"yeast", "flour", "meat"},
			want:        []string{"bread", "sandwich", "burger"},
		},
		{
			recipes: recipes1, ingredients: ingredients1, supplies: supplies1,
			want: []string{"fzjnm", "ju"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findAllRecipes(tt.recipes, tt.ingredients, tt.supplies)

			sort.Strings(got)
			sort.Strings(tt.want)

			require.Equal(t, tt.want, got)
		})
	}
}
