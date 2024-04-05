package ceklogin

import (
  "os/exec"
  "strings"
  "slices"
  "os"
  "io/ioutil"
)

type FileError struct {
  message string
}

func (fileerror *FileError) Error() string {
  return fileerror.message
}

func WriteLog(cookie *string, token *string) error {
  var files map[string]string = map[string]string{
    "./data/cookies.txt": *cookie,
    "./data/auth": "true",
  }
  for Name, Value := range files {
    open, err := os.OpenFile(Name, os.O_CREATE|os.O_WRONLY, 0666)
    defer open.Close()
    if err != nil {
      return err
    }
    open.WriteString(Value)
  }
  return nil
}

func CekFile() (string, error) {
  commandX, _ := exec.Command("ls", "./data").Output()
  output := strings.Split(string(commandX), "\n")
  if !slices.Contains(output, "auth") {
    return "", &FileError{"file \"auth\" tidak ada"}
  }
  file, err := os.OpenFile("./data/cookies.txt", os.O_RDONLY, 0666)
  if err != nil {
    return "", err
  }
  strfile, err := ioutil.ReadAll(file)
  if err != nil {
    return "", err
  }
  return string(strfile), nil
}
