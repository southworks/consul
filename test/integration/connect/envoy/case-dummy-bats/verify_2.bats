#!/usr/bin/env bats

load dummy-function

setup() {
  source $(pwd)/workdir/vars.sh
  source $(pwd)/workdir/setup.sh "Content of the created setup.txt file in setup.sh" $TXT_FILE_NAME
}

teardown() {
  rm $TXT_FILE_NAME
}

@test "Test with dummyFunction invoked" {
  FIRST_ARG="First Argument"

  run dummyFunction "$FIRST_ARG"

  [ $status -eq 0 ]
  [ -n "$output" ] # Not empty
  [ "$output" = "$FIRST_ARG $COMMON_VAR" ]
}

@test "Test skipped" {
  skip

  run not_existing_function

  [ "$status" -eq 100000 ]
}
