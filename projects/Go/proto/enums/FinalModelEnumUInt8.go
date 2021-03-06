// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumUInt8 final model class
type FinalModelEnumUInt8 struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelEnumUInt8) FBEAllocationSize(value EnumUInt8) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelEnumUInt8) FBESize() int { return 1 }

// Get the final offset
func (fm FinalModelEnumUInt8) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumUInt8) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumUInt8) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumUInt8) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumUInt8(buffer *fbe.Buffer, offset int) *FinalModelEnumUInt8 {
    return &FinalModelEnumUInt8{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelEnumUInt8) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelEnumUInt8) Get() (EnumUInt8, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumUInt8(0), 0, errors.New("model is broken")
    }

    return EnumUInt8(fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumUInt8) Set(value EnumUInt8) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint8(value))
    return fm.FBESize(), nil
}
