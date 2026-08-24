package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta001SourceContract(t *testing.T) {
    source, err := os.ReadFile("histogram.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(value) < 2 || value[0] != '[' || value[len(value)-1] != ']' {") {
        t.Fatalf("expected source contract is missing")
    }
}
