package main

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta002SourceContract(t *testing.T) {
    source, err := os.ReadFile("flags.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if err == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
