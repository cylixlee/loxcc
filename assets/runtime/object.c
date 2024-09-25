#include "object.h"
#include <string.h>
#include "gc.h"

// the underlying function to create objects.
//
// this is necessary to add newly created objects to VM's heap, to perform GC on.
static inline LRT_Object *LRT_AllocateObject(size_t size, LRT_ObjectType type);

// convenient macro for allocating objects.
#define ALLOCATE_OBJ(_Type, _ObjectType) ((_Type *)LRT_AllocateObject(sizeof(_Type), _ObjectType))

LRT_StringObject *LRT_NewString(const char *chars, size_t length)
{
    LRT_StringObject *string = ALLOCATE_OBJ(LRT_StringObject, LOBJ_String);
    string->length = length;
    string->chars = strcpy(string->chars, chars);
    return string;
}

static inline LRT_Object *LRT_AllocateObject(size_t size, LRT_ObjectType type)
{
    LRT_Object *object = LRT_Reallocate(NULL, 0, size);
    object->type = type;
    return object;
}