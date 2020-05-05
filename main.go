package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
)

var (
	privateKeyFile = "pk-APKAIOKHNLXZIO3FRZ3Q.pem"
	ACCESS_KEY     = "PPKOIOKHKLXZLQ3FRZ3A"

	PRIVATE_KEY *rsa.PrivateKey

	cdnDomain = "d2sjg3xny2v4pz.cloudfront.net"
	s3Domain  = "shub-storage-siujq.s3.ap-southeast-1.amazonaws.com"
	s3URL     = "https://shub-storage-siujq.s3.ap-southeast-1.amazonaws.com/tests/559350/file_url/Bang_Con_Vow.pdf"
	cdnRawURL = strings.Replace(s3URL, s3Domain, cdnDomain, -1)
)

func init() {
	data, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalf("Read private key error: %v", err)
	}

	pemData, _ := pem.Decode([]byte(data))
	PRIVATE_KEY, err = x509.ParsePKCS1PrivateKey(pemData.Bytes)
	if err != nil {
		log.Fatal("Parse private key error: %v", err)
	}

}

func main() {
	signer := sign.NewURLSigner(ACCESS_KEY, PRIVATE_KEY)
	expireTime := time.Hour * 48

	signedURL, err := signer.Sign(cdnRawURL, time.Now().Add(expireTime))
	if err != nil {
		log.Fatalf("Create CDN url error: %v", err)
	}

	log.Println("CDN_URL", signedURL)
}
