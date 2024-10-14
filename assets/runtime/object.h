#ifndef LOXCRT_OBJECT_H
#define LOXCRT_OBJECT_H

#include "prelude.h"
#include "value.h"  // for Object-Value type check.
#include <stdarg.h> // for vararg functions.

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
        LOBJ_Function,
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
        struct LRT_Object *next; // a field for intrusive linked list
    } LRT_Object;

    /**
     * Lox Strings.
     *
     * Lox Strings are immutable objects, which contains a cluster of chars and the length
     * of that. Trailing `\0` is preserved in order to convert to C-style strings easily.
     */
    typedef struct
    {
        LRT_Object intrinsic;
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
     * Function pointer compatible with Lox calls.
     *
     * However, we could not use a signature that has definite parameters to represent all
     * Lox functions, which have different arities. The C varargs is adopted so, with the
     * first argument as the arity.
     */
    typedef LRT_Value (*LRT_Fn)(size_t arity, va_list args);

    /**
     * The Lox function.
     *
     * Lox, like other dynamic-typed languages, adopts function as first-class values.
     * That is, we need to represent a function as a LRT_Value.
     *
     * The underlying code is referenced by a function pointer in C. For recording the
     * callstack (for panic tracing), we'll need a String object to store the function
     * name.
     */
    typedef struct
    {
        LRT_Object intrinsic;
        LRT_StringObject *name;
        LRT_Fn fn;
    } LRT_FunctionObject;

    // Create a function object from its name and a function pointer.
    LRT_FunctionObject *LRT_NewFunction(LRT_StringObject *name, LRT_Fn fn);

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
// Convenient macro for checking whether a Value is a Function.
#define IS_FUNCTION(_Value) isinstance(_Value, LOBJ_Function)

// Convenient macro for converting a Value to a String.
#define AS_STRING(_Value) ((LRT_StringObject *)AS_OBJECT(_Value))
// Convenient macro for converting a Value to a C-style string.
#define AS_CSTR(_Value) (AS_STRING(_Value)->chars)
// Convenient macro for converting a Value to a Function.
#define AS_FUNCTION(_Value) ((LRT_FunctionObject *)AS_OBJECT(_Value))
// Convenient macro for converting a Value to a C-callable function.
#define AS_FN(_Value) (AS_FUNCTION(_Value)->fn)

#ifdef __cplusplus
}
#endif

#endif