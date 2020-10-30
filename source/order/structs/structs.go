package structs

import (
	"fmt"
)

func init() {
	fmt.Println("package: order.structs - initialized")
}

// HTTPError struct
type HTTPError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

// SecretBody struct
type SecretBody struct {
	SecretName string `json:"secretName"`
}
