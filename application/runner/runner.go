package runner

import (
  "io"
  "os"
  "fmt"
  "log"
  "time"
  "strings"
  "os/exec"
  "io/ioutil"
  "path/filepath"
)

const (
  version = 1.0
  logExtension = ".log"
)

var (
  // ciDIR is the absolute path to the CI directory
  ciDIR = filepath.Join(os.Getenv("ROOT_DIR"), "ci")
  // LogDIR is the absolute path to the CI log directory
  logDir = filepath.Join(ciDIR, "logs")
)

type runnerDetails struct {

  // Define the logs information for this runner.
  logName string
  logSize string

  // Define the runner information for this runner.
  runnerName string
  runnerLanguage string
  runnerEndingDate time.Time
  runnerStartingDate time.Time

  // Define the project main informations.
  projectRepositoryUrl string
  projectRepositoryName string
  projectRepositoryBranch string
  projectRepositoryLanguage string
}

func init() {

}

func startBuild() {

  fmt.Println("Build started at", time.Now())
}

func pauseBuild() {

}

func cancelBuild() {

}

func getBuildStatus() {

}

func updateBuildStatus() {

}

func prepareEnvVariables() {

}

/**
 * Allow to separate the runner from the main runner process.
 */
func separateRunnerFromMain(separated bool, filename string) error {
  if separated {
    dir, file := filepath.Split(filename)
    log.Printf("Making dir: %s for file: %s\n", dir, file)
    return os.MkdirAll(dir, 0777)
  }

  dir, file := filepath.Split(filename)
  log.Printf("Making dir: %s for file: %s\n", dir, file)
  return os.MkdirAll(dir, 0755)
}
