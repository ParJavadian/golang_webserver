package service

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const pgMethodName = "pg"
const dhMethodName = "dh"

type PgParams struct {
	P           int
	G           int
	Nonce       string
	ServerNonce string
	MessageId   int
}

func (pg PgParams) String() string {
	return fmt.Sprintf("%d %d %s %s %d", pg.P, pg.G, pg.Nonce, pg.ServerNonce, pg.MessageId)
}

type DHParams struct {
	Nonce       string
	ServerNonce string
	MessageId   int
	PublicKey   int
}

func GetPg(nonce string, requestMessageId int) PgParams {
	// generate the response
	responseMessageId := randomOddInt()
	serverNonce := randomString(20)
	pgResponse := PgParams{23, 5, nonce, serverNonce, responseMessageId}

	// save to cache
	cacheKey := getCacheKeyPg(nonce, serverNonce, pgMethodName)
	CacheData(nil, cacheKey, pgResponse.String(), 20*time.Minute)

	return pgResponse
}

func GetDHParams(nonce string, serverNonce string, messageId int, requestPublicKey int) (DHParams, error) {
	b := randomInt()
	// get PgParams from cache
	pgCacheKey := getCacheKeyPg(nonce, serverNonce, pgMethodName)
	pgParamsString, err := GetValue(nil, pgCacheKey)
	if err != nil {
		return DHParams{}, fmt.Errorf("pgParams not found or expired for nonce %s and serverNonce %s", nonce, serverNonce)
	}

	pgParams, err := getPgParamsFromString(pgParamsString)
	if err != nil {
		return DHParams{}, fmt.Errorf("pgParams not found or expired for nonce %s and serverNonce %s", nonce, serverNonce)
	}

	responsePublicKey := (pgParams.G ^ b) % pgParams.P
	commonKey := (requestPublicKey ^ b) % pgParams.P

	dhCacheKey := getCacheKeyDh(string(rune(commonKey)), dhMethodName)
	dhCacheValue := getCacheValuePg(nonce, serverNonce)
	CacheData(nil, dhCacheKey, dhCacheValue, 20*time.Minute)

	responseMessageId := randomOddInt()
	dhParams := DHParams{
		Nonce:       nonce,
		ServerNonce: serverNonce,
		MessageId:   responseMessageId,
		PublicKey:   responsePublicKey,
	}
	return dhParams, nil
}

func getPgParamsFromString(pgParamsString string) (PgParams, error) {
	pgParams := PgParams{}
	_, err := fmt.Sscanf(
		pgParamsString,
		"%d %d %s %s %d",
		&pgParams.P,
		&pgParams.G,
		&pgParams.Nonce,
		&pgParams.ServerNonce,
		&pgParams.MessageId,
	)
	if err != nil {
		fmt.Println(err)
	}
	return pgParams, err

}

func getCacheKeyPg(nonce string, serverNonce string, methodName string) string {
	return methodName + "_" + hashWithSHA1(nonce+serverNonce)
}

func getCacheKeyDh(authKey string, methodName string) string {
	return methodName + "_" + authKey
}

func getCacheValuePg(nonce string, serverNonce string) string {
	return nonce + " " + serverNonce
}

func randomString(length int) string {
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func randomInt() int {
	return rand.Int()
}

func randomOddInt() int {
	return 2*randomInt() + 1
}

func hashWithSHA1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	hashInBytes := hash.Sum(nil)
	return string(hashInBytes)
}
