package copier

import (
	"testing"
)

// Test case for issue where case-insensitive field matching
// fails when there are both exported and unexported fields with similar names
func TestCaseInsensitiveFieldMatching(t *testing.T) {
	// Simulate protobuf generated struct with both unexported 'state' and exported 'State' fields
	type ProtoStruct struct {
		state      uint32
		PageNumber int32 `json:"pageNumber"`
		State      int32 `json:"state"`
	}

	type SourceStruct struct {
		PageNumber int32 `json:"pageNumber"`
		State      int32 `json:"state"`
	}

	source := SourceStruct{
		PageNumber: 1,
		State:      99,
	}

	dest := &ProtoStruct{}

	err := Copy(dest, &source)
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}

	// Verify that the exported State field was copied correctly
	if dest.State != source.State {
		t.Errorf("State field not copied correctly. Expected: %d, Got: %d", source.State, dest.State)
	}

	if dest.state == uint32(source.State) {
		t.Errorf("state field not copied correctly. Expected: %d, Got: %d", source.State, dest.state)
	}

	if dest.PageNumber != source.PageNumber {
		t.Errorf("PageNumber field not copied correctly. Expected: %d, Got: %d", source.PageNumber, dest.PageNumber)
	}
}

// Test that exact case matching still works
func TestExactCaseMatching(t *testing.T) {
	type Source struct {
		Name string
		Age  int
	}

	type Dest struct {
		Name string
		Age  int
	}

	source := Source{Name: "John", Age: 30}
	dest := &Dest{}

	err := Copy(dest, &source)
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}

	if dest.Name != source.Name || dest.Age != source.Age {
		t.Errorf("Fields not copied correctly. Expected: %+v, Got: %+v", source, *dest)
	}
}
