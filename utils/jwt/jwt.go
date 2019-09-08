package jwt

import (
	"github.com/HRDVV/cover-note/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type Jwt struct {
	claims      jwt.Claims
	expiresTime int64
}

func (j *Jwt) GenerateToken(user model.User) (string, error) {
	j.expiresTime = time.Now().Unix() + viper.GetInt64("jwt.expTime")
	j.claims = jwt.StandardClaims{
		Audience:  user.Username, // 受众
		ExpiresAt: j.expiresTime, // 失效时间
		//Id:        user.Username,       // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "covernote",       // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}
	jwtSecret := []byte(viper.GetString("jwt.secret"))
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func (j *Jwt) ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
