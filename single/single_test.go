package single

import "testing"

func TestGetInstance(t *testing.T) {

	t.Helper()
	instance1 := GetInstance()
	instance2 := GetInstance()

	t.Run("is_true", func(t *testing.T) {
		b := instance1 == instance2
		if !b {
			t.Error("error")
		}
	})
}
