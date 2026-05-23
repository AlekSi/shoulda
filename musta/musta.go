package musta

import "github.com/AlekSi/shoulda"

//go:generate go run ../internal/mustagen

// TB is a subset of [testing.TB] that is used by this package.
type TB = shoulda.TB
