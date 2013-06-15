package nqueens

import (
	"testing"
)

const nQueens = 14

func BenchmarkSequential(b *testing.B) {
	solutions := SolveSequential(nQueens)
	b.Logf("%d solutions found for %d queens.", solutions, nQueens)
}

func BenchmarkParallel(b *testing.B) {
	solutions := SolveParallel(nQueens)
	b.Logf("%d solutions found for %d queens.", solutions, nQueens)
}

func TestSequential(t *testing.T) {
	solutions := SolveSequential(8)
	if solutions != 92 {
		t.Fail()
	}
}

func TestParallel(t *testing.T) {
	solutions := SolveParallel(8)
	if solutions != 92 {
		t.Fail()
	}
}
