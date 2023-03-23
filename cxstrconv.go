package cxstrconv
  
import (
    "fmt"
    "strconv"
)
  
func CXstrconv3BytesToUint16(in string) uint16 {
    return uint16(strconv.FormatUint16(in,10)
}
                  
func CXstrconvUint16To3Bytes(in uint16) string {
  return strconv.ParseUint16(in,36)
}
