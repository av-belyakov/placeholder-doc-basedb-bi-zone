package datamodels

import "time"

type TestOptions struct {
	ValueMapAny      map[string]any
	ValueSliceAny    []any
	ValueSliceString []string
	ValueSliceUint64 []uint64
	ValueSliceInt    []int
	ValueSliceBool   []bool
	ValueTime        time.Time
	ValueAny         any
	ValueString      string
	ValueUint64      uint64
	ValueInt         int
	ValueBool        bool
	SetFunc          func()
	GetFunc          func()
}
