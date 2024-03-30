#!/bin/bash
set -e

OUTPUT_DIR=output

export ver=v1 && \
    go test -run '^$' -bench '^BenchmarkSum$' -benchtime 10s -count 6 \
        -cpu 4 \
        -benchmem \
        -memprofile=$OUTPUT_DIR/${ver}.mem.pprof \
        -cpuprofile=$OUTPUT_DIR/${ver}.cpu.pprof \
    | tee $OUTPUT_DIR/${ver}.txt
