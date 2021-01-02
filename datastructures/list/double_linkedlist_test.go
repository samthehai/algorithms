package list

import "testing"

func TestDoubleLinkedListSize(t *testing.T) {
	l := New()

	if got, want := l.Size(), 0; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListIsEmpty(t *testing.T) {
	l := New()

	if got, want := l.Size(), 0; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListString(t *testing.T) {
	l := New()
	l.Add("b")

	if got, want := l.String(), "[b]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListAdd(t *testing.T) {
	l := New()

	l.Add("a")
	l.Add("b")
	l.Add("c")
	l.Add("d")
	if got, want := l.Size(), 4; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.String(), "[a, b, c, d]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListAddLast(t *testing.T) {
	l := New()

	l.AddLast("a")
	l.AddLast("b")
	l.AddLast("c")
	l.AddLast("d")
	if got, want := l.Size(), 4; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.String(), "[a, b, c, d]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListAddFirst(t *testing.T) {
	l := New()

	l.AddFirst("a")
	l.AddFirst("b")
	l.AddFirst("c")
	l.AddFirst("d")
	if got, want := l.Size(), 4; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.String(), "[d, c, b, a]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListAddAt(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	l.AddAt(1, "d")
	if got, want := l.Size(), 4; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.String(), "[a, d, b, c]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListPeekFirst(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	l.AddAt(1, "d")
	first, err := l.PeekFirst()
	if got, want := first, "a"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListPeekFirst_Empty(t *testing.T) {
	l := New()

	first, err := l.PeekFirst()
	if got, want := first, interface{}(nil); got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err != nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListPeekLast(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	first, err := l.PeekLast()
	if got, want := first, "c"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListPeekLast_Empty(t *testing.T) {
	l := New()

	last, err := l.PeekLast()
	if got, want := last, interface{}(nil); got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err != nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListRemoveFirst(t *testing.T) {
	l := New()
	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")

	first, err := l.RemoveFirst()
	if got, want := first, "a"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.String(), "[b, c]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Size(), 2; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListRemoveFirst_Empty(t *testing.T) {
	l := New()

	first, err := l.RemoveFirst()
	if got, want := first, interface{}(nil); got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err != nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListRemoveLast(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	first, err := l.RemoveLast()
	if got, want := first, "c"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.String(), "[a, b]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Size(), 2; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListRemoveLast_Empty(t *testing.T) {
	l := New()

	first, err := l.RemoveLast()
	if got, want := first, interface{}(nil); got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err != nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListIndexOf(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	if got, want := l.IndexOf("a"), 0; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.IndexOf("b"), 1; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.IndexOf("c"), 2; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.IndexOf("z"), -1; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListRemove(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	err := l.Remove("a")
	if got, want := l.String(), "[b, c]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}

	err = l.Remove("c")
	if got, want := l.String(), "[b]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}

	err = l.Remove("b")
	if got, want := l.String(), "[]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListRemoveAt(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	v, err := l.RemoveAt(0)
	if got, want := v, "a"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}

	v, err = l.RemoveAt(1)
	if got, want := v, "c"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}

	v, err = l.RemoveAt(1)
	if got, want := v, interface{}(nil); got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err != nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}

	v, err = l.RemoveAt(0)
	if got, want := v, "b"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := err == nil, true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}

	if got, want := l.String(), "[]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListContains(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	if got, want := l.Contains("a"), true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Contains("b"), true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Contains("c"), true; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Contains("d"), false; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Contains(nil), false; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Contains(1), false; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestDoubleLinkedListClear(t *testing.T) {
	l := New()

	l.AddAt(0, "a")
	l.AddAt(1, "b")
	l.AddAt(2, "c")
	l.Clear()
	if got, want := l.String(), "[]"; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
	if got, want := l.Size(), 0; got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}
