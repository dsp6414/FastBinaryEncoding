// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.1.0.0

package protoex

import "errors"
import "../fbe"

// Fast Binary Encoding OrderType final model class
type FinalModelOrderType struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelOrderType) FBEAllocationSize(value OrderType) int { return fm.FBESize() }

// Get the final size
func (fm FinalModelOrderType) FBESize() int { return 1 }

// Get the final offset
func (fm FinalModelOrderType) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelOrderType) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelOrderType) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelOrderType) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelOrderType(buffer *fbe.Buffer, offset int) *FinalModelOrderType {
    return &FinalModelOrderType{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelOrderType) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm FinalModelOrderType) Get() (OrderType, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return OrderType(0), 0, errors.New("model is broken")
    }

    return OrderType(fbe.ReadByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelOrderType) Set(value OrderType) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), byte(value))
    return fm.FBESize(), nil
}
