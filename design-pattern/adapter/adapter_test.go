package adapter

import (
	"testing"
)

type Client struct {
	scoreOpt ScoreOperation
}

func (c *Client) ListScores(scores []float32) []float32 {
	return c.scoreOpt.Sort(scores)
}

func TestStructAdapter(t *testing.T) {
	scores := []float32{2.0, 3.0, 1.0}
	qka := NewQuickSortStructAdapter()
	client := &Client{scoreOpt: qka}
	scores = client.ListScores(scores)
	if scores[0] != 1.0 {
		t.Error("failed to list scores by sort")
	}
}

func TestObjectAdapter(t *testing.T) {
	scores := []float32{2.0, 3.0, 1.0}
	qka := NewQuickSortObjectAdapter()
	client := &Client{scoreOpt: qka}
	scores = client.ListScores(scores)
	if scores[0] != 1.0 {
		t.Error("failed to list scores by sort")
	}
}
