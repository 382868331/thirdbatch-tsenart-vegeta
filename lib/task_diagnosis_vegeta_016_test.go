package vegeta

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisVegeta016SourceContract(t *testing.T) {
    source, err := os.ReadFile("targets.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err = jl.Error(); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if err = jl.Error(); err == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
