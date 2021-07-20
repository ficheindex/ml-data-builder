package main

/*
Demo server that spits out 1000 random records of a certain kind in a response
*/

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

const (
	recordCount    = 1000
	randomIntRange = 10000
)

func randIntStrings(count int) []string {
	output := make([]string,