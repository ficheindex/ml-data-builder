package main

/*
Demo server that spits out 1000 random records of a certain kind in a response
*/

import (
	"encoding/json"
	"math/rand"
	"net/