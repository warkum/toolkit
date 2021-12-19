package encryptx

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		log.Println(err)
	}

	key := hex.EncodeToString(bytes)

	type args struct {
		stringToEncrypt string
		keyString       string
	}
	tests := []struct {
		name              string
		args              args
		wantDecryptString string
		wantErr           bool
	}{
		{
			"Test success encrypted descrypted",
			args{
				stringToEncrypt: "hello world",
				keyString:       key,
			},
			"hello world",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncryptedString, err := Encrypt(tt.args.stringToEncrypt, tt.args.keyString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotDescryptString, err := Decrypt(gotEncryptedString, tt.args.keyString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDescryptString != tt.wantDecryptString {
				t.Errorf("Encrypt() = %v, want %v", gotEncryptedString, tt.wantDecryptString)
			}
		})
	}
}
