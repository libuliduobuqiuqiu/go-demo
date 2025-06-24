package goothers

import (
	"fmt"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

func ParseRegular(fileName, password string) error {

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	// fmt.Println(string(fileBytes))
	// 使用标准 Decode 获取私钥和证书
	_, certificate, err := pkcs12.Decode(fileBytes, password)
	if err != nil {
		return fmt.Errorf("解析 p12 失败: %w", err)
	}

	// 打印基本信息
	fmt.Println("证书信息：")
	fmt.Printf("Subject: %s\n", certificate.Subject)
	fmt.Printf("Issuer : %s\n", certificate.Issuer)
	fmt.Printf("Alg    : %s\n", certificate.PublicKeyAlgorithm)

	return nil
}
