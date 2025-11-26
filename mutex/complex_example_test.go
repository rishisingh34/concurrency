package mutex

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestComplexMutexExample(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe() 
	
	os.Stdout = w  
	
	ComplexMutexExample()
	
	_ = w.Close()
	
	result, _ := io.ReadAll(r)
	
	output := string(result) 
	
	os.Stdout = stdOut
	
	if ! strings.Contains(output, "$34320.00") {
		t.Error("wrong balance returned")
	}
}