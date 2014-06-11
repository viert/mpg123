mpg123
======

Simple mpeg decoder for Go

Example:

```go
package main

import (
  "fmt"
  "mpg123"
  "os"
)

func main() {

  f, err := os.Open("/home/viert/Downloads/56.mp3")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  of, err := os.Create("output.wav")
  if err != nil {
    panic(err)
  }
  defer of.Close()
  var inc, outc int

  decoder, err := mpg123.NewWriter(of)
  if err != nil {
    panic(err)
  }

  buf := make([]byte, 16384)

  for {
    n, err := f.Read(buf)
    if err != nil {
      break
    }
    inc += n

    m, err := decoder.Write(buf[0:n])
    if err != nil {
      break
    }
    outc += m
  }
  fmt.Println(inc, "bytes in,", outc, "bytes out")
}
```
