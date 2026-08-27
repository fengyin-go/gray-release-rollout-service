package service

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash/crc32"
)

// ChecksumType 校验和算法类型。
type ChecksumType string

// 支持的校验和算法。
const (
	ChecksumSHA256 ChecksumType = "sha256"
	ChecksumSHA1   ChecksumType = "sha1"
	ChecksumMD5    ChecksumType = "md5"
	ChecksumCRC32  ChecksumType = "crc32"
)

// ComputeChecksum 按算法计算字符串的十六进制校验和。
func ComputeChecksum(algorithm ChecksumType, data string) string {
	switch algorithm {
	case ChecksumSHA1:
		h := sha1.Sum([]byte(data))
		return hex.EncodeToString(h[:])
	case ChecksumMD5:
		h := md5.Sum([]byte(data))
		return hex.EncodeToString(h[:])
	case ChecksumCRC32:
		return crc32Hex(data)
	default:
		h := sha256.Sum256([]byte(data))
		return hex.EncodeToString(h[:])
	}
}

func crc32Hex(data string) string {
	v := crc32.ChecksumIEEE([]byte(data))
	b := []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
	return hex.EncodeToString(b)
}

// VerifyChecksum 校验数据是否匹配给定校验和。
func VerifyChecksum(algorithm ChecksumType, data, checksum string) bool {
	if checksum == "" {
		return false
	}
	return ComputeChecksum(algorithm, data) == checksum
}

// VerifyVersionChecksum 校验版本制品校验和（未配置校验和则视为通过）。
func (s *Service) VerifyVersionChecksum(versionID, data string) (bool, error) {
	v, err := s.store.GetVersion(versionID)
	if err != nil {
		return false, err
	}
	if v.Checksum == "" {
		return true, nil
	}
	return VerifyChecksum(ChecksumSHA256, data, v.Checksum), nil
}
