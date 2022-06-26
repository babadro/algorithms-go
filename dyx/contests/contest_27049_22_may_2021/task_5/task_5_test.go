package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_checkRequests(t *testing.T) {
	tests := []struct {
		userLimit    int
		serviceLimit int
		duration     int
		reqs         [][2]int
		wantStatuses []int
	}{
		{
			userLimit: 2, serviceLimit: 5, duration: 5,
			reqs:         [][2]int{{1, 100}, {1, 100}, {2, 100}, {2, 200}, {2, 300}, {2, 400}, {2, 500}, {3, 500}, {5, 200}, {6, 100}, {7, 200}},
			wantStatuses: []int{200, 200, 429, 200, 200, 200, 503, 503, 503, 429, 200},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d %d", tt.userLimit, tt.serviceLimit, tt.duration), func(t *testing.T) {
			reqs, statuses := make(chan [2]int), make(chan int)
			go checkRequestsChan(tt.userLimit, tt.serviceLimit, tt.duration, reqs, statuses)

			gotStatuses := make([]int, 0, len(tt.reqs))
			for _, req := range tt.reqs {
				reqs <- req
				gotStatuses = append(gotStatuses, <-statuses)
			}

			close(reqs)

			if !reflect.DeepEqual(gotStatuses, tt.wantStatuses) {
				t.Errorf("checkRequests() = %v, want %v", gotStatuses, tt.wantStatuses)
			}
		})
	}
}
