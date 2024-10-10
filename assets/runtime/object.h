#ifndef LOXCRT_OBJECT_H
#define LOXCRT_OBJECT_H

#include "prelude.h"
#include "value.h" // for Object-Value type check.

#ifdef __cplusplus
extern "C"
{
#endif

    /**
     * All possible types of Objects.
     *
     * From a runtime's perspective, categories of objects are exhaustible. For example,
     * user-defined classes are of type `Class`, and instances of them are of type
     * `Instance`.
     */
    typedef enum
    {
        LOBJ_String,
    } LRT_ObjectType;

    /**
     * The basic definition of objects.
     *
     * Specifically, in C, this struct is placed as a certain object type's first field,
     * in order to make the pointer types of which mutual convertible. This is, to some
     * extent, a poor guy's polymorphism.
     */
    typedef struct LRT_Object
    {
        LRT_ObjectType type;     // type information
        struct LRT_Object *next; // a field for **intrusive** linked list
    } LRT_Object;

    /**
     * Lox Strings.
     *
     * Lox Strings are immutable objects, which contains a cluster of chars and the length
     * of that. Trailing `\0` is preserved in order to convert to C-style strings easily.
     */
    typedef struct
    {
        LRT_Object meta;
        size_t length;
        char *chars;
        uint32_t hash;
    } LRT_StringObject;

    // Create a Lox String from a C-style string literal.
    LRT_StringObject *LRT_NewString(const char *literal, size_t length);
    /**
     * Create a Lox String from a pointer to chars and the length.
     *
     * Note that the pointer is then **taken** by the created String object, which means
     * the pointer should not be freed somewhere else. The chars should be `\0` ended, and
     * the length should be correct.
     */
    LRT_StringObject *LRT_TakeString(char *chars, size_t length);

    /**
     * The unified function to finalize an object.
     *
     * This function should be called **only** by GC, as other parts of code should not
     * release memory manually. This function promises to free the object correctly
     * according to its actual definition.
     */
    void LRT_FinalizeObject(LRT_Object *object);

// Convenient way to get a Object Value's ObjectType.
#define TYPEOF(_Value) (AS_OBJECT(_Value)->type)

    /**
     * Check whether a value is an instance of a certain ObjectType.
     *
     * It only returns true when the value is an object, and it matches the given
     * ObjectType. It's not a macro because the value may be an expression, which could be
     * evaluated twice in a macro.
     *
     * It's `static inline`ed to avoid conflict definition.
     */
    static inline bool isinstance(LRT_Value value, LRT_ObjectType type)
    {
        return IS_OBJECT(value) && TYPEOF(value) == type;
    }

// Convenient macro for checking whether a Value is a String.
#define IS_STRING(_Value) isinstance(_Value, LOBJ_String)

// Convenient macro for convert a Value to a String.
#define AS_STRING(_Value) ((LRT_StringObject *)AS_OBJECT(_Value))
// Convenient macro for convert a Value to a C-style string.
#define AS_CSTR(_Value) (AS_STRING(_Value)->chars)

#ifdef __cplusplus
}
#endif

#endif