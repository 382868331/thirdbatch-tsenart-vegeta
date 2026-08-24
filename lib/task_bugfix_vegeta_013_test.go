package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixVegeta013SourceContract(t *testing.T) {
    source, err := os.ReadFile("pacer.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return sp.Period <= 0 || sp.Mean.hitsPerNs() <= 0 || sp.Amp.hitsPerNs() >= sp.Mean.hitsPerNs()") {
        t.Fatalf("expected source contract is missing")
    }
}
