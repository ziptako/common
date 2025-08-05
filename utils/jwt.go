package utils

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Auth struct {
	AccessSecret string
	AccessExpire int64
}
type Payload struct {
	Id int64 `json:"id"`
}

func GenerateJwtToken(a Auth, p *Payload) (string, error) {
	// 创建一个 MapClaims 类型的声明
	claims := make(jwt.MapClaims)
	// 计算过期时间
	claims["exp"] = time.Now().Unix() + a.AccessExpire // 设置 JWT 的过期时间（exp），通常需要一个 UNIX 时间戳
	claims["iat"] = time.Now().Unix()                  // 设置签发时间（iat）
	claims["payload"] = p                              // 自定义的负载（payload），可以设置为任何信息，例如用户名、用户ID等
	// 创建新的 JWT
	token := jwt.New(jwt.SigningMethodHS256) // 使用 HMAC SHA256 签名方法创建新的 JWT
	// 将声明分配给 JWT
	token.Claims = claims
	// 使用 secretKey 签名JWT，并返回生成的字符串和错误（如果有）
	return token.SignedString([]byte(a.AccessSecret))
}

// GetPayloadFromContext 从context中获取用户payload信息
// 返回payload和错误信息
func GetPayloadFromContext(ctx context.Context) (*Payload, error) {
	unverifiedPayload := ctx.Value("payload")
	if unverifiedPayload == nil {
		return nil, errors.New("payload not found in context")
	}

	payload, ok := unverifiedPayload.(Payload)
	if !ok {
		return nil, errors.New("invalid payload type")
	}

	return &payload, nil
}
