// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "math"
import "time"
import "github.com/google/uuid"

// Fast Binary Encoding buffer based on dynamic byte array
type Buffer struct {
    data   []byte // Bytes memory buffer
    size   int    // Bytes memory buffer size
    offset int    // Bytes memory buffer offset
}

// Is the buffer empty?
func (b Buffer) Empty() bool { return (len(b.data) == 0) || (b.size <= 0) }
// Get bytes memory buffer
func (b Buffer) Data() []byte { return b.data }
// Get bytes memory buffer capacity
func (b Buffer) Capacity() int { return len(b.data) }
// Get bytes memory buffer size
func (b Buffer) Size() int { return b.size }
// Get bytes memory buffer offset
func (b Buffer) Offset() int { return b.offset }

// Create an empty buffer
func NewEmptyBuffer() *Buffer {
    return &Buffer{data: make([]byte, 0)}
}

// Create an empty buffer with a given capacity
func NewCapacityBuffer(capacity int) *Buffer {
    return &Buffer{data: make([]byte, capacity)}
}

// Create a buffer with attached bytes memory buffer
func NewAttachedBuffer(buffer []byte, offset int, size int) *Buffer {
    result := NewEmptyBuffer()
    result.AttachBuffer(buffer, offset, size)
    return result
}

// Attach an empty memory buffer
func (b *Buffer) AttachNew() {
    b.data = make([]byte, 0)
    b.size = 0
    b.offset = 0
}

// Attach an empty memory buffer with a given capacity
func (b *Buffer) AttachCapacity(capacity int) {
    b.data = make([]byte, capacity)
    b.size = 0
    b.offset = 0
}

// Attach a given memory buffer
func (b *Buffer) AttachBuffer(buffer []byte, offset int, size int) {
    if len(buffer) < size {
        panic("Invalid buffer!")
    }
    if size <= 0 {
        panic("Invalid size!")
    }
    if offset > size {
        panic("Invalid offset!")
    }

    b.data = buffer
    b.size = size
    b.offset = offset
}

// Allocate memory in the current write buffer and return offset to the allocated memory block
func (b *Buffer) Allocate(size int) int {
    if size < 0 {
        panic("Invalid allocation size!")
    }

    offset := b.size

    // Calculate a new buffer size
    total := b.size + size

    if total <= len(b.data) {
        b.size = total
        return offset
    }

    length := 2 * len(b.data)
    if length < total {
        length = total
    }

    data := make([]byte, length)
    copy(data, b.data[:b.size])
    b.data = data
    b.size = total
    return offset
}

// Remove some memory of the given size from the current write buffer
func (b *Buffer) Remove(offset int, size int) {
    if (offset + size) > len(b.data) {
        panic("Invalid offset & size!")
    }

    copy(b.data[offset:], b.data[offset+size:])
    b.size -= size
    if b.offset >= (offset + size) {
        b.offset -= size
    } else if b.offset >= offset {
        b.offset -= b.offset - offset
        if b.offset > b.size {
            b.offset = b.size
        }
    }
}

// Reserve memory of the given capacity in the current write bufferb
func (b *Buffer) Reserve(capacity int) {
    if capacity < 0 {
        panic("Invalid reserve capacity!")
    }

    if capacity > len(b.data) {
        length := 2 * len(b.data)
        if length < capacity {
            length = capacity
        }

        data := make([]byte, length)
        copy(data, b.data[:b.size])
        b.data = data
    }
}

// Resize the current write buffer
func (b *Buffer) Resize(size int) {
    b.Reserve(size)
    b.size = size
    if b.offset > b.size {
        b.offset = b.size
    }
}

// Reset the current write buffer and its offset
func (b *Buffer) Reset() {
    b.size = 0
    b.offset = 0
}

// Buffer I/O methods

func ReadBool(buffer []byte, offset int) bool {
    return buffer[offset] != 0
}

func ReadByte(buffer []byte, offset int) byte {
    return buffer[offset]
}

func ReadChar(buffer []byte, offset int) rune {
    return rune(ReadUInt8(buffer, offset))
}

func ReadWChar(buffer []byte, offset int) rune {
    return rune(ReadUInt32(buffer, offset))
}

func ReadInt8(buffer []byte, offset int) int8 {
    return int8(buffer[offset])
}

func ReadUInt8(buffer []byte, offset int) uint8 {
    return uint8(buffer[offset])
}

func ReadInt16(buffer []byte, offset int) int16 {
    return (int16(buffer[offset + 0]) << 0) | (int16(buffer[offset + 1]) << 8)
}

func ReadUInt16(buffer []byte, offset int) uint16 {
    return (uint16(buffer[offset + 0]) << 0) | (uint16(buffer[offset + 1]) << 8)
}

func ReadInt32(buffer []byte, offset int) int32 {
    return (int32(buffer[offset + 0]) <<  0) |
           (int32(buffer[offset + 1]) <<  8) |
           (int32(buffer[offset + 2]) << 16) |
           (int32(buffer[offset + 3]) << 24)
}

func ReadUInt32(buffer []byte, offset int) uint32 {
    return (uint32(buffer[offset + 0]) <<  0) |
           (uint32(buffer[offset + 1]) <<  8) |
           (uint32(buffer[offset + 2]) << 16) |
           (uint32(buffer[offset + 3]) << 24)
}

func ReadInt64(buffer []byte, offset int) int64 {
    return (int64(buffer[offset + 0]) <<  0) |
           (int64(buffer[offset + 1]) <<  8) |
           (int64(buffer[offset + 2]) << 16) |
           (int64(buffer[offset + 3]) << 24) |
           (int64(buffer[offset + 4]) << 32) |
           (int64(buffer[offset + 5]) << 40) |
           (int64(buffer[offset + 6]) << 48) |
           (int64(buffer[offset + 7]) << 56)
}

func ReadUInt64(buffer []byte, offset int) uint64 {
    return (uint64(buffer[offset + 0]) <<  0) |
           (uint64(buffer[offset + 1]) <<  8) |
           (uint64(buffer[offset + 2]) << 16) |
           (uint64(buffer[offset + 3]) << 24) |
           (uint64(buffer[offset + 4]) << 32) |
           (uint64(buffer[offset + 5]) << 40) |
           (uint64(buffer[offset + 6]) << 48) |
           (uint64(buffer[offset + 7]) << 56)
}

func ReadFloat(buffer []byte, offset int) float32 {
    bits := ReadUInt32(buffer, offset)
    return math.Float32frombits(bits)
}

func ReadDouble(buffer []byte, offset int) float64 {
    bits := ReadUInt64(buffer, offset)
    return math.Float64frombits(bits)
}

func ReadTimestamp(buffer []byte, offset int) time.Time {
    nanoseconds := ReadUInt64(buffer, offset)
    return time.Unix(int64(nanoseconds / 1000000000), int64(nanoseconds % 1000000000))
}

func ReadUUID(buffer []byte, offset int) uuid.UUID {
    bytes := ReadBytes(buffer, offset, 16)
    result, _ := uuid.FromBytes(bytes)
    return result
}

func ReadBytes(buffer []byte, offset int, size int) []byte {
    return buffer[offset:offset + size]
}

func ReadString(buffer []byte, offset int, size int) string {
    return string(buffer[offset:offset + size])
}

func WriteBool(buffer []byte, offset int, value bool) {
    if value {
        buffer[offset] = 1
    } else {
        buffer[offset] = 0
    }
}

func WriteByte(buffer []byte, offset int, value byte) {
    buffer[offset] = value
}

func WriteChar(buffer []byte, offset int, value rune) {
    WriteUInt8(buffer, offset, uint8(value))
}

func WriteWChar(buffer []byte, offset int, value rune) {
    WriteUInt32(buffer, offset, uint32(value))
}

func WriteInt8(buffer []byte, offset int, value int8) {
    buffer[offset] = byte(value)
}

func WriteUInt8(buffer []byte, offset int, value uint8) {
    buffer[offset] = byte(value)
}

func WriteInt16(buffer []byte, offset int, value int16) {
    buffer[offset + 0] = byte(value >> 0)
    buffer[offset + 1] = byte(value >> 8)
}

func WriteUInt16(buffer []byte, offset int, value uint16) {
    buffer[offset + 0] = byte(value >> 0)
    buffer[offset + 1] = byte(value >> 8)
}

func WriteInt32(buffer []byte, offset int, value int32) {
    buffer[offset + 0] = byte(value >>  0)
    buffer[offset + 1] = byte(value >>  8)
    buffer[offset + 2] = byte(value >> 16)
    buffer[offset + 3] = byte(value >> 24)
}

func WriteUInt32(buffer []byte, offset int, value uint32) {
    buffer[offset + 0] = byte(value >>  0)
    buffer[offset + 1] = byte(value >>  8)
    buffer[offset + 2] = byte(value >> 16)
    buffer[offset + 3] = byte(value >> 24)
}

func WriteInt64(buffer []byte, offset int, value int64) {
    buffer[offset + 0] = byte(value >>  0)
    buffer[offset + 1] = byte(value >>  8)
    buffer[offset + 2] = byte(value >> 16)
    buffer[offset + 3] = byte(value >> 24)
    buffer[offset + 4] = byte(value >> 32)
    buffer[offset + 5] = byte(value >> 40)
    buffer[offset + 6] = byte(value >> 48)
    buffer[offset + 7] = byte(value >> 56)
}

func WriteUInt64(buffer []byte, offset int, value uint64) {
    buffer[offset + 0] = byte(value >>  0)
    buffer[offset + 1] = byte(value >>  8)
    buffer[offset + 2] = byte(value >> 16)
    buffer[offset + 3] = byte(value >> 24)
    buffer[offset + 4] = byte(value >> 32)
    buffer[offset + 5] = byte(value >> 40)
    buffer[offset + 6] = byte(value >> 48)
    buffer[offset + 7] = byte(value >> 56)
}

func WriteFloat(buffer []byte, offset int, value float32) {
    WriteUInt32(buffer, offset, math.Float32bits(value))
}

func WriteDouble(buffer []byte, offset int, value float64) {
    WriteUInt64(buffer, offset, math.Float64bits(value))
}

func WriteTimestamp(buffer []byte, offset int, value time.Time) {
    nanoseconds := value.UnixNano()
    WriteUInt64(buffer, offset, uint64(nanoseconds))
}

func WriteUUID(buffer []byte, offset int, value uuid.UUID) {
    bytes, _ := value.MarshalBinary()
    WriteBytes(buffer, offset, bytes)
}

func WriteBytes(buffer []byte, offset int, value []byte) {
    copy(buffer[offset:offset + len(value)], value)
}

func WriteSlice(buffer []byte, offset int, value []byte, valueOffset int, valueSize int) {
    copy(buffer[offset:offset + len(value)], value[valueOffset:valueOffset + valueSize])
}

func WriteCount(buffer []byte, offset int, value byte, valueCount int) {
    for i := 0; i < valueCount; i++ {
        buffer[offset + i] = value
    }
}

func WriteString(buffer []byte, offset int, value string) {
    WriteBytes(buffer, offset, []byte(value))
}
