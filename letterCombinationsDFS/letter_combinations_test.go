package letterCombinationsDFS

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	expected := []string{"ad","ae","af","bd","be","bf","cd","ce","cf"}
	res := LetterCombinations("23")
	
	// 比较两个slice不能直接比较
	if !reflect.DeepEqual(expected, res) {
		t.Errorf("expected '%s' got '%s'", expected, res)
	}
}