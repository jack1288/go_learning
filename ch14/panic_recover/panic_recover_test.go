package  panic_recover_test


import (

   "testing"
   "fmt"
   "errors"
)




func TestPanicVxExit(t *testing.T) {
     //defer func() {
     //   fmt.Println("Finally!")
     //}()

      defer func() {
         if err := recover(); err != nil {
            fmt.Println("recovered from", err)
         }
     }()
     fmt.Println("start")
     panic(errors.New("Something wrong!"))

}
