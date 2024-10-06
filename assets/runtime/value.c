#include "value.h"
#include <stdio.h>
#include <string.h>
#include "object.h"
#include "gc.h"

static void LRT_PrintObject(LRT_Value value);
static LRT_Value LRT_Concatenate(LRT_Value left, LRT_Value right);

#define BINARY_OP(_As, _Operator, _Left, _Right) \
    if (!IS_NUMBER(_Left) || !IS_NUMBER(_Right)) \
    {                                            \
        LRT_Panic("operands must be numbers");   \
    }                                            \
    return _As(AS_NUMBER(_Left) _Operator AS_NUMBER(_Right));

LRT_Value LRT_Add(LRT_Value left, LRT_Value right)
{
    if (IS_STRING(left) && IS_STRING(right))
    {
        return LRT_Concatenate(left, right);
    }
    BINARY_OP(NUMBER_VAL, +, left, right);
}

// clang-format off
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
    case LVAL_Object:
        LRT_StringObject *leftString  = AS_STRING(left);
        LRT_StringObject *rightString = AS_STRING(right);
        return BOOLEAN_VAL(
            leftString->length == rightString->length &&
            memcmp(
                leftString->chars, 
                rightString->chars, 
                leftString->length
            ) == 0
        );
    default:
        LRT_Panic("unreachable code (LOXCRT::Equal)");
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
    case LVAL_Object: LRT_PrintObject(value);         break;
    default:
        LRT_Panic("unreachable code (LOXCRT::Print)");
    }
    // clang-format on
}

static void LRT_PrintObject(LRT_Value value)
{
    switch (OBJ_TYPE(value))
    {
    case LOBJ_String:
        printf("%s", AS_CSTR(value));
        break;
    default:
        LRT_Panic("unreachable code (LOXCRT::PrintObject)");
    }
}

static LRT_Value LRT_Concatenate(LRT_Value left, LRT_Value right)
{
    LRT_StringObject *leftString = AS_STRING(left);
    LRT_StringObject *rightString = AS_STRING(right);

    size_t length = leftString->length + rightString->length;
    char *chars = ALLOCATE(char, length + 1);
    memcpy(chars, leftString->chars, leftString->length);
    memcpy(chars + leftString->length, rightString->chars, rightString->length);
    chars[length] = '\0';

    LRT_StringObject *result = LRT_TakeString(chars, length);
    return OBJECT_VAL(result);
}