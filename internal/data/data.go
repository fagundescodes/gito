package data

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

const GITDir = ".gito"

func Init() (string, error) {
	if err := os.Mkdir(GITDir, 0o755); err != nil {
		return "", err
	}
	if err := os.Mkdir(GITDir+"/objects", 0o755); err != nil {
		return "", err
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return wd, nil
}

func HashObject(content []byte, objectType string) (string, error) {
	obj := append([]byte(objectType), 0)
	obj = append(obj, content...)

	sum := sha1.Sum(obj)
	oid := hex.EncodeToString(sum[:])

	if err := os.WriteFile(GITDir+"/objects/"+oid, obj, 0o644); err != nil {
		return "", err
	}

	return oid, nil
}

func GetObject(oid string) ([]byte, error) {
	obj, err := os.ReadFile(GITDir + "/objects/" + oid)
	if err != nil {
		return nil, err
	}

	_, after, ok := bytes.Cut(obj, []byte{0})
	if !ok {
		return nil, fmt.Errorf("invalid object format")
	}

	return after, nil
}
