package main

import (
  "fmt"
  "goshare/lib"
  "goshare/ceklogin"
  "goshare/auth"
  "goshare/terminal"
  "os"
  "slices"
  "goshare/share"
  "time"
  "goshare/intercode"
  "strconv"
)

type GoShare interface {
  uInterface(lib.Rule, intercode.GetZ)
  loginView(lib.Rule, intercode.GetZ)
  createData(string, string)
  verifikasiLogin(lib.Rule, intercode.GetZ)
  clearTerm()
  logout(lib.Rule, intercode.GetZ)
  shareSubmit(lib.Rule)
}

func (sharebot ShareBot) uInterface(b lib.Rule, i intercode.GetZ){
  sharebot.clearTerm()
  fmt.Println(b.Banner())
  fmt.Println(b.Line())
  i.SpttyChan()
  fmt.Println("[1] Mulai bot share facebook")
  fmt.Println("[0] Logout")
  var input string
  for {
    fmt.Println(b.Line())
    fmt.Print("[?] Pilih: ")
    fmt.Scanln(&input)
    if input == "" {
      fmt.Println(b.Line())
      fmt.Println("[!] Jangan kosong")
    } else {
      if slices.Contains([]string{"0","00","1","01"}, input) {
        break
      } else {
        input = ""
        fmt.Println(b.Line())
        fmt.Println("[!] Pilihan tidak tersedia")
      }
    }
  }
  if input == "1" || input == "01" {
    sharebot.shareSubmit(b)
  } else {
    sharebot.logout(b, i)
  }
}

func (sharebot ShareBot) loginView(b lib.Rule, i intercode.GetZ) {
  sharebot.clearTerm()
  fmt.Println(b.Banner())
  fmt.Println(b.Line())
  i.SpttyChan()
  fmt.Println("[!] Anda belum login")
  fmt.Println("[!] Login untuk melanjutkan")
  fmt.Println("[!] Pastikan facebook tidak dalam mode gratis")
  var cookie string
  for {
    fmt.Println(b.Line())
    fmt.Print("[?] Cookies facebook: ")
    fmt.Scanln(&cookie)
    if cookie == "" {
      fmt.Println(b.Line())
      fmt.Println("[!] Jangan kosong")
    } else {
      token, err := auth.GetToken(&cookie, 0)
      if err != nil {
        fmt.Println(b.Line())
        fmt.Println("[!] Login gagal, periksa cookies anda")
      } else {
        fmt.Println("[✓] Login berhasil ...")
        err := ceklogin.WriteLog(&cookie, token)
        if err != nil {
          fmt.Println("[!] Error: ", err.Error())
          break
        }
        break
      }
    }
  }
}

func (sharebot *ShareBot) shareSubmit(b lib.Rule) {
  var shareUrl string
  for {
    fmt.Println(b.Line())
    fmt.Print("[?] Masukkan link postingan publik: ")
    fmt.Scanln(&shareUrl)
    if shareUrl == "" {
      fmt.Println(b.Line())
      fmt.Println("[!] Jangan kosong")
    } else {
      fmt.Println(b.Line())
      break
    }
  }
  total := 0
  for {
    stop := 0
    for i := 0; i < 30; i++ {
      go share.Go(sharebot.Cookie, sharebot.Token, shareUrl, &total, &stop)
    }
    for {
      if stop == 30 {
        break
      }
    }
    dot := "[/]"
    for second := 0; second < 15; second++ {
      strSecond := strconv.Itoa(15 - second)
      if len(strSecond) == 1 {
        strSecond = "0" + strSecond
      }
      fmt.Print(fmt.Sprintf("\r[GoShare] Delay %s detik             %s", strSecond, dot))
      if dot == "[/]" {
        dot = "[\\]"
      } else {
        dot = "[/]"
      }
      time.Sleep(1*time.Second)
    }
  }
}

func (sharebot *ShareBot) logout(b lib.Rule, i intercode.GetZ){
  var konfirmasi string
  for {
    fmt.Println(b.Line())
    fmt.Print("[?] Yakin ingin logout (y/t): ")
    fmt.Scanln(&konfirmasi)
    if konfirmasi == "" {
      fmt.Println(b.Line())
      fmt.Println("[!] Jangan kosong")
    } else {
      if slices.Contains([]string{"y","Y","t","T"}, konfirmasi) {
        break
      } else {
        konfirmasi = ""
        fmt.Println(b.Line())
        fmt.Println("[!] Pilihan tidak tersedia")
      }
    }
  }
  if konfirmasi == "y" || konfirmasi == "Y" {
    os.RemoveAll("./data")
    os.Mkdir("./data", os.ModePerm)
    fmt.Println("[✓] Logout berhasil")
  } else {
    sharebot.uInterface(b, i)
  }
}

func newFanda(f *string, b lib.Rule) intercode.GetZ {
  newfan := new(Fanda)
  newfan.author = fmt.Sprintf("%s\n%s", *f, b.Line())
  return newfan
}

func (sharebot *ShareBot) clearTerm(){
  err := terminal.Clear()
  if err != nil {
    fmt.Println(err.Error())
  }
}

func (fanda *Fanda) SpttyChan(){
  fmt.Println(fanda.author)
}

type ShareBot struct {
  Cookie string
  Token string
}

type Fanda struct {
  author string
}

func (sharebot *ShareBot) createData(cookie, token string){
  if cookie != "" {
    sharebot.Cookie = cookie
  }
  if token != "" {
    sharebot.Token = token
  }
}

func (sharebot *ShareBot) verifikasiLogin(b lib.Rule, i intercode.GetZ){
  fmt.Println(b.Banner())
  fmt.Println(b.Line())
  i.SpttyChan()
  fmt.Println("[.] Memverifikasi login ...")
}

func main(){
  var run GoShare = new(ShareBot)
  var banner *lib.VStr = lib.NewBan()
  var author string = "[=] Author  : Fanda\n[=] Facebook: http://fb.com/profile.php?id=100024425583446"
  impthor := newFanda(&author, banner)
  cookie, err := ceklogin.CekFile()
  run.createData(cookie, "")
  if err != nil {
    run.loginView(banner, impthor)
  } else {
    run.verifikasiLogin(banner, impthor)
    verif, err := auth.GetToken(&cookie, 1)
    if err != nil {
      os.RemoveAll("./data")
      os.Mkdir("./data", os.ModePerm)
      fmt.Println("[×] Cookie invalid")
    } else {
      run.createData("", *verif)
      run.uInterface(banner, impthor)
    }
  }
}
