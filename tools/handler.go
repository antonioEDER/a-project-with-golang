package tools

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func FindByFilterHandler(c echo.Context) (err error) {
	plaintext := []byte("Eder Antonio")

	ciphertext, err := Encrypt(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(ciphertext))

	// ciphertext, _ := hex.DecodeString("ac986000761ea009fffa133eeef6684df6d108393f8326ed0b3434b7a5b7e852")

	// result, err := Decrypt(ciphertext)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", result)

	return c.JSON(200, "")
}
