#!/bin/bash
set -e

export ver=1 && \
    go test -run '^$' -bench '^BenchmarkSum$' -benchtime 10s -count 6 \
        -cpu 4 \
        -benchmem \
        -memprofile=${ver}.mem.pprof -cpuprofile=${ver}.cpu.pprof \
    | tee ${ver}.txt
