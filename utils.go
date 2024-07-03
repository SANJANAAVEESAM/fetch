package main

import (
    "crypto/rand"
    "encoding/hex"
)

func generateID() string {
    bytes := make([]byte, 16)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}
