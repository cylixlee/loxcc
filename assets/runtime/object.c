#include "object.h"
#include <string.h> // for string manipulation
#include "gc.h"     // for unified allocation

#ifdef GC_TRACE
#include <stdio.h> // for Object finalization tracing
#endif

/**
 * The universal function to allocate a string object.
 *
 * NewString and TakeString call this eventually. It's very convenient for string
 * interning.
 */
static LRT_StringObject *LRT_AllocateString(char *chars, size_t length, uint32_t hash);
/**
 * FNV-1a hash algorithm.
 *
 * Actually I don't know how this works; the code is just copied from clox.
 */
static uint32_t LRT_HashString(const char *key, size_t length);

LRT_StringObject *LRT_NewString(const char *chars, size_t length)
{
    char *ownedChars = ALLOCATE(char, length + 1);
    strcpy(ownedChars, chars);
    ownedChars[length] = '\0';
    return LRT_TakeString(ownedChars, length);
}

LRT_StringObject *LRT_TakeString(char *chars, size_t length)
{
    uint32_t hash = LRT_HashString(chars, length);
    return LRT_AllocateString(chars, length, hash);
}

LRT_FunctionObject *LRT_NewFunction(LRT_StringObject *name, LRT_Fn fn)
{
    LRT_FunctionObject *object = ALLOCATE_OBJ(LRT_FunctionObject, LOBJ_Function);
    object->name = name;
    object->fn = fn;
}

void LRT_FinalizeObject(LRT_Object *object)
{
#ifdef GC_TRACE
    printf("=== Finalize object %p\n", object);
#endif
    switch (object->type)
    {
    case LOBJ_String:
        LRT_StringObject *strobj = (LRT_StringObject *)object;
        FREE(strobj->chars, char, strobj->length + 1);
        FREE(strobj, LRT_StringObject, 1);
        break;
    case LOBJ_Function:
        FREE(object, LRT_FunctionObject, 1); // function objects contain no owned allocations.
        break;
    default:
        LRT_Panic("unreachable code (LOXCRT::FinalizeObject)");
    }
}

static LRT_StringObject *LRT_AllocateString(char *chars, size_t length, uint32_t hash)
{
    // use the interned one if any
    LRT_StringObject *interned = LRT_GCFindInterned(chars, length, hash);
    if (interned != NULL)
    {
        FREE(chars, char, length + 1);
        return interned;
    }

    // otherwise, allocate the string and intern it.
    LRT_StringObject *string = ALLOCATE_OBJ(LRT_StringObject, LOBJ_String);
    string->chars = chars;
    string->length = length;
    string->hash = hash;
    LRT_GCInternString(string);
    return string;
}

static uint32_t LRT_HashString(const char *key, size_t length)
{
    uint32_t hash = 2166136261u;
    for (int i = 0; i < length; i++)
    {
        hash ^= (uint8_t)key[i];
        hash *= 16777619;
    }
    return hash;
}