package sandal

import (
	"fmt"
	"github.com/draftcode/sandal/lang"
	"io/ioutil"
	"testing"
)

var sampleFiles = []string{
	"two_phase_commit.sandal",
	// "sample2.sandal",
	// "sample.sandal",
}

const ROOT_PATH = "/home/draftcode/src/go/src/github.com/draftcode/sandal/"

func TestCompileSamples(t *testing.T) {
	for _, sampleFile := range sampleFiles {
		filePath := ROOT_PATH + sampleFile
		body, err := ioutil.ReadFile(filePath)
		if err != nil {
			t.Fatal(sampleFile, err)
		}
		err, compiled := lang.CompileFile(string(body))
		if err != nil {
			t.Fatal(sampleFile, err)
		}
		fmt.Print(compiled)
	}
}
