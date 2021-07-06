package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

type Config struct {
	outputFile   string
	outputStdout bool
}

func parseArgs() (*Config, []string) {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		slog.Error("wrong arguments")
		os.Exit(1)
	}

	config := Config{outputStdout: true, outputFile: ""}
	if len(os.Args) == 4 {
		config.outputStdout = false
		config.outputFile = os.Args[3]
	}

	fileNames := []string{os.Args[1]}

	if len(os.Args) > 2 {
		fileNames = append(fileNames, os.Args[2])
	}

	return &config, fileNames
}

func parseFiles(fileNames []string) ([]string, error) {
	result := []string{}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			return []string{}, fmt.Errorf("error during open file %w", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			tmpStr := scanner.Text()
			fmt.Println(tmpStr)
			result = append(result, tmpStr)
		}
	}
	return result, nil
}

func writeOutput(c *Config, lines []string) {
	stringResult := strings.Join(lines, "\n")

	fmt.Printf("%s", lines)
	if c.outputStdout {
		fmt.Printf("%s", stringResult)
	} else {
		d1 := []byte(stringResult)
		err := ioutil.WriteFile(c.outputFile, d1, 0o644)
		if err != nil {
			slog.Fatal(err)
		}
	}
}

func main() {
	slog = logger.NewSugaredLogger()
	slog.Sync() //nolint:errcheck

	slog.Debug(len(os.Args))

	config, fileNames := parseArgs()

	result, err := parseFiles(fileNames)
	if err != nil {
		slog.Error(err)
		os.Exit(1)
	}

	writeOutput(config, result)
}
