package cryptodemo

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"software.sslmate.com/src/go-pkcs12"
)

func FingerprintSHA256(filePath string) (string, error) {
	// 读取文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))

	// 解析 PEM
	block, _ := pem.Decode(data)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block")
	}

	// 对 DER 编码做 SHA-256
	hash := sha256.Sum256(block.Bytes)
	hexStr := hex.EncodeToString(hash[:])

	fmt.Println(hexStr)

	// 转冒号分隔大写
	var parts []string
	for i := 0; i < len(hexStr); i += 2 {
		parts = append(parts, strings.ToUpper(hexStr[i:i+2]))
	}

	return strings.Join(parts, ":"), nil
}

// FingerprintP12SHA256 计算 P12/PFX 文件的证书 SHA-256 Fingerprint
func FingerprintP12SHA256(filePath, password string) (string, error) {
	// 读取文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Decode 返回 (privateKey, cert, caCerts, err)
	_, cert, _, err := pkcs12.DecodeChain(data, password)
	if err != nil {
		return "", err
	}

	if cert == nil {
		return "", fmt.Errorf("no certificate found in P12 file")
	}

	// 对证书 DER 编码计算 SHA-256
	hash := sha256.Sum256(cert.Raw)
	hexStr := hex.EncodeToString(hash[:])
	fmt.Println(hexStr)

	// 转成冒号分隔大写
	var parts []string
	for i := 0; i < len(hexStr); i += 2 {
		parts = append(parts, strings.ToUpper(hexStr[i:i+2]))
	}

	return strings.Join(parts, ":"), nil
}
