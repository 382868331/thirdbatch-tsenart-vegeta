package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta003SourceContract(t *testing.T) {
    source, err := os.ReadFile("attack.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "case n == NoFollow:") {
        t.Fatalf("expected source contract is missing")
    }
}
