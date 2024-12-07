package env

import (
	"fmt"
	"os"
	"testing"
)

// TestParseInt calls env.ParseInt with a number as string.
func TestParseInt(t *testing.T) {
	intAsString := "73"
	expectedIntNumber := 73
	intNumber, err := ParseInt(intAsString)
	if err != nil {
		t.Fatalf("env.ParseInt(%s) got error: %s", intAsString, err)
	}
	if intNumber != expectedIntNumber {
		t.Fatalf(
			"env.ParseInt(%s) = %d, expected - %d",
			intAsString, intNumber, expectedIntNumber,
		)
	}
}

// TestGetRequiredEnvValue calls env.GetRequiredEnvValue with existed environment key.
func TestGetRequiredEnvValue(t *testing.T) {
	// Set environment variable
	envVariableName := "HOST"
	envVariableValue := "127.0.0.1"
	err := os.Setenv(envVariableName, envVariableValue)
	if err != nil {
		t.Fatalf("error setting environment variable: %s", err)
	}

	t.Run("Test GetRequiredEnvValue with existed env variable", func(t *testing.T) {
		targetEnvVariableValue, err := GetRequiredEnvValue(envVariableName)
		if err != nil {
			t.Fatalf("env.GetRequiredEnvValue(\"%s\") got error: %s", envVariableName, err)
		}
		if targetEnvVariableValue != envVariableValue {
			t.Fatalf(
				"env.GetRequiredEnvValue(\"%s\") = %s, expected - %s",
				envVariableName, targetEnvVariableValue, envVariableValue,
			)
		}
	})

	// Unset environment variable
	err = os.Unsetenv(envVariableName)
	if err != nil {
		t.Fatalf("error setting environment variable: %s", err)
	}
}

// TestGetRequiredEnvValueWithoutEnv calls env.GetRequiredEnvValue with not existed environment key.
func TestGetRequiredEnvValueWithoutEnv(t *testing.T) {
	envVariableName := "HOST"
	expectedErrorMessage := fmt.Sprintf("environment variable `%s` unset", envVariableName)

	_, err := GetRequiredEnvValue(envVariableName)
	if err == nil {
		t.Fatalf("env.GetRequiredEnvValue(\"%s\") should return an error", envVariableName)
	}

	if err.Error() != expectedErrorMessage {
		t.Fatalf("Error message should be `%s`, but got `%s`", expectedErrorMessage, err.Error())
	}
}