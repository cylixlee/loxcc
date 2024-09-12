#ifndef LOXCRT_VALUE_H
#define LOXCRT_VALUE_H

#include "prelude.h"

typedef enum
{
    LBoolean,
    LNil,
    LNumber,
} LRT_ValueType;

typedef struct
{
    LRT_ValueType type;
    union
    {
        bool boolean;
        double number;
    } as;
} LRT_Value;

// clang-format off

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

#define BOOLEAN_VAL(_Value)   ((LRT_Value){LBoolean, {.boolean = (_Value)}})
#define NIL_VAL               ((LRT_Value){LNil, {.number = 0}})
#define NUMBER_VAL(_Value)    ((LRT_Value){LNumber, {.number = (_Value)}})

#define IS_BOOLEAN(_Value) ((_Value).type == LBoolean)
#define IS_NIL(_Value)     ((_Value).type == LNil)
#define IS_NUMBER(_Value)  ((_Value).type == LNumber)

#define AS_BOOLEAN(_Value)   ((_Value).as.boolean)
#define AS_NUMBER(_Value)    ((_Value).as.number)

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

#endif