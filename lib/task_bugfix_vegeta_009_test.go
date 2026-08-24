package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta009SourceContract(t *testing.T) {
    source, err := os.ReadFile("pacer.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "case p.StartAt.Per < 0 || p.StartAt.Freq < 0:") {
        t.Fatalf("expected source contract is missing")
    }
}
