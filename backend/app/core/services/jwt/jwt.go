package jwt

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"time"

	service "github.com/lipidbilayer/boiler/app/core/services"
)

const (
	ISSUER_KEY             = "iss"
	ISSUED_AT_KEY          = "iat"
	EXPIRATION_KEY         = "exp"
	SUBJECT_KEY            = "sub"
	EXPIRE_OFFSET          = 3600
	TOKEN_CLAIMS_KEY       = "jwt.auth.claims"
	AUTHENTICATED_USER_KEY = "jwt.auth.user"
)

// Objects implementing the AuthHandler interface can be
// registered to Authenticate User for application
type AuthHandler interface {
	Authenticate(username, password string) (string, bool)
}

// The AuthHandlerFunc type is an adapter to allow the use of
// ordinary functions as Auth handlers.
type AuthHandlerFunc func(string, string) (string, bool)

// Authenticate calls f(u, p).
func (f AuthHandlerFunc) Authenticate(u, p string) (string, bool) {
	return f(u, p)
}

var (
	Realm             string
	issuer            string
	privateKey        *rsa.PrivateKey
	publicKey         *rsa.PublicKey
	expiration        time.Duration // in minutues
	refreshExpiration time.Duration //in minutes
	isIssuerExists    bool
	// handler        AuthHandler
)

/*
Method Init initializes JWT auth provider based on given config values from app.conf
*/
func Init(config service.ConfigService, file service.EmbeddedFile) {
	Realm = config.GetJWTRealmName()
	issuer = config.GetJWTIssuerName()
	expiration = time.Duration(config.GetJWTExpiration())
	refreshExpiration = time.Duration(config.GetJWTRefreshExpiration())

	// if _, ok := authHandler.(AuthHandler); !ok {
	// 	revel.AppLog.Fatal("Auth Handler doesn't implement interface jwt.AuthenticationHandler")
	// }

	Realm = fmt.Sprintf(`Bearer realm="%s"`, Realm)

	isIssuerExists = len(issuer) > 0
	// handler = authHandler.(AuthHandler)
	// statikFS, err := fs.New()
	// if err != nil {
	// 	log.Fatal("File storage not found")
	// }
	privateKey = loadPrivateKey(config.GetJWTPrivateKeyPath(), file.GetEmbeddedFile())
	publicKey = loadPublicKey(config.GetJWTPublicKeyPath(), file.GetEmbeddedFile())
}

// func Authenticate(username, password string) (string, bool) {
// 	return handler.Authenticate(username, password)
// }

// Method IsInBlocklist is checks against logged out tokens
func IsInBlocklist(token string) bool {
	var existingToken string
	// cache.Get(token, &existingToken)

	if len(existingToken) > 0 {
		log.Printf("Yes, blocklisted token [%v]", existingToken)
		return true
	}

	return false
}

// Private Methods
func loadPrivateKey(keyPath string, file http.FileSystem) *rsa.PrivateKey {
	keyData := readKeyFile(keyPath, file)

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(keyData.Bytes)
	if err != nil {
		log.Fatalf("Private key import error [%v]", keyPath)
	}

	return privateKeyImported
}

func loadPublicKey(keyPath string, file http.FileSystem) *rsa.PublicKey {
	keyData := readKeyFile(keyPath, file)

	rsaPublicKey, err := x509.ParsePKCS1PublicKey(keyData.Bytes)
	if err != nil {
		log.Fatalf("Public key import error [%v]", keyPath)
	}

	return rsaPublicKey
}

func readKeyFile(keyPath string, file http.FileSystem) *pem.Block {
	keyFile, err := file.Open(keyPath)
	if err != nil {
		log.Fatalf("Key file open error [%v]", keyPath)
	}
	defer keyFile.Close()

	pemFileInfo, _ := keyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(keyFile)
	_, err = buffer.Read(pemBytes)
	if err != nil {
		log.Fatalf("Key file read error [%v]", keyPath)
	}

	keyData, _ := pem.Decode([]byte(pemBytes))

	return keyData
}
