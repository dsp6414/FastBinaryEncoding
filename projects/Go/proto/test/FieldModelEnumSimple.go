// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"

// Fast Binary Encoding EnumSimple field model class
type FieldModelEnumSimple struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm FieldModelEnumSimple) FBESize() int { return 4 }
// Get the field extra size
func (fm FieldModelEnumSimple) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelEnumSimple) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumSimple) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumSimple) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumSimple) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelEnumSimple(buffer *fbe.Buffer, offset int) *FieldModelEnumSimple {
    return &FieldModelEnumSimple{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FieldModelEnumSimple) Verify() bool { return true }

// Get the value
func (fm FieldModelEnumSimple) Get() (EnumSimple, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm FieldModelEnumSimple) GetDefault(defaults EnumSimple) (EnumSimple, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumSimple(0), nil
    }

    return EnumSimple(fbe.ReadInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelEnumSimple) Set(value EnumSimple) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int32(value))
    return nil
}
