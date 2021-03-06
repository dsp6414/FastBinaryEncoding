// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumWChar field model class
type FieldModelEnumWChar struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm FieldModelEnumWChar) FBESize() int { return 4 }
// Get the field extra size
func (fm FieldModelEnumWChar) FBEExtra() int { return 0 }

// Get the field offset
func (fm FieldModelEnumWChar) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumWChar) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumWChar) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumWChar) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelEnumWChar(buffer *fbe.Buffer, offset int) *FieldModelEnumWChar {
    return &FieldModelEnumWChar{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FieldModelEnumWChar) Verify() bool { return true }

// Get the value
func (fm FieldModelEnumWChar) Get() (EnumWChar, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm FieldModelEnumWChar) GetDefault(defaults EnumWChar) (EnumWChar, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumWChar(0), nil
    }

    return EnumWChar(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelEnumWChar) Set(value EnumWChar) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(value))
    return nil
}
