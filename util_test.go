package dsbldr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBasicOAuthHeader(t *testing.T) {
	consumerKey := "consumerKey"
	nonce := "nonc