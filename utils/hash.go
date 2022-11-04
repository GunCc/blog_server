package utils

import "golang.org/x/crypto/bcrypt"

// BcryptHash 使用 bcrypt 对密码进行加密
func BcrypHash(password string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
