package benchmarks

import "testing"

func Benchmark_MarshalLimitClause(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q, err := marshalLimitClause(5)
		if err != nil {
			b.Errorf("failed: %s", err.Error())
		}
		if q != "5" {
			b.Errorf("wrong query: %s", q)
		}
	}
}

func Benchmark_MarshalLimitClauseAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q, err := marshalLimitClauseAlt(5)
		if err != nil {
			b.Errorf("failed: %s", err.Error())
		}
		if q != "5" {
			b.Errorf("wrong query: %s", q)
		}
	}
}
