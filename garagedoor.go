package main

import (
        "fmt"
	"os"
        "flag"
        "time"
	"github.com/stianeikeland/go-rpio"
)

func main() {
  if err := rpio.Open(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer rpio.Close()

  pinPtr := flag.Int("pin", 0, "GPIO pin number.  Ex: --pin=17")
  flag.Parse()

  pin := rpio.Pin(*pinPtr)
  if pin == 0 {
    os.Exit(1)

  } else {
    pin.Output()
    time.Sleep(time.Millisecond * 250)
    pin.Input()
  }
}
