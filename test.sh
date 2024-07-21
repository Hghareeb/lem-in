#!/bin/bash

RESULTS_DIR="results"
mkdir -p $RESULTS_DIR

# Function to run a test case
run_test() {
    local input_file=$1
    local expected_turns=$2
    local description=$3
    local result_file="$RESULTS_DIR/$(basename "$input_file" .txt)_result.txt"

    echo "Running test: $description" | tee $result_file
    output=$(go run . "$input_file")
    actual_turns=$(echo "$output" | grep -o "^L" | wc -l)

    echo "$output" | tee -a $result_file

    if [ "$actual_turns" -le "$expected_turns" ]; then
        echo "PASS: $description" | tee -a $result_file
    else
        echo "FAIL: $description (Expected at most $expected_turns turns, but got $actual_turns)" | tee -a $result_file
    fi

    echo | tee -a $result_file
}

# Check if the Go program can be run
if ! go build -o lem-in .; then
    echo "Error: Failed to build the Go program." | tee -a $RESULTS_DIR/build_error.txt
    exit 1
fi

# Run all test cases in the examples_run folder
for input_file in examples_run/*; do
    result_file="$RESULTS_DIR/$(basename "$input_file" .txt)_result.txt"
    if [[ $input_file == *"badexample"* ]]; then
        echo "Running test: $input_file" | tee $result_file
        output=$(./lem-in "$input_file")
        if [[ $output == *"ERROR: invalid data format"* ]]; then
            echo "PASS: $input_file" | tee -a $result_file
        else
            echo "FAIL: $input_file (Expected error message)" | tee -a $result_file
        fi
    else
        echo "Running test: $input_file" | tee $result_file
        output=$(./lem-in "$input_file")
        echo "$output" | tee -a $result_file
    fi
    echo | tee -a $result_file
done

# Additional test for performance with larger examples
if [ -f "examples_run/example06.txt" ]; then
    result_file="$RESULTS_DIR/example06_result.txt"
    echo "Running performance test with example06 (100 ants)" | tee $result_file
    time ./lem-in examples_run/example06.txt 2>&1 | tee -a $result_file
fi

if [ -f "examples_run/example07.txt" ]; then
    result_file="$RESULTS_DIR/example07_result.txt"
    echo "Running performance test with example07 (1000 ants)" | tee $result_file
    time ./lem-in examples_run/example07.txt 2>&1 | tee -a $result_file
fi

# Cleanup
rm -f lem-in

echo "All tests completed."
