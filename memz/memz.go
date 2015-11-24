package main

import "os"
import "io/ioutil"
import "fmt"
import "github.com/abeconnelly/memz"

var g_normalize_flag bool

func main() {

  if len(os.Args)<3 {
    fmt.Fprintf(os.Stderr, "provide filenames\n")
    os.Exit(1)
  }

  fn0 := os.Args[1]
  fn1 := os.Args[2]

  s0,e := ioutil.ReadFile(fn0)
  if e!=nil { panic(e) }
  s1,e := ioutil.ReadFile(fn1)
  if e!=nil { panic(e) }

  X,Y,sc := memz.Hirschberg(s0,s1)

  if g_normalize_flag {
    memz.SeqPairNormalize(X,Y)
  }

  fmt.Printf("%d\n%s\n%s\n", sc, X, Y)
}
