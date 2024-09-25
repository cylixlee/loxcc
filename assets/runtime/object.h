#ifndef LOXCRT_OBJECT_H
#define LOXCRT_OBJECT_H

#include "prelude.h"
#include "value.h"

// tells C++ compiler to treat the code as C source.
#ifdef __cplusplus
extern "C"
{
#endif

    /**
     * Type declarations and typedefs.
     *
     * Instead of writing typedefs everywhere a struct is defined, writing them together
     * at the beginning of a header file is more tidy and good for circular referencing.
     */

    // clang-format off

    enum   LRT_ObjectType;
    struct LRT_Object;
    struct LRT_StringObject;

    typedef enum   LRT_ObjectType   LRT_ObjectType;
    typedef struct LRT_Object       LRT_Object;
    typedef struct LRT_StringObject LRT_StringObject;

    // clang-format on

    enum LRT_ObjectType
    {
        LOBJ_String,
    };

    struct LRT_Object
    {
        LRT_ObjectType type;
    };

    struct LRT_StringObject
    {
        LRT_Object meta;
        size_t length;
        char *chars;
    };

    LRT_StringObject *LRT_NewString(const char *, size_t length);

    /**
     * Utilities for objects' type check.
     */

    // clang-format off

#define OBJ_TYPE(_Value) (AS_OBJECT(_Value)->type)

    inline bool isinstance(LRT_Value value, LRT_ObjectType type)
    {
        return IS_OBJECT(value) && OBJ_TYPE(value) == type;
    }

#define IS_STRING(_Value) isinstance(_Value, LOBJ_String)

#define AS_STRING(_Value) ((LRT_StringObject*)AS_OBJECT(_Value))
#define AS_CSTR(_Value)   (AS_STRING(_Value)->chars)

    // clang-format on

// tells C++ compiler to treat the code as C source.
#ifdef __cplusplus
}
#endif

#endif