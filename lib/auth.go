package lib



func NewBan() *VStr {
  vstr := new(VStr)
  vstr.banner = "╔═╗┌─┐╔═╗┬ ┬┌─┐┬─┐┌─┐\n║ ╦│ │╚═╗├─┤├─┤├┬┘├┤\n╚═╝└─┘╚═╝┴ ┴┴ ┴┴└─└─┘"
  vstr.line = "============================"
  return vstr
}

type Rule interface {
  Banner() string
  Line() string
}

type VStr struct {
  banner string
  line string
}

func (vstr *VStr) Banner() string {
  return vstr.banner
}

func (vstr *VStr) Line() string {
  return vstr.line
}
