package context

import (
	"context"
	"log"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	// cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)

			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}

}

func (s *SpyStore) Cancel() {
	// s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	// if !s.cancelled {
	// 	s.t.Error("store was not told to cancel")
	// }
}
