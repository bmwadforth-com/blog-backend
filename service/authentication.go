package service

import (
	"blog-backend/util"
	"context"
	armorUtil "github.com/bmwadforth-com/armor-go/src/util"
	"github.com/bmwadforth-com/armor-go/src/util/jwt"
	"github.com/bmwadforth-com/armor-go/src/util/jwt/common"
	"google.golang.org/api/idtoken"
	"time"
)

func NewBearerToken(username string) string {
	key := []byte(util.Config.JwtSigningKey)

	tokenBuilder := jwt.NewJWSToken(common.HS256, key)

	claims := common.NewClaimSet()
	claims.Add(string(common.Audience), "blog-backend")
	claims.Add(string(common.Subject), username)
	claims.Add(string(common.IssuedAt), time.Now())

	token, err := tokenBuilder.AddClaims(claims).Serialize()
	if err != nil {
		panic(err)
	}

	return token
}

func ValidateBearerToken(tokenString string) bool {
	key := []byte(util.Config.JwtSigningKey)

	tokenBuilder, err := jwt.DecodeToken(tokenString, key)
	if err != nil {
		armorUtil.LogError("token decode failed: %v", err)
		return false
	}

	_, err = tokenBuilder.Validate()
	if err != nil {
		armorUtil.LogError("token validation failed: %v", err)
		return false
	}

	return true
}

func GetTokenClaims(tokenString string) map[string]interface{} {
	key := []byte(util.Config.JwtSigningKey)

	tokenBuilder, err := jwt.DecodeToken(tokenString, key)
	if err != nil {
		armorUtil.LogError("token decode failed: %v", err)
		return nil
	}

	return tokenBuilder.GetClaims()
}

func FetchIdentityToken(ctx context.Context, audience string) (string, error) {
	ts, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		return "", err
	}
	token, err := ts.Token()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
