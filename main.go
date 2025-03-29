package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodeWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str)) // Sum returns the MD5 checksum of the data.
	fmt.Println("\nhashCode (Byte Array) is : ", hashCode)
	return hex.EncodeToString(hashCode[:]) // Encode the byte array of Hash to a string
}

func main() {
	var password string
	fmt.Println("Enter Password below ")
	fmt.Scanln(&password)
	fmt.Println("\nPassword encrypted to: ", encodeWithMD5(password))
}
