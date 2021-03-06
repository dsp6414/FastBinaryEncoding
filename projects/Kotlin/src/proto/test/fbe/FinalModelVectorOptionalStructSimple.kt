// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

@file:Suppress("UnusedImport", "unused")

package test.fbe

import java.io.*
import java.lang.*
import java.lang.reflect.*
import java.math.*
import java.nio.charset.*
import java.time.*
import java.util.*

import fbe.*
import test.*

// Fast Binary Encoding OptionalStructSimple vector final model class
class FinalModelVectorOptionalStructSimple(buffer: Buffer, offset: Long) : FinalModel(buffer, offset)
{
    private val _model = FinalModelOptionalStructSimple(buffer, offset)

    // Get the allocation size
    fun fbeAllocationSize(values: ArrayList<StructSimple?>): Long
    {
        var size: Long = 4
        for (value in values)
            size += _model.fbeAllocationSize(value)
        return size
    }
    fun fbeAllocationSize(values: LinkedList<StructSimple?>): Long
    {
        var size: Long = 4
        for (value in values)
            size += _model.fbeAllocationSize(value)
        return size
    }
    fun fbeAllocationSize(values: HashSet<StructSimple?>): Long
    {
        var size: Long = 4
        for (value in values)
            size += _model.fbeAllocationSize(value)
        return size
    }

    // Check if the vector is valid
    override fun verify(): Long
    {
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return Long.MAX_VALUE

        val fbeVectorSize = readUInt32(fbeOffset).toLong()

        var size: Long = 4
        _model.fbeOffset = fbeOffset + 4
        var i = fbeVectorSize
        while (i-- > 0)
        {
            val offset = _model.verify()
            if (offset == Long.MAX_VALUE)
                return Long.MAX_VALUE
            _model.fbeShift(offset)
            size += offset
        }
        return size
    }

    // Get the vector as ArrayList
    fun get(values: ArrayList<StructSimple?>): Long
    {
        values.clear()

        assert((_buffer.offset + fbeOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return 0

        val fbeVectorSize = readUInt32(fbeOffset).toLong()
        if (fbeVectorSize == 0L)
            return 4

        values.ensureCapacity(fbeVectorSize.toInt())

        var size: Long = 4
        val offset = Size()
        _model.fbeOffset = fbeOffset + 4
        for (i in 0 until fbeVectorSize)
        {
            offset.value = 0
            val value = _model.get(offset)
            values.add(value)
            _model.fbeShift(offset.value)
            size += offset.value
        }
        return size
    }

    // Get the vector as LinkedList
    fun get(values: LinkedList<StructSimple?>): Long
    {
        values.clear()

        assert((_buffer.offset + fbeOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return 0

        val fbeVectorSize = readUInt32(fbeOffset).toLong()
        if (fbeVectorSize == 0L)
            return 4

        var size: Long = 4
        val offset = Size()
        _model.fbeOffset = fbeOffset + 4
        for (i in 0 until fbeVectorSize)
        {
            offset.value = 0
            val value = _model.get(offset)
            values.add(value)
            _model.fbeShift(offset.value)
            size += offset.value
        }
        return size
    }

    // Get the vector as HashSet
    fun get(values: HashSet<StructSimple?>): Long
    {
        values.clear()

        assert((_buffer.offset + fbeOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return 0

        val fbeVectorSize = readUInt32(fbeOffset).toLong()
        if (fbeVectorSize == 0L)
            return 4

        var size: Long = 4
        val offset = Size()
        _model.fbeOffset = fbeOffset + 4
        for (i in 0 until fbeVectorSize)
        {
            offset.value = 0
            val value = _model.get(offset)
            values.add(value)
            _model.fbeShift(offset.value)
            size += offset.value
        }
        return size
    }

    // Set the vector as ArrayList
    fun set(values: ArrayList<StructSimple?>): Long
    {
        assert((_buffer.offset + fbeOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return 0

        write(fbeOffset, values.size.toUInt())

        var size: Long = 4
        _model.fbeOffset = fbeOffset + 4
        for (value in values)
        {
            val offset = _model.set(value)
            _model.fbeShift(offset)
            size += offset
        }
        return size
    }

    // Set the vector as LinkedList
    fun set(values: LinkedList<StructSimple?>): Long
    {
        assert((_buffer.offset + fbeOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return 0

        write(fbeOffset, values.size.toUInt())

        var size: Long = 4
        _model.fbeOffset = fbeOffset + 4
        for (value in values)
        {
            val offset = _model.set(value)
            _model.fbeShift(offset)
            size += offset
        }
        return size
    }

    // Set the vector as HashSet
    fun set(values: HashSet<StructSimple?>): Long
    {
        assert((_buffer.offset + fbeOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return 0

        write(fbeOffset, values.size.toUInt())

        var size: Long = 4
        _model.fbeOffset = fbeOffset + 4
        for (value in values)
        {
            val offset = _model.set(value)
            _model.fbeShift(offset)
            size += offset
        }
        return size
    }
}
