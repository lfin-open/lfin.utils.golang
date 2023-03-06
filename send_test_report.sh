#! /bin/bash

# run test coverage and send to deepsource

# ----------------------
# test coverage
# https://deepsource.io/docs/analyzer/test-coverage/
# ----------------------
# Set DEEPSOURCE_DSN env variable from repository settings page
export DEEPSOURCE_DSN=https://333f2cfad267446796f19387eb6b10f1@deepsource.io

# test coverage, report to deepsource
go test -coverprofile=cover.out ./...

# Install 'deepsource CLI' if not installed
if [ -f ./bin/deepsource ]; then
    echo "deepsource CLI already installed"
else
    echo "deepsource CLI not installed, installing now"
    curl https://deepsource.io/cli | sh
fi


# From the root directory, run the report coverage command
./bin/deepsource report --analyzer test-coverage --key go --value-file ./cover.out
# -----------------------
