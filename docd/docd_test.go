package docd

import (
	"fmt"
	"testing"
)

func TestInitDoc(t *testing.T) {
	d, err := InitDocd("test_doc", "doc.json", "testing", "dev")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("file created successfully: %s\n", d.FullPath)
}
