package bench_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/exaring/hoglet"
	circuit "github.com/rubyist/circuitbreaker"
	"github.com/sony/gobreaker"
)

func BenchmarkHoglet(b *testing.B) {
	circuit, err := hoglet.NewCircuit(
		func(context.Context, any) (any, error) { return nil, nil },
		hoglet.NewEWMABreaker(10, 0.9),
		hoglet.WithHalfOpenDelay(time.Second),
		//        hoglet.WithConcurrencyLimit(100_000_000, false),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = circuit.Call(ctx, nil)
		}
	})
}

func BenchmarkHystrix(b *testing.B) {
	cmdName := "foo"

	hystrix.ConfigureCommand(cmdName, hystrix.CommandConfig{
		Timeout: 250,
		//		MaxConcurrentRequests:  100_000_000, // unlimited; otherwise we skew the comparison
		ErrorPercentThreshold:  10,
		RequestVolumeThreshold: 20,
		SleepWindow:            2000,
	})

	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = hystrix.DoC(ctx, cmdName, func(context.Context) error { return nil }, nil)
		}
	})
}

func BenchmarkGoBreaker(b *testing.B) {
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name: "foo",
		// MaxRequests: 100_000_000, // unlimited; otherwise we skew the comparison
		Interval: 10 * time.Second,
		Timeout:  time.Second,
	})

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = cb.Execute(func() (interface{}, error) { return nil, nil })
		}
	})
}

func BenchmarkCircuit(b *testing.B) {
	cb := circuit.NewRateBreaker(0.9, 10)

	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// why doesn't it use the context's timeout? 🤷
			_ = cb.CallContext(ctx, func() error { return nil }, time.Second)
		}
	})
}
