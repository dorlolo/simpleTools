// Package utils -----------------------------
// @file      : map_test.go
// @author    : JuneXu
// @contact   : 428192774@qq.com
// @time      : 2022/5/26 16:24
// -------------------------------------------
package utils

import "testing"

func TestGeoDistance(t *testing.T) {
	length := GeoDistance(120.651767, 31.374309, 120.632569, 31.431902)
	t.Log(length)
}
