package main

import "github.com/Nv7-Github/screenlapse/encoders"

var encs map[string]func() encoders.Encoder = map[string]func() encoders.Encoder{"PNG": encoders.NewPNGEncoder}
