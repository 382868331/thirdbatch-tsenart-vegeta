package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisVegeta004SourceContract(t *testing.T) {
    source, err := os.ReadFile("attack.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if ip == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if ip != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
