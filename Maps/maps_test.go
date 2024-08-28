package maps

import "testing"

func MapsTest (t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("key exists", func (t * testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("key not exists", func (t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Errorf("expected to get an error")
		}
	})

	t.Run("add and check", func (t *testing.T) {
		dictionary.Add("test2", "this is a test2")

		got, _ := dictionary.Search("test2")
		want := "this is a test2"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}