// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding

package fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

// Fast Binary Encoding UByte final model class
class FinalModelUInt8(buffer: Buffer, offset: Long) : FinalModel(buffer, offset)
{
    // Get the allocation size
    @Suppress("UNUSED_PARAMETER")
    fun FBEAllocationSize(value: UByte): Long = FBESize

    // Final size
    override val FBESize: Long = 1

    // Check if the value is valid
    override fun verify(): Long
    {
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return Long.MAX_VALUE

        return FBESize
    }

    // Get the value
    fun get(size: Size): UByte
    {
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return 0.toUByte()

        size.value = FBESize
        return readUInt8(FBEOffset)
    }

    // Set the value
    fun set(value: UByte): Long
    {
        assert(_buffer.offset + FBEOffset + FBESize <= _buffer.size) { "Model is broken!" }
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return 0

        write(FBEOffset, value)
        return FBESize
    }
}