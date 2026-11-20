// SPDX-License-Identifier: GPL-2.0-or-later

package main

import (
	"net"
	"testing"

	"github.com/antoniocosta/unifi-openwrt/internal/translate"
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

func TestNetSig(t *testing.T) {
	a := translate.Op{Key: "network.lan.proto", Value: "dhcp", Old: "static"}
	b := translate.Op{Key: "network.lan.ipaddr", Del: true, Old: "192.168.1.2"}
	w := translate.Op{Key: "wireless.radio0.channel", Value: "6", Old: "1"}
	if netSig([]translate.Op{w}) != "" {
		t.Error("no network change: a signature")
	}
	if netSig([]translate.Op{a, w, b}) != netSig([]translate.Op{b, a}) {
		t.Error("the same network changes: different signatures")
	}
	if netSig([]translate.Op{a}) == netSig([]translate.Op{a, b}) {
		t.Error("different network changes: one signature")
	}
}
