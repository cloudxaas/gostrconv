package cxstrconv

func Int8toa(n int8) []byte {
    if n == 0 {
        return []byte{'0'}
    }

    buf := make([]byte, 0, 3)
    neg := false

    if n < 0 {
        neg = true
        n = -n
    }

    for n > 0 {
        buf = append(buf, byte(n%10)+'0')
        n /= 10
    }

    if neg {
        buf = append(buf, '-')
    }

    for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
        buf[i], buf[j] = buf[j], buf[i]
    }

    return buf
}

func Int16toa(n int16) []byte {
    if n == 0 {
        return []byte{'0'}
    }

    buf := make([]byte, 0, 5)
    neg := false

    if n < 0 {
        neg = true
        n = -n
    }

    for n > 0 {
        buf = append(buf, byte(n%10)+'0')
        n /= 10
    }

    if neg {
        buf = append(buf, '-')
    }

    for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
        buf[i], buf[j] = buf[j], buf[i]
    }

    return buf
}

func Int32toa(n int32) []byte {
    if n == 0 {
        return []byte{'0'}
    }

    buf := make([]byte, 0, 11)
    neg := false

    if n < 0 {
        neg = true
        n = -n
    }

    for n > 0 {
        buf = append(buf, byte(n%10)+'0')
        n /= 10
    }

    if neg {
        buf = append(buf, '-')
    }

    for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
        buf[i], buf[j] = buf[j], buf[i]
    }

    return buf
}

func Int64toa(n int64) []byte {
    if n == 0 {
        return []byte{'0'}
    }

    buf := make([]byte, 0, 20)
    neg := false

    if n < 0 {
        neg = true
        n = -n
    }

    for n > 0 {
        buf = append(buf, byte(n%10)+'0')
        n /= 10
    }

    if neg {
        buf = append(buf, '-')
    }

    for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
        buf[i], buf[j] = buf[j], buf[i]
    }

    return buf
}


func Uint8toa(n uint8) string {
    if n == 0 {
        return "0"
    }
    buf := [3]byte{}
    i := len(buf)
    for n > 0 {
        i--
        buf[i] = '0' + n%10
        n /= 10
    }
    return string(buf[i:])
}

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

func Uint32toa(n uint32) string {
    if n == 0 {
        return "0"
    }
    buf := [11]byte{} // max uint32 string length is 10 digits
    i := len(buf) - 1
    for n > 0 {
        buf[i] = byte(n%10) + '0'
        n /= 10
        i--
    }
    return string(buf[i+1:])
}

func Uint64toa(n uint64) string {
    if n == 0 {
        return "0"
    }
    buf := [20]byte{}
    i := len(buf) - 1
    for n > 0 {
        q := n / 10
        r := n % 10
        buf[i] = byte(r + '0')
        i--
        n = q
    }
    return string(buf[i+1:])
}
