package terminal

import (
  "fmt"
  "os"
  "os/exec"
  "runtime"
)

type ErrHndl struct {
  message string
}

func init(){
  err := Clear()
  if err != nil {
    fmt.Println(err.Error())
  }
}

func (errhndl *ErrHndl) Error() string {
  return errhndl.message
}

func Clear() error {
  dev := runtime.GOOS
  var command string
  if dev == "linux" || dev == "android" {
    command = "clear"
  } else if dev == "windows" {
    command = "cls"
  } else {
    return &ErrHndl{"Error: os tidak diketahui"}
  }
  execute := exec.Command(command)
  execute.Stdout = os.Stdout
  execute.Run()
  return nil
}
