package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("aaaa"))
	conn.Write(intToBytes(1))
	conn.Write(intToBytes(8))
	conn.Write(intToBytes(int(time.Now().Unix())))
	conn.Write(intToBytes(123))
	conn.Write([]byte("asdfqwer"))
	conn.Write([]byte("bbbb"))

	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf[:len])
}

func intToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &tmp)
	return bytesBuffer.Bytes()
}
