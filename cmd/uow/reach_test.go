// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"net"
	"testing"
)

func TestReachable(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	url := "http://" + l.Addr().String() + "/inform"
	if err := reachable(url); err != nil {
		t.Errorf("listening: %v", err)
	}
	l.Close()
	if reachable(url) == nil {
		t.Error("closed port: reachable")
	}
	if reachable("not a url") == nil {
		t.Error("bad URL: reachable")
	}
}
