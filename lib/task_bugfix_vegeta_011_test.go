package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta011SourceContract(t *testing.T) {
    source, err := os.ReadFile("results.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "seq++") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "seq--") {
        t.Fatalf("mutated source contract is still present")
    }
}
