package _1773_count_items_matching_a_rule

import (
	"fmt"
	"testing"
)

func Test_countMatches(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		items     [][]string
		ruleKey   string
		ruleValue string
		want      int
	}{
		{
			items:     [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}},
			ruleKey:   "color",
			ruleValue: "silver",
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.items), func(t *testing.T) {
			if got := countMatches(tt.items, tt.ruleKey, tt.ruleValue); got != tt.want {
				t.Errorf("countMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
