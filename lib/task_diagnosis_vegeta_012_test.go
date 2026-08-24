package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisVegeta012SourceContract(t *testing.T) {
    source, err := os.ReadFile("targets.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return false") {
        t.Fatalf("expected source contract is missing")
    }
}
