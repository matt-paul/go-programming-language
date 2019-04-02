package allinclusive

import "testing"

func TestContainAllRots(t *testing.T) {

	t.Run("returns true when string matches the any item in the array", func(t *testing.T) {
		testSliceLiteral := []string{"qbsj", "bsjq", "sjqb", "twZNsslC", "jqbs"}
		result := ContainAllRots("bsjq", testSliceLiteral)
		if !result {
			t.Fatalf("Oh no, this has has gone awfully wrong hasn't it?")
		}
	})

	t.Run("returns false if the string is not in the array", func(t *testing.T) {
		testSliceLiteral := []string{"aaa", "bbb", "cccc"}
		result := ContainAllRots("bsjq", testSliceLiteral)
		if result {
			t.Fatalf("This value should not be in our array")
		}
	})

	t.Run("returns false when one rotation is missing", func(t *testing.T) {
		testSliceLiteral := []string{"bsjq", "sjqb", "jqbs", "nono"}
		result := ContainAllRots("bsjq", testSliceLiteral)
		if result {
			t.Fatalf("Work to do!")
		}
	})

	t.Run("returns true when all rotations are included", func(t *testing.T) {
		testSliceLiteral := []string{"bsjq", "sjqb", "jqbs", "qbsj"}
		result := ContainAllRots("bsjq", testSliceLiteral)
		if !result {
			t.Fatalf("the function has not found enlightenment")
		}
	})
}
