package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta019SourceContract(t *testing.T) {
    source, err := os.ReadFile("results.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if jw.Error != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
