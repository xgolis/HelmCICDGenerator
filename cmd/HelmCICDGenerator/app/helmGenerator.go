package app

import (
	"fmt"
	"io/ioutil"
	"strings"

	cp "github.com/otiai10/copy"
)

type Config struct {
	HelmPath string
}

func createHelmChart(app UserApp) error {
	filePrefix := "./helm"
	err := changeMainChart(filePrefix, app)
	if err != nil {
		return err
	}
	err = changeValuesChart(filePrefix, app)
	if err != nil {
		return err
	}

	cp.Copy(filePrefix, "./"+app.Name+"/helm")

	return nil
}

func changeMainChart(prefix string, app UserApp) error {
	input, err := ioutil.ReadFile(prefix + "/Chart.yaml")
	if err != nil {
		return fmt.Errorf("error while reading chat.yaml")
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "name: ") {
			lines[i] = "name: " + app.Name
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(prefix+"/Chart.yaml", []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("error rewriting reading chat.yaml")
	}

	return nil
}

func changeValuesChart(prefix string, app UserApp) error {
	input, err := ioutil.ReadFile(prefix + "/values.yaml")
	if err != nil {
		return fmt.Errorf("error while reading values.yaml")
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "fullImage:") {
			lines[i] = "  fullImage: " + app.Image
			// fmt.Printf(lines[i])
		}
		if strings.Contains(line, " nodePort:") {
			lines[i] = "  nodePort: " + app.Port
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(prefix+"/values.yaml", []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("error while rewriting values.yaml")
	}

	return nil
}
