// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import fbe.*;
import enums.*;

import com.google.gson.*;

// Fast Binary Encoding enums JSON class
public final class Json
{
    private static final Gson _engine;

    // Get the JSON engine
    public static Gson getEngine() { return _engine; }

    static
    {
        _engine = register(new GsonBuilder()).create();
    }

    private Json() {}

    public static GsonBuilder register(GsonBuilder builder)
    {
        fbe.Json.register(builder);
        builder.registerTypeAdapter(EnumByte.class, new EnumByteJson());
        builder.registerTypeAdapter(EnumChar.class, new EnumCharJson());
        builder.registerTypeAdapter(EnumWChar.class, new EnumWCharJson());
        builder.registerTypeAdapter(EnumInt8.class, new EnumInt8Json());
        builder.registerTypeAdapter(EnumUInt8.class, new EnumUInt8Json());
        builder.registerTypeAdapter(EnumInt16.class, new EnumInt16Json());
        builder.registerTypeAdapter(EnumUInt16.class, new EnumUInt16Json());
        builder.registerTypeAdapter(EnumInt32.class, new EnumInt32Json());
        builder.registerTypeAdapter(EnumUInt32.class, new EnumUInt32Json());
        builder.registerTypeAdapter(EnumInt64.class, new EnumInt64Json());
        builder.registerTypeAdapter(EnumUInt64.class, new EnumUInt64Json());
        return builder;
    }
}
