package benchmarks

import "testing"

func BenchmarkLoadBalancer(b *testing.B) {
	connCount := 100
	triesCount := 1000
	connections := make([]*Connection, connCount)

	for i := 0; i < connCount; i++ {
		connections[i] = &Connection{}
	}
	ChanBalancer := NewLoadBalancerChan(connections)
	ChanBalancer.Init()
	defer ChanBalancer.Close()
	AtomicBalancer := NewLoadBalancerAtomic(connections)
	MitexBalancer := NewLoadBalancerMutex(connections)
	b.ResetTimer()

	b.Run("chan", func(b *testing.B) {
		for n := 0; n < triesCount; n++ {
			ChanBalancer.NextConn()
		}
	})
	b.Run("atomic", func(b *testing.B) {
		for n := 0; n < triesCount; n++ {
			AtomicBalancer.NextConn()
		}
	})
	b.Run("mutex", func(b *testing.B) {
		for n := 0; n < triesCount; n++ {
			MitexBalancer.NextConn()
		}
	})
}
