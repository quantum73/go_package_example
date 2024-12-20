package env

import (
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

// TestParseInt calls env.ParseInt with incorrect string.
func TestParseIntWithIncorrectString(t *testing.T) {
	incorrectString := "foo"
	_, err := ParseInt(incorrectString)
	if err == nil {
		t.Fatalf("env.ParseInt(%s) got no error", incorrectString)
	}
}

// TestGetRequiredEnvValue calls env.GetRequiredValue with existed environment key.
func TestGetRequiredEnvValue(t *testing.T) {
	// Set environment variable
	envVariableName := "HOST"
	envVariableValue := "127.0.0.1"
	err := os.Setenv(envVariableName, envVariableValue)
	if err != nil {
		t.Fatalf("error setting environment variable: %s", err)
	}

	t.Run("Test GetRequiredValue with existed env variable", func(t *testing.T) {
		targetEnvVariableValue, err := GetRequiredValue(envVariableName)
		if err != nil {
			t.Fatalf("env.GetRequiredValue(\"%s\") got error: %s", envVariableName, err)
		}
		if targetEnvVariableValue != envVariableValue {
			t.Fatalf(
				"env.GetRequiredValue(\"%s\") = %s, expected - %s",
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

// TestGetRequiredEnvValueWithoutEnv calls env.GetRequiredValue with not existed environment key.
func TestGetRequiredEnvValueWithoutEnv(t *testing.T) {
	envVariableName := "HOST"
	_, err := GetRequiredValue(envVariableName)
	if err == nil {
		t.Fatalf("env.GetRequiredValue(\"%s\") should return an error", envVariableName)
	}
}
