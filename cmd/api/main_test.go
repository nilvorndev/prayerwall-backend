package main

import "testing"

func TestEnvOr(t *testing.T) {
	t.Setenv("PRAYERWALL_TEST_VALUE", "configured")
	if got := envOr("PRAYERWALL_TEST_VALUE", "fallback"); got != "configured" {
		t.Fatalf("envOr() = %q, want %q", got, "configured")
	}
	if got := envOr("PRAYERWALL_MISSING_VALUE", "fallback"); got != "fallback" {
		t.Fatalf("envOr() = %q, want %q", got, "fallback")
	}
}
