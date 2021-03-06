// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumInt8 final model class
type FinalModelEnumInt8 struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelEnumInt8) FBEAllocationSize(value EnumInt8) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelEnumInt8) FBESize() int { return 1 }

// Get the final offset
func (fm FinalModelEnumInt8) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumInt8) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumInt8) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumInt8) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumInt8(buffer *fbe.Buffer, offset int) *FinalModelEnumInt8 {
    return &FinalModelEnumInt8{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelEnumInt8) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelEnumInt8) Get() (EnumInt8, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumInt8(0), 0, errors.New("model is broken")
    }

    return EnumInt8(fbe.ReadInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumInt8) Set(value EnumInt8) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int8(value))
    return fm.FBESize(), nil
}
