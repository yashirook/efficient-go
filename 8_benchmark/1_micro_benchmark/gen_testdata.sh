#!/bin/bash
set -e

for i in {1..2000000}; do echo $RANDOM; done > $1
