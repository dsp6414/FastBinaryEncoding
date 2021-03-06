// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "errors"

// Fast Binary Encoding int32 final model class
type FinalModelInt32 struct {
    buffer *Buffer  // Final model buffer
    offset int      // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelInt32) FBEAllocationSize(value int32) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelInt32) FBESize() int { return 4 }

// Get the final offset
func (fm FinalModelInt32) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelInt32) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelInt32) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelInt32) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelInt32(buffer *Buffer, offset int) *FinalModelInt32 {
    return &FinalModelInt32{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelInt32) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelInt32) Get() (int32, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, 0, errors.New("model is broken")
    }

    return ReadInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelInt32) Set(value int32) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    WriteInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize(), nil
}
