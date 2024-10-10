#include "object.h"
#include <string.h> // for string manipulation
#include "gc.h"     // for unified allocation

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
    LRT_StringObject *string = ALLOCATE_OBJ(LRT_StringObject, LOBJ_String);
    string->chars = chars;
    string->length = length;
    string->hash = LRT_HashString(chars, length);
    return string;
}

void LRT_FinalizeObject(LRT_Object *object)
{
    switch (object->type)
    {
    case LOBJ_String:
        LRT_StringObject *strobj = (LRT_StringObject *)object;
        FREE(strobj->chars, char, strobj->length + 1);
        FREE(strobj, LRT_StringObject, 1);
        break;
    default:
        LRT_Panic("unreachable code (LOXCRT::FinalizeObject)");
    }
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