package cli

// Args is the interface for arguments which are appended to a command
type Args interface {
	// Get returns the nth argument, or else a blank string
	Get(n int) string
	// First returns the first argument, or else a blank string
	First() string
	// Tail returns the rest of the arguments (not the first one)
	// or else an empty string slice
	Tail() []string
	// Len returns the length of the wrapped slice
	Len() int
	// Present checks if there are any arguments present
	Present() bool
	// Slice returns a copy of the internal slice
	Slice() []string
}

type args []string

// Get returns the nth argument, or else a blank string
func (a *args) Get(n int) string {
	if len(*a) > n {
		return (*a)[n]
	}
	return ""
}

// First returns the first argument, or else a blank string
func (a *args) First() string {
	return a.Get(0)
}

// Tail returns the rest of the arguments (not the first one)
// or else an empty string slice
func (a *args) Tail() []string {
	if a.Len() >= 2 {
		tail := []string((*a)[1:])
		ret := make([]string, len(tail))
		copy(ret, tail)
		return ret
	}
	return []string{}
}

// Len returns the length of the wrapped slice
func (a *args) Len() int {
	return len(*a)
}

// Present checks if there are any arguments present
func (a *args) Present() bool {
	return a.Len() != 0
}

// Slice returns a copy of the internal slice
func (a *args) Slice() []string {
	ret := make([]string, len(*a))
	copy(ret, *a)
	return ret
}
