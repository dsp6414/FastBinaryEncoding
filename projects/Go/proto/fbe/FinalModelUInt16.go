// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "errors"

// Fast Binary Encoding uint16 final model class
type FinalModelUInt16 struct {
    buffer *Buffer  // Final model buffer
    offset int      // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelUInt16) FBEAllocationSize(value uint16) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelUInt16) FBESize() int { return 2 }

// Get the final offset
func (fm FinalModelUInt16) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelUInt16) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelUInt16) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelUInt16) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelUInt16(buffer *Buffer, offset int) *FinalModelUInt16 {
    return &FinalModelUInt16{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelUInt16) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelUInt16) Get() (uint16, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, 0, errors.New("model is broken")
    }

    return ReadUInt16(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelUInt16) Set(value uint16) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    WriteUInt16(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize(), nil
}
