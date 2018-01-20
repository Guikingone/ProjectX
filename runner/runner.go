package runner

import (
  "os"
  "fmt"
  "logs"
	"strings"
  "os/exec"
	"path/filepath"
)

const (
  version = 1.0
  logExtension = ".log"
)

type runnerDetails struct {

  // Define the logs information for this runner.
  logName string
  logSize string

  // Define the runner information for this runner.
  runnerName string
  runnerLanguage string

  // Define the project main informations.
  projectRepositoryName string
  projectRepositoryBranch string
  projectRepositoryUrl string
  projectRepositoryLanguage string
}

func init() {

}

func startBuild() {

}

func pauseBuild() {

}

func cancelBuild() {

}

func getBuildStatus() {

}

func updateBuildStatus() {

}
