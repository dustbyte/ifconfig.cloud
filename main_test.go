package main

import "testing"

func TestXForwardedForEmpty(t *testing.T) {
	result := fromXForwardedFor("")
	if result != "" {
		t.Errorf("Expected empty result")
	}
}

// Shouldn't happen, but never trust proxy input
func TestXForwardedForNoComma(t *testing.T) {
	result := fromXForwardedFor("127.0.0.1")
	if result != "127.0.0.1" {
		t.Errorf("Expected 127.0.0.1")
	}
}

func TestXForwardedFor(t *testing.T) {
	result := fromXForwardedFor("127.0.0.1,192.168.1.1")
	if result != "127.0.0.1" {
		t.Errorf("Expected 127.0.0.1")
	}
}

func TestEmptyRemoteAddr(t *testing.T) {
	result := fromRemoteAddr("")
	if result != "" {
		t.Errorf("Expected empty result")
	}
}

func TestRemoteAddrNoPort(t *testing.T) {
	result := fromRemoteAddr("127.0.0.1")
	if result != "127.0.0.1" {
		t.Errorf("Expected 127.0.0.1")
	}
}

func TestRemoteAddr(t *testing.T) {
	result := fromRemoteAddr("127.0.0.1:1234")
	if result != "127.0.0.1" {
		t.Errorf("Expected 127.0.0.1")
	}
}
