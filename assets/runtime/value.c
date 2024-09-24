#include "value.h"
#include <stdio.h>

#define BINARY_OP(_Ty, _Operator, _Left, _Right) \
    if (!IS_NUMBER(_Left) || !IS_NUMBER(_Right)) \
    {                                            \
        LRT_Panic("operands must be numbers");   \
    }                                            \
    return _Ty(AS_NUMBER(_Left) _Operator AS_NUMBER(_Right));

// clang-format off

LRT_Value LRT_Add(LRT_Value left, LRT_Value right)      { BINARY_OP(NUMBER_VAL, +, left, right); }
LRT_Value LRT_Subtract(LRT_Value left, LRT_Value right) { BINARY_OP(NUMBER_VAL, -, left, right); }
LRT_Value LRT_Multiply(LRT_Value left, LRT_Value right) { BINARY_OP(NUMBER_VAL, *, left, right); }
LRT_Value LRT_Divide(LRT_Value left, LRT_Value right)   { BINARY_OP(NUMBER_VAL, /, left, right); }

LRT_Value LRT_Equal(LRT_Value left, LRT_Value right)
{
    if (left.type != right.type)
    {
        return BOOLEAN_VAL(false);
    }
    
    switch (left.type)
    {
    case LVAL_Boolean: return BOOLEAN_VAL(AS_BOOLEAN(left) == AS_BOOLEAN(right));
    case LVAL_Nil:     return BOOLEAN_VAL(true);
    case LVAL_Number:  return BOOLEAN_VAL(AS_NUMBER(left) == AS_NUMBER(right));
    default:
        LRT_Panic("unreachable LRT_Equal");
    }
}

LRT_Value LRT_Greater(LRT_Value left, LRT_Value right) { BINARY_OP(BOOLEAN_VAL, >, left, right); }
LRT_Value LRT_Less(LRT_Value left, LRT_Value right)    { BINARY_OP(BOOLEAN_VAL, <, left, right); }

LRT_Value LRT_NotEqual(LRT_Value left, LRT_Value right)     { return LRT_Not(LRT_Equal(left, right)); }
LRT_Value LRT_LessEqual(LRT_Value left, LRT_Value right)    { return LRT_Not(LRT_Greater(left, right)); }
LRT_Value LRT_GreaterEqual(LRT_Value left, LRT_Value right) { return LRT_Not(LRT_Less(left, right)); }

// clang-format on

LRT_Value LRT_Negate(LRT_Value value)
{
    if (!IS_NUMBER(value))
    {
        LRT_Panic("operand must be a number");
    }
    AS_NUMBER(value) = -AS_NUMBER(value);
    return value;
}

LRT_Value LRT_Not(LRT_Value value) { return BOOLEAN_VAL(LRT_FalsinessOf(value)); }

bool LRT_FalsinessOf(LRT_Value value)
{
    return IS_NIL(value) || (IS_BOOLEAN(value) && !AS_BOOLEAN(value));
}

void LRT_Print(LRT_Value value)
{
    // clang-format off
    switch (value.type)
    {
    case LVAL_Boolean:
        printf("%s", AS_BOOLEAN(value) ? "true" : "false");
        break;

    case LVAL_Nil:    printf("nil");                  break;
    case LVAL_Number: printf("%g", AS_NUMBER(value)); break;
    }
    // clang-format on
}