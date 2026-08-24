package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta005SourceContract(t *testing.T) {
    source, err := os.ReadFile("attack.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if dial == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && dial == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
