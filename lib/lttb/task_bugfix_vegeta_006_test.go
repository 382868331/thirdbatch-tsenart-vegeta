package lttb

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta006SourceContract(t *testing.T) {
    source, err := os.ReadFile("lttb.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if threshold >= count || threshold == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
