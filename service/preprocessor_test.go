package service

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

var (
	contextData []byte
	bookData    []byte
	resultData  []byte
)

func init() {
	contextData, _ = os.ReadFile("test/context.json")
	bookData, _ = os.ReadFile("test/book.json")
	resultData, _ = os.ReadFile("test/result.json")
}

func TestProcessBook(t *testing.T) {
	tested := NewPreprocessor()

	var context map[string]any
	json.Unmarshal(contextData, &context)
	var book map[string]any
	json.Unmarshal(bookData, &book)

	result := tested.ProcessBook(context, book)

	var expectedResult map[string]any
	json.Unmarshal(resultData, &expectedResult)

	if !reflect.DeepEqual(result, expectedResult) {
		r, _ := json.MarshalIndent(result, "", "  ")
		e, _ := json.MarshalIndent(expectedResult, "", "  ")

		t.Errorf(`Processed book did not match the expected output!
Result:
%s
---
Expected:
%s`, string(r), string(e))
	}
}
