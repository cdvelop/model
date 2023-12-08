package model

type CipherAdapter interface {
	Encrypt(content []byte) (out, err string)
	Decrypt(content string) (out, err string)
}
