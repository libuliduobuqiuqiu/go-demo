package goothers

import (
	"fmt"
	"os"

	"github.com/tjfoc/gmsm/pkcs12"
	"github.com/tjfoc/gmsm/x509"
)

func ParseSM2P12(fileName, password string) error {

	// 使用 gm 的 pkcs12 解码器（而不是标准库）
	cert, key, err := pkcs12.SM2P12Decrypt(fileName, password)
	if err != nil {
		err = fmt.Errorf("解析失败: %v", err)
		return err
	}

	fmt.Println(key)
	fmt.Println(cert.Subject, cert.Issuer, cert.PublicKeyAlgorithm)
	fmt.Println(cert.NotAfter, cert.NotBefore)
	return nil
}

func ParseSM2P12Bydata(fileName, password string) error {

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	_, cert, err := pkcs12.DecodeAll(fileBytes, password)
	if err != nil {
		return err
	}

	fmt.Println(cert[0].Subject, cert[0].SignatureAlgorithm, cert[0].PublicKeyAlgorithm)
	fmt.Println(x509.SM2)
	fmt.Println(cert[0].PublicKeyAlgorithm == x509.SM2)
	return nil
}
