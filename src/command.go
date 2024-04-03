package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func (ctx *AppContext) runDelegationVerifyCommand(command, input string) ([]Submission, error) {
	var cmd string
	if ctx.AppConfig.NoChecks {
		ctx.Log.Info("Note! Running with --no-checks flag. This will skip some checks.")
		cmd = fmt.Sprintf("%v stdin --no-checks", command)
	} else {
		cmd = fmt.Sprintf("%v stdin", command)
	}
	out, err := runCommand(cmd, input)
	if err != nil {
		return nil, fmt.Errorf("error running %v: %w", command, err)
	}
	submissions, err := parseDelegationVerifyOutput(out)
	if err != nil {
		return nil, fmt.Errorf("error parsing submissions: %w", err)
	}
	return submissions, nil
}

// Output from the delegation verification binary is expected to be a newline-separated JSON array of Submission objects.
// We parse this into a slice of Submission objects.
func parseDelegationVerifyOutput(data string) ([]Submission, error) {
	var submissions []Submission

	// Split the input data into separate records based on newline.
	records := strings.Split(data, "\n")

	for _, record := range records {
		if record == "" {
			continue // Skip empty lines
		}

		var submission Submission
		if err := json.Unmarshal([]byte(record), &submission); err != nil {
			return nil, err // Return error if any record fails to unmarshal
		}

		submissions = append(submissions, submission)
	}

	return submissions, nil
}

func runCommand(command, input string) (string, error) {
	cmdParts := strings.Split(command, " ")
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	cmd.Stdin = bytes.NewBufferString(input)

	// Run the command and capture its standard output.
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run command: %w", err)
	}

	return stdout.String(), nil
}
