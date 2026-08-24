package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta014SourceContract(t *testing.T) {
    source, err := os.ReadFile("attack.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "case ticks <- struct{}{}:") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "case ticks <=- struct{}{}:") {
        t.Fatalf("mutated source contract is still present")
    }
}
