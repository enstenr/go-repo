go test -bench=. -benchmem algo_dynami_program.go fibo_test.go 

go test -bench=. -benchmem algo3_quicksort.go algo3_test.go

go test -bench=. -benchmem algo3_quicksort_2.go algo3_quicksort_2_test.go

go test -bench=. -benchmem c3.go c3_test.go

go test -bench=BenchmarkNewAeroplane -cpuprofile=cpu.pprof


go tool pprof -http=:8080 cpu.pprof

go tool pprof cpu.pprof
(pprof) top10