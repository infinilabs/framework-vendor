package extsort

import (
	"bytes"
	"sort"
)

// Less compares byte chunks.
type Less func(a, b []byte) bool

// Equal compares two byte chunks for equality.
type Equal func(a, b []byte) bool

func stdLess(a, b []byte) bool {
	return bytes.Compare(a, b) < 0
}

// Options contains sorting options
type Options struct {
	// WorkDir specifies the working directory.
	// By default os.TempDir() is used.
	WorkDir string

	// Less defines the compare function.
	// Default: bytes.Compare() < 0
	Less Less

	// Sort defines the sort function that is used.
	// Default: sort.Sort
	Sort func(sort.Interface)

	// Dedupe defines the compare function for de-duplication.
	// Default: nil (= do not de-dupe)
	Dedupe Equal

	// BufferSize limits the memory buffer used for sorting.
	// Default: 64MiB (must be at least 64KiB)
	BufferSize int

	// Compression optionally uses compression for temporary output.
	Compression Compression
}

func (o *Options) norm() *Options {
	var opt Options
	if o != nil {
		opt = *o
	}

	if opt.Less == nil {
		opt.Less = stdLess
	}

	if opt.Sort == nil {
		opt.Sort = sort.Sort
	}

	if std := (1 << 26); opt.BufferSize < 1 {
		opt.BufferSize = std
	} else if min := (1 << 16); opt.BufferSize < min {
		opt.BufferSize = min
	}

	opt.Compression = opt.Compression.norm()

	return &opt
}
