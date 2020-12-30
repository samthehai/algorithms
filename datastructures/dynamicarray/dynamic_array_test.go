package dynamicarray

import "testing"

func TestSize(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := a.Size(), 0; got != want {
		t.Errorf("Size = %v, want %v", got, want)
	}
}

func TestAdd(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	a.Add("a")
	a.Add("b")
	if got, want := a.Size(), 2; got != want {
		t.Errorf("Size = %v, want %v", got, want)
	}
	got, want := a.arr, []string{"a", "b"}
	for i, v := range got {
		if v.(string) != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestRemove(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	a.Add("a")
	a.Add("b")
	a.Remove("a")
	if got, want := a.Size(), 1; got != want {
		t.Errorf("Size = %v, want %v", got, want)
	}
	got, want := a.arr, []string{"b"}
	for i, v := range got {
		if v.(string) != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestRemoveAt(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	a.Add("a")
	a.Add("b")
	a.RemoveAt(0)
	if got, want := a.Size(), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	got, want := a.arr, []string{"b"}
	for i, v := range got {
		if v.(string) != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestIndexOf(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	a.Add("a")
	a.Add("b")
	a.Add("c")
	a.Add(nil)
	if got, want := a.IndexOf("a"), 0; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.IndexOf("b"), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.IndexOf("c"), 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.IndexOf(nil), 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.IndexOf((*string)(nil)), -1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.IndexOf("z"), -1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestContains(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	a.Add("a")
	a.Add("b")
	if got, want := a.Contains("a"), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.Contains("b"), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.Contains("c"), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := a.Contains(1), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestString(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	a.Add("a")
	a.Add("b")
	if got, want := a.String(), "[a],[b]"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
