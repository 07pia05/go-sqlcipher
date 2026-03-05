package sqlite3

/*
#include <stdint.h>

// enable encryption codec in sqlite
#cgo CFLAGS: -DSQLITE_HAS_CODEC

// use memory for temporay storage in sqlite
#cgo CFLAGS: -DSQLITE_TEMP_STORE=2

// use libtomcrypt implementation in sqlcipher
// #cgo CFLAGS: -DSQLCIPHER_CRYPTO_LIBTOMCRYPT

// use OpenSSL implementation in sqlcipher
#cgo CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL

// disable assertions
#cgo CFLAGS: -DNDEBUG

// set operating specific sqlite flags
#cgo linux CFLAGS: -DSQLITE_OS_UNIX=1 -D_GNU_SOURCE
#cgo windows CFLAGS: -DSQLITE_OS_WIN=1

// SQLCipher extra init/shutdown
#cgo CFLAGS: -DSQLITE_EXTRA_INIT=sqlcipher_extra_init -DSQLITE_EXTRA_SHUTDOWN=sqlcipher_extra_shutdown

// Link OpenSSL for SQLCipher crypto
#cgo linux LDFLAGS: -lssl -lcrypto
*/
import "C"
