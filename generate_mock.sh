#!/bin/bash
FILE_PATH="internal/twilio/twilio-sdk"

mkdir -p "$FILE_PATH/mock"
mockgen -source="$FILE_PATH/twilio.go" -destination "$FILE_PATH/mock/twilio_mock.go"