// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumInt16 final model class
type FinalModelEnumInt16 struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelEnumInt16) FBEAllocationSize(value EnumInt16) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelEnumInt16) FBESize() int { return 2 }

// Get the final offset
func (fm FinalModelEnumInt16) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumInt16) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumInt16) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumInt16) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumInt16(buffer *fbe.Buffer, offset int) *FinalModelEnumInt16 {
    return &FinalModelEnumInt16{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelEnumInt16) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelEnumInt16) Get() (EnumInt16, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumInt16(0), 0, errors.New("model is broken")
    }

    return EnumInt16(fbe.ReadInt16(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumInt16) Set(value EnumInt16) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteInt16(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int16(value))
    return fm.FBESize(), nil
}
