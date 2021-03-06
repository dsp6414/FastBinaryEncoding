// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "errors"

// Fast Binary Encoding float64 final model class
type FinalModelDouble struct {
    buffer *Buffer  // Final model buffer
    offset int      // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelDouble) FBEAllocationSize(value float64) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelDouble) FBESize() int { return 8 }

// Get the final offset
func (fm FinalModelDouble) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelDouble) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelDouble) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelDouble) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelDouble(buffer *Buffer, offset int) *FinalModelDouble {
    return &FinalModelDouble{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelDouble) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelDouble) Get() (float64, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0.0, 0, errors.New("model is broken")
    }

    return ReadDouble(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelDouble) Set(value float64) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    WriteDouble(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize(), nil
}
