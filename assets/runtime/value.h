#ifndef LOXCRT_VALUE_H
#define LOXCRT_VALUE_H

#include "prelude.h"

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

    struct LRT_Object;    // external, in "object.h"
    enum   LRT_ValueType;
    struct LRT_Value;

    typedef struct LRT_Object    LRT_Object;
    typedef enum   LRT_ValueType LRT_ValueType;
    typedef struct LRT_Value     LRT_Value;

    // clang-format on

    enum LRT_ValueType
    {
        LVAL_Boolean,
        LVAL_Nil,
        LVAL_Number,
        LVAL_Object,
    };

    struct LRT_Value
    {
        LRT_ValueType type;
        union
        {
            bool boolean;
            double number;
            LRT_Object *object;
        } as;
    };

    /**
     * Lox Value Utilities
     *
     * Compiling Lox values to C is a difficult work. Without language-level support for some
     * features (e.g. implicit type conversion, operator overloading, etc.), we have to seek
     * for macros and functions' help.
     *
     * These macros & functions below are about value initialization, type-check, conversion
     * and other operations. By simply compiling Lox expressions to macro-wrapped C
     * expressions, we can simplify LoxCC a lot.
     */

    // clang-format off

#define BOOLEAN_VAL(_Value)   ((LRT_Value){LVAL_Boolean, {.boolean = (_Value)}})
#define NIL_VAL               ((LRT_Value){LVAL_Nil, {.number = 0}})
#define NUMBER_VAL(_Value)    ((LRT_Value){LVAL_Number, {.number = (_Value)}})
#define OBJECT_VAL(_Value)    ((LRT_Value){LVAL_Object, {.object = ((LRT_Object *)(_Value))}})

#define IS_BOOLEAN(_Value) ((_Value).type == LVAL_Boolean)
#define IS_NIL(_Value)     ((_Value).type == LVAL_Nil)
#define IS_NUMBER(_Value)  ((_Value).type == LVAL_Number)
#define IS_OBJECT(_Value)  ((_Value).type == LVAL_Object)

#define AS_BOOLEAN(_Value) ((_Value).as.boolean)
#define AS_NUMBER(_Value)  ((_Value).as.number)
#define AS_OBJECT(_Value)  ((_Value).as.object)

    // clang-format on

    /**
     * Binary operations
     */

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

    /**
     * Unary operations
     */

    LRT_Value LRT_Negate(LRT_Value);
    LRT_Value LRT_Not(LRT_Value);

    /**
     * Builtin functionalities
     */

    bool LRT_FalsinessOf(LRT_Value);
    void LRT_Print(LRT_Value);

// tells C++ compiler to treat the code as C source.
#ifdef __cplusplus
}
#endif

#endif