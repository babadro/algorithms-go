package _1773_count_items_matching_a_rule

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var filter func(item []string, value string) bool
	switch ruleKey {
	case "type":
		filter = matchType
	case "color":
		filter = matchColor
	default:
		filter = matchName
	}

	counter := 0
	for _, item := range items {
		if filter(item, ruleValue) {
			counter++
		}
	}

	return counter
}

func matchType(item []string, typ string) bool {
	return item[0] == typ
}

func matchColor(item []string, color string) bool {
	return item[1] == color
}

func matchName(item []string, name string) bool {
	return item[2] == name
}
