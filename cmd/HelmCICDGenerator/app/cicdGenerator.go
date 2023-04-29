package app

import (
	"fmt"
	"io/ioutil"
	"strings"

	cp "github.com/otiai10/copy"
)

func createCICDPipelines(app UserApp) error {
	prefix := "./pipelines"

	cp.Copy(prefix+"/.gitlab-ci.yml", "./"+app.Name)
	cp.Copy(prefix+"/.github", "./"+app.Name)

	err := createGitlabPipeline(prefix, app)
	if err != nil {
		return fmt.Errorf("error while creating gitlab pipeline: %v", err)
	}

	err = createGithubPipeline(prefix, app)
	if err != nil {
		return fmt.Errorf("error while creating github pipeline: %v", err)
	}

	return nil
}

func createGitlabPipeline(prefix string, app UserApp) error {
	input, err := ioutil.ReadFile(prefix + "/.gitlab-ci.yml")
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "  IMAGE_NAME: ") {
			lines[i] = "  IMAGE_NAME: " + app.Image
		}
		if strings.Contains(line, "  NAMESPACE: ") {
			lines[i] = "  NAMESPACE: " + app.UserName
		}
		if strings.Contains(line, "APPNAME: ") {
			lines[i] = "  APPNAME: " + app.Name
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./"+app.Name+"/.gitlab-ci.yml", []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

func createGithubPipeline(prefix string, app UserApp) error {
	input, err := ioutil.ReadFile(prefix + "/.github/workflows/deploy.yaml")
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		// if strings.Contains(line, "  IMAGE_NAME: ") {
		// 	lines[i] = "  IMAGE_NAME: " + app.Image
		// }
		if strings.Contains(line, "  NAMESPACE: ") {
			lines[i] = "  NAMESPACE: " + app.UserName
		}
		if strings.Contains(line, "APPNAME: ") {
			lines[i] = "  APPNAME: " + app.Name
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./"+app.Name+"/.github/workflows/deploy.yaml", []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}
