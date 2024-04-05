package share

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func Go(cookie, token, urlShare string, total, x *int){
  url := fmt.Sprintf("https://graph.facebook.com/v13.0/me/feed?link=%s&published=0&access_token=%s", urlShare, token)
  client := &http.Client{}
  req, err := http.NewRequest("POST", url, nil)
  if err != nil {
    fmt.Println(err.Error())
  }
  var headerMap map[string]string = map[string]string{
    "User-Agent": "Mozilla/5.0 (Linux; Android 8.1.0; S45B Build/OPM2.171019.012;) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/96.0.4664.45 Mobile Safari/537.36",
    "Content-Type": "application/json",
    "cookie": cookie,
  }
  for key, value := range headerMap {
    req.Header.Add(key, value)
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err.Error())
  }
  app, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    fmt.Println(err.Error())
  }
  JSON := make(map[string]interface{})
  json.Unmarshal([]byte(app), &JSON)
  if JSON["id"] != nil {
    fmt.Println(fmt.Sprintf("\r[✓] Berhasil: %s => %d", JSON["id"].(string), *total + 1))
  } else {
    fmt.Println(fmt.Sprintf("\r[×] Gagal ngeshare => %d", *total + 1))
  }
  *total++
  *x++
  fmt.Print("\r[GoShare] Tekan CTRL + Z untuk berhenti")
}
