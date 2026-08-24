package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisVegeta020SourceContract(t *testing.T) {
    source, err := os.ReadFile("results.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "bytes.Equal(r.Body, other.Body) &&") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "bytes.Equal(r.Body, other.Body) ||") {
        t.Fatalf("mutated source contract is still present")
    }
}
