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

func main() {
	slog = logger.NewSugaredLogger()
	slog.Sync() //nolint:errcheck

	slog.Info(len(os.Args))

	if len(os.Args) < 2 || len(os.Args) > 4 {
		slog.Error("wrong arguments")
		return
	}

	outputStdout := true
	outputFile := ""
	if len(os.Args) == 4 {
		outputStdout = false
		outputFile = os.Args[3]
	}

	fileNames := []string{os.Args[1]}

	if len(os.Args) > 2 {
		fileNames = append(fileNames, os.Args[2])
	}

	result := []string{}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			slog.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			tmpStr := scanner.Text()
			fmt.Println(tmpStr)
			result = append(result, tmpStr)
		}
	}

	stringResult := strings.Join(result, "\n")

	fmt.Printf("%s", result)
	if outputStdout {
		fmt.Printf("%s", stringResult)
	} else {
		d1 := []byte(stringResult)
		err := ioutil.WriteFile(outputFile, d1, 0o644)
		if err != nil {
			slog.Fatal(err)
		}
	}
}
