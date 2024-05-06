#!/usr/bin/env bash

air &

AIR_PID=$!

cleanup() {
    kill -SIGINT $AIR_PID
    exit 0
}

trap cleanup EXIT SIGINT SIGTERM

wait $AIR_PID
