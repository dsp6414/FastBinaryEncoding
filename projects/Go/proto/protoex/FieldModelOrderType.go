// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.1.0.0

package protoex

import "errors"
import "../fbe"

// Fast Binary Encoding OrderType field model class
type FieldModelOrderType struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm FieldModelOrderType) FBESize() int { return 1 }
// Get the field extra size
func (fm FieldModelOrderType) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelOrderType) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelOrderType) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelOrderType) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelOrderType) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelOrderType(buffer *fbe.Buffer, offset int) *FieldModelOrderType {
    return &FieldModelOrderType{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FieldModelOrderType) Verify() bool { return true }

// Get the value
func (fm FieldModelOrderType) Get() (OrderType, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm FieldModelOrderType) GetDefault(defaults OrderType) (OrderType, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return OrderType(0), nil
    }

    return OrderType(fbe.ReadByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelOrderType) Set(value OrderType) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), byte(value))
    return nil
}
