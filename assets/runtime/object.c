#include "object.h"
#include <string.h> // for string manipulation
#include "gc.h"     // for unified allocation

LRT_StringObject *LRT_NewString(const char *chars, size_t length)
{
    LRT_StringObject *string = ALLOCATE_OBJ(LRT_StringObject, LOBJ_String);
    string->length = length;
    string->chars = ALLOCATE(char, length + 1);
    strcpy(string->chars, chars);
    string->chars[length] = '\0';
    return string;
}

LRT_StringObject *LRT_TakeString(char *chars, size_t length)
{
    LRT_StringObject *string = ALLOCATE_OBJ(LRT_StringObject, LOBJ_String);
    string->chars = chars;
    string->length = length;
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