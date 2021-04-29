package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5HashOrig : return hashed input as string
//  this is the version without the trailing '\n'
func MD5HashOrig(m string) string {

	mHashed := md5.Sum([]byte(m)) // returns [16]byte
	mOut := hex.EncodeToString(mHashed[:])

	return mOut
}

// MD5Hash : return hashed input as string
//  this is the version which accomodates `echo` logic
func MD5Hash(m string) string {

	mByte := []byte(m + "\n") // returns [16]byte
	mHashed := md5.Sum(mByte)
	mOut := hex.EncodeToString(mHashed[:])

	return mOut
}
