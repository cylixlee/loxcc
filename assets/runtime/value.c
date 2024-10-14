#include "value.h"
#include <stdio.h>  // for output
#include <string.h> // for string manipulation
#include <stdarg.h> // for vararg functions
#include "object.h" // for object definition
#include "gc.h"     // for allocation

/**
 * Print an object to `stdout`.
 *
 * This is a separate function from `LRT_PrintValue` because there're several object
 * types. Stuffing all the logic to `LRT_PrintValue` is not a good idea.
 */
static void LRT_PrintObject(LRT_Value value);
/**
 * Concatenate two strings, returns a new one as a Lox Value.
 */
static LRT_Value LRT_Concatenate(LRT_Value left, LRT_Value right);

/**
 * A convenient macro for operations between two number values.
 */
#define BINARY_OP(_As, _Operator, _Left, _Right) \
    if (!IS_NUMBER(_Left) || !IS_NUMBER(_Right)) \
    {                                            \
        LRT_Panic("operands must be numbers");   \
    }                                            \
    return _As(AS_NUMBER(_Left) _Operator AS_NUMBER(_Right));

/**
 * Addition between two numbers, or concatenation between two strings.
 */
LRT_Value LRT_Add(LRT_Value left, LRT_Value right)
{
    if (IS_STRING(left) && IS_STRING(right))
    {
        return LRT_Concatenate(left, right);
    }
    BINARY_OP(NUMBER, +, left, right);
}

// clang-format off
LRT_Value LRT_Subtract(LRT_Value left, LRT_Value right) { BINARY_OP(NUMBER, -, left, right); }
LRT_Value LRT_Multiply(LRT_Value left, LRT_Value right) { BINARY_OP(NUMBER, *, left, right); }
LRT_Value LRT_Divide(LRT_Value left, LRT_Value right)   { BINARY_OP(NUMBER, /, left, right); }

LRT_Value LRT_Equal(LRT_Value left, LRT_Value right)
{
    if (left.type != right.type)
    {
        return BOOLEAN(false);
    }
    
    switch (left.type)
    {
    case LVAL_Boolean: return BOOLEAN(AS_BOOLEAN(left) == AS_BOOLEAN(right));
    case LVAL_Nil:     return BOOLEAN(true);
    case LVAL_Number:  return BOOLEAN(AS_NUMBER(left) == AS_NUMBER(right));
    case LVAL_Object:  return BOOLEAN(AS_OBJECT(left) == AS_OBJECT(right));
    default:
        LRT_Panic("unreachable code (LOXCRT::Equal)");
    }
}

LRT_Value LRT_Greater(LRT_Value left, LRT_Value right) { BINARY_OP(BOOLEAN, >, left, right); }
LRT_Value LRT_Less(LRT_Value left, LRT_Value right)    { BINARY_OP(BOOLEAN, <, left, right); }

LRT_Value LRT_NotEqual(LRT_Value left, LRT_Value right)     { return LRT_Not(LRT_Equal(left, right)); }
LRT_Value LRT_LessEqual(LRT_Value left, LRT_Value right)    { return LRT_Not(LRT_Greater(left, right)); }
LRT_Value LRT_GreaterEqual(LRT_Value left, LRT_Value right) { return LRT_Not(LRT_Less(left, right)); }

// clang-format on

LRT_Value LRT_And(LRT_Value left, LRT_Value right)
{
    if (LRT_FalsinessOf(left))
    {
        return left;
    }
    return right;
}

LRT_Value LRT_Or(LRT_Value left, LRT_Value right)
{
    if (!LRT_FalsinessOf(left))
    {
        return left;
    }
    return right;
}

LRT_Value LRT_Negate(LRT_Value value)
{
    if (!IS_NUMBER(value))
    {
        LRT_Panic("operand must be a number");
    }
    AS_NUMBER(value) = -AS_NUMBER(value);
    return value;
}

LRT_Value LRT_Not(LRT_Value value) { return BOOLEAN(LRT_FalsinessOf(value)); }

LRT_Value LRT_Call(LRT_Value callee, size_t arity, ...)
{
    // type check and prepare fn.
    if (!IS_FUNCTION(callee))
    {
        LRT_Panic("attempt to invoke a non-callable object");
    }
    LRT_Fn fn = AS_FN(callee);

    // accessing vararg list, and call fn.
    va_list varlist;
    va_start(varlist, arity);
    LRT_Value returnValue = fn(arity, varlist); // record return value. fn WILL return.
    va_end(varlist);

    return returnValue;
}

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
    printf("\n");
    // clang-format on
}

static void LRT_PrintObject(LRT_Value value)
{
    switch (TYPEOF(value))
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
    return OBJECT(result);
}