package cxstrconv
  
import (
    "fmt"
    "strconv"
)
  

func Uint16toa(n uint16) string {
    if n == 0 {
        return "0"
    }
    
    var b [5]byte
    i := len(b) - 1
    for n > 0 {
        b[i] = byte(n%10) + '0'
        n /= 10
        i--
    }
    
    return string(b[i+1:])
}

