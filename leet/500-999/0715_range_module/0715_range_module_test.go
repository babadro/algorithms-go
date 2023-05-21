package _0715_range_module

import (
	"testing"

	"github.com/babadro/algorithms-go/utils"
	"github.com/stretchr/testify/require"
)

func TestRangeModule1(t *testing.T) {
	r := Constructor()

	r.AddRange(10, 20)
	r.RemoveRange(14, 16)
	t.Log(r.QueryRange(10, 14))
	t.Log(r.QueryRange(13, 15))
	t.Log(r.QueryRange(16, 17))
}

func TestRangeModule_(t *testing.T) {
	tests := []struct {
		ops    []*string
		input  []*[]int
		output []*bool
	}{
		{
			utils.ParseArr[string](`["RangeModule", "addRange", "removeRange", "queryRange", "queryRange", "queryRange"]`),
			utils.ParseArr[[]int](`[[], [10, 20], [14, 16], [10, 14], [13, 15], [16, 17]]`),
			utils.ParseArr[bool](`[null, null, null, true, false, true]`),
		},
		{
			utils.ParseArr[string](`["RangeModule","queryRange","queryRange","addRange","addRange","queryRange","queryRange","queryRange","removeRange","addRange","removeRange","addRange","removeRange","removeRange","queryRange","queryRange","queryRange","queryRange","removeRange","addRange","removeRange","queryRange","addRange","addRange","removeRange","queryRange","removeRange","removeRange","removeRange","addRange","removeRange","addRange","queryRange","queryRange","queryRange","queryRange","queryRange","addRange","removeRange","addRange","addRange","addRange","queryRange","removeRange","addRange","queryRange","addRange","queryRange","removeRange","removeRange","addRange","addRange","queryRange","queryRange","addRange","addRange","removeRange","removeRange","removeRange","queryRange","removeRange","removeRange","addRange","queryRange","removeRange","addRange","addRange","queryRange","removeRange","queryRange","addRange","addRange","addRange","addRange","addRange","removeRange","removeRange","addRange","queryRange","queryRange","removeRange","removeRange","removeRange","addRange","queryRange","removeRange","queryRange","addRange","removeRange","removeRange","queryRange"]`),
			utils.ParseArr[[]int](`[[],[21,34],[27,87],[44,53],[69,89],[23,26],[80,84],[11,12],[86,91],[87,94],[34,52],[1,59],[62,96],[34,83],[11,59],[59,79],[1,13],[21,23],[9,61],[17,21],[4,8],[19,25],[71,98],[23,94],[58,95],[12,78],[46,47],[50,70],[84,91],[51,63],[26,64],[38,87],[41,84],[19,21],[18,56],[23,39],[29,95],[79,100],[76,82],[37,55],[30,97],[1,36],[18,96],[45,86],[74,92],[27,78],[31,35],[87,91],[37,84],[26,57],[65,87],[76,91],[37,77],[18,66],[22,97],[2,91],[82,98],[41,46],[6,78],[44,80],[90,94],[37,88],[75,90],[23,37],[18,80],[92,93],[3,80],[68,86],[68,92],[52,91],[43,53],[36,37],[60,74],[4,9],[44,80],[85,93],[56,83],[9,26],[59,64],[16,66],[29,36],[51,96],[56,80],[13,87],[42,72],[6,56],[24,53],[43,71],[36,83],[15,45],[10,48]]`),
			utils.ParseArr[bool](`[null,false,false,null,null,false,true,false,null,null,null,null,null,null,false,false,true,true,null,null,null,false,null,null,null,false,null,null,null,null,null,null,true,true,false,false,false,null,null,null,null,null,true,null,null,false,null,true,null,null,null,null,false,false,null,null,null,null,null,false,null,null,null,false,null,null,null,true,null,false,null,null,null,null,null,null,null,null,false,false,null,null,null,null,true,null,false,null,null,null,false]`),
		},
		{
			utils.ParseArr[string](`["RangeModule","addRange","addRange","removeRange","queryRange","queryRange","removeRange","removeRange","removeRange","removeRange","removeRange","queryRange","removeRange","addRange","removeRange","addRange","queryRange","queryRange","addRange","addRange","queryRange","removeRange","queryRange","addRange","queryRange","removeRange","removeRange","addRange","addRange","removeRange","removeRange","removeRange","addRange","addRange","queryRange","queryRange","queryRange","queryRange","queryRange","removeRange","removeRange","queryRange","addRange","addRange","addRange","queryRange","addRange","addRange","removeRange","addRange","queryRange","removeRange","addRange","queryRange","addRange","addRange","addRange","queryRange","addRange","queryRange","removeRange","removeRange","removeRange","removeRange","queryRange","removeRange","queryRange","queryRange","removeRange","queryRange","addRange","addRange","queryRange","removeRange","removeRange","queryRange","addRange","removeRange","removeRange","addRange","addRange","addRange","queryRange","queryRange","addRange","queryRange","removeRange","queryRange","removeRange","addRange","queryRange"]`),
			utils.ParseArr[[]int](`[[],[55,62],[1,29],[18,49],[6,98],[59,71],[40,45],[4,58],[57,69],[20,30],[1,40],[73,93],[32,93],[38,100],[50,64],[26,72],[8,74],[15,53],[44,85],[10,71],[54,70],[10,45],[30,66],[47,98],[1,7],[44,78],[31,49],[62,63],[49,88],[47,72],[8,50],[49,79],[31,47],[54,87],[77,78],[59,100],[8,9],[50,51],[67,93],[25,86],[8,92],[31,87],[90,95],[28,56],[10,42],[27,34],[75,81],[17,63],[78,90],[9,18],[51,74],[20,54],[35,72],[2,29],[28,41],[17,95],[73,75],[34,43],[57,96],[51,72],[21,67],[40,73],[14,26],[71,86],[34,41],[10,25],[27,68],[18,32],[30,31],[45,61],[64,66],[18,93],[13,21],[13,46],[56,99],[6,93],[25,36],[27,88],[82,83],[30,71],[31,73],[10,41],[71,72],[9,56],[22,76],[38,74],[2,77],[33,61],[74,75],[11,43],[27,75]]`),
			utils.ParseArr[bool](`[null,null,null,null,false,false,null,null,null,null,null,false,null,null,null,null,false,false,null,null,true,null,false,null,false,null,null,null,null,null,null,null,null,null,true,true,false,false,true,null,null,false,null,null,null,true,null,null,null,null,false,null,null,false,null,null,null,true,null,true,null,null,null,null,false,null,false,false,null,false,null,null,false,null,null,false,null,null,null,null,null,null,true,true,null,true,null,false,null,null,false]`),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			require.Equal(t, len(tt.ops), len(tt.input))

			r := Constructor()

			for i, op := range tt.ops {
				require.NotNil(t, op)
				input := tt.input[i]
				switch *op {
				case "RangeModule":
					continue
				case "addRange":
					r.AddRange((*input)[0], (*input)[1])
				case "removeRange":
					r.RemoveRange((*input)[0], (*input)[1])
				case "queryRange":
					want := tt.output[i]
					require.NotNil(t, want)

					require.NotPanics(t, func() {
						got := r.QueryRange((*input)[0], (*input)[1])
						require.Equalf(t, *want, got, "%d, input# %d", *input, i)
					})

				default:
					t.Errorf("unknown operation %s", *op)
				}
			}
		})
	}
}
