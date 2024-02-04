package _2115_find_all_possible_recipes_from_given_suppliers

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	availableIngredients := make(map[string]bool, len(supplies))

	for _, supply := range supplies {
		availableIngredients[supply] = true
	}

	recipeToIDx := make(map[string]int, len(recipes))
	for i, recipe := range recipes {
		recipeToIDx[recipe] = i
	}

	var res []string
	for _, recipe := range recipes {
		if possible(recipe, recipeToIDx, ingredients, availableIngredients, make(map[string]bool)) {
			res = append(res, recipe)
		}
	}

	return res
}

func possible(
	recipe string,
	recipeToIDx map[string]int,
	ingredients [][]string,
	availableIngredients map[string]bool,
	visited map[string]bool,
) bool {
	if visited[recipe] {
		return false
	}

	visited[recipe] = true

	if ok, keyExists := availableIngredients[recipe]; keyExists {
		return ok
	}

	idx, keyExists := recipeToIDx[recipe]
	if !keyExists {
		return false
	}

	res := true
	for _, ingredient := range ingredients[idx] {
		if !possible(ingredient, recipeToIDx, ingredients, availableIngredients, visited) {
			res = false
		}
	}

	availableIngredients[recipe] = res

	return res
}
