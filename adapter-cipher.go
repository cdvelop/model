package model

type CipherAdapter interface {
	Encrypt(content string) (out, err string)
	Decrypt(content string) (out, err string)
}
