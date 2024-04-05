package auth

import (
  "regexp"
  "net/http"
  "os"
  "io/ioutil"
)

type ErHndl struct {
  message string
}

func (erhndl *ErHndl) Error() string {
  return erhndl.message
}

func GetToken(cookie *string, condition int8) (*string, error) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://business.facebook.com/business_locations", nil)
  headers := map[string]string{
    "User-Agent":"Mozilla/5.0 (Linux; Android 8.1.0; MI 8 Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36",
    "Host":"business.facebook.com",
    "Content-Type":"text/html; charset=utf-8",
    "Cookie": *cookie,
  }
  for khead, vhead := range headers {
    req.Header.Add(khead, vhead)
  }
  if err != nil {
    return cookie, err
  }
  res, err2 := client.Do(req)
  defer res.Body.Close()
  if err2 != nil {
    return cookie, err2
  }
  strout, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return cookie, err
  }
  reSynt := regexp.MustCompile("(EAAG\\w+)")
  if !reSynt.MatchString(string(strout)) {
    if condition == 1 {
      os.RemoveAll("./data")
      os.Mkdir("./data", os.ModePerm)
    }
    return cookie, &ErHndl{"invalid"}
  }
  var token string = reSynt.FindStringSubmatch(string(strout))[1]
  return &token, nil
}
