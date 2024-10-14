#ifndef LOXCRT_VALUE_H
#define LOXCRT_VALUE_H

#include "prelude.h"

#ifdef __cplusplus
extern "C"
{
#endif

    /**
     * All possible types of Lox Values.
     */
    typedef enum
    {
        LVAL_Boolean,
        LVAL_Nil,
        LVAL_Number,
        LVAL_Object,
    } LRT_ValueType;

    struct LRT_Object;                    // external, in "object.h"
    typedef struct LRT_Object LRT_Object; // corresponding typedef

    /**
     * Lox Values.
     *
     * Lox is a dynamically typed language, so the Value struct should contain all
     * possible situations of a value. Tagged unions, as a common practice, is how the
     * struct is defined.
     *
     * Values are designed to be small enough to be frequently passed. Thus, object types,
     * which are often much bigger than PODs, are heap-allocated and the Value structs may
     * keep pointers to them.
     *
     * NOTE: All helper macros about type checking and conversion are of `Value`
     * parameters. We don't use `Value.as.object` as any parameter of macros or functions.
     * Just use `Value`s.
     */
    typedef struct
    {
        LRT_ValueType type;
        union
        {
            bool boolean;
            double number;
            LRT_Object *object;
        } as;
    } LRT_Value;

    /**
     * Lox Value Utilities
     *
     * Compiling Lox values to C is a difficult work. Without language-level support for
     * some features (e.g. implicit type conversion, operator overloading, etc.), we have
     * to seek for macros and functions' help.
     *
     * These macros & functions below are about value initialization, type-check,
     * conversion and other operations. By simply compiling Lox expressions to
     * macro-wrapped C expressions, we can simplify LoxCC a lot.
     */

    // clang-format off

#define BOOLEAN(_Value) ((LRT_Value){LVAL_Boolean, {.boolean = (_Value)}})
#define NIL             ((LRT_Value){LVAL_Nil, {.number = 0}})
#define NUMBER(_Value)  ((LRT_Value){LVAL_Number, {.number = (_Value)}})
#define OBJECT(_Value)  ((LRT_Value){LVAL_Object, {.object = ((LRT_Object *)(_Value))}})

#define IS_BOOLEAN(_Value) ((_Value).type == LVAL_Boolean)
#define IS_NIL(_Value)     ((_Value).type == LVAL_Nil)
#define IS_NUMBER(_Value)  ((_Value).type == LVAL_Number)
#define IS_OBJECT(_Value)  ((_Value).type == LVAL_Object)

#define AS_BOOLEAN(_Value) ((_Value).as.boolean)
#define AS_NUMBER(_Value)  ((_Value).as.number)
#define AS_OBJECT(_Value)  ((_Value).as.object)

    // clang-format on

    // === Binary operations ===

    LRT_Value LRT_Add(LRT_Value, LRT_Value);
    LRT_Value LRT_Subtract(LRT_Value, LRT_Value);
    LRT_Value LRT_Multiply(LRT_Value, LRT_Value);
    LRT_Value LRT_Divide(LRT_Value, LRT_Value);

    LRT_Value LRT_Equal(LRT_Value, LRT_Value);
    LRT_Value LRT_Greater(LRT_Value, LRT_Value);
    LRT_Value LRT_Less(LRT_Value, LRT_Value);

    LRT_Value LRT_NotEqual(LRT_Value, LRT_Value);
    LRT_Value LRT_LessEqual(LRT_Value, LRT_Value);
    LRT_Value LRT_GreaterEqual(LRT_Value, LRT_Value);

    LRT_Value LRT_And(LRT_Value, LRT_Value);
    LRT_Value LRT_Or(LRT_Value, LRT_Value);

    // === Unary operations ===

    LRT_Value LRT_Negate(LRT_Value);
    LRT_Value LRT_Not(LRT_Value);

    // === Invocation operations ===
    LRT_Value LRT_Call(LRT_Value callee, size_t arity, ...);

    // === Builtin functionalities ===

    /**
     * Check whether the value can be evaluated as a `false`.
     *
     * This is very common in dynamically typed languages. For example, a `nil` can be
     * implicitly converted to false using `!` operator.
     *
     * As the language defines, only `nil` and `false` can be converted to false; other
     * values are implicitly `true`.
     */
    bool LRT_FalsinessOf(LRT_Value);

    /**
     * Print a Lox Value to `stdout`, no newline is printed.
     */
    void LRT_Print(LRT_Value);

#ifdef __cplusplus
}
#endif

#endif