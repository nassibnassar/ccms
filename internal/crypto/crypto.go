package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
)

func HashPassword(key string, salt, secretKey []byte) string {
	concat := append(secretKey, salt...)
	return base64.RawStdEncoding.EncodeToString(argon2.IDKey([]byte(key), concat, 1, 64*1024, 4, 32))
}

func RandomKey() []byte {
	key := make([]byte, 16)
	_, _ = rand.Read(key[:])
	return key
}

func EncodeToHexString(b []byte) string {
	return hex.EncodeToString(b)
}

func DecodeFromHexString(s string) ([]byte, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func InputPassword(prompt string, confirm bool) (string, error) {
	// avoid leaving terminal in no-echo state
	var sigc = make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	go func() {
		<-sigc
	}()
	defer signal.Reset(os.Interrupt)

	fmt.Print(prompt)
	p, err := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	if err != nil {
		return "", err
	}

	if confirm {
		fmt.Print("(Confirming) " + prompt)
		q, err := term.ReadPassword(syscall.Stdin)
		fmt.Println()
		if err != nil {
			return "", err
		}
		if string(p) != string(q) {
			return "", errors.New("passwords do not match")
		}
	}

	return string(p), nil
}
