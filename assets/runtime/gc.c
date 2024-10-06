#include "gc.h"
#include <stdlib.h>
#include <string.h>

struct
{
    LRT_Object *objects;
} GC;

void LRT_InitializeGC()
{
    GC.objects = NULL;
}

void LRT_FinalizeGC()
{
    LRT_Object *object = GC.objects;
    while (object != NULL)
    {
        LRT_Object *next = object->next;
        LRT_FinalizeObject(object);
        object = next;
    }
}

LRT_Object *LRT_AllocateObject(size_t size, LRT_ObjectType type)
{
    LRT_Object *object = LRT_Reallocate(NULL, 0, size);
    object->type = type;
    object->next = GC.objects;
    GC.objects = object;
    return object;
}

void *LRT_Reallocate(void *pointer, size_t oldSize, size_t newSize)
{
    // nothing happens if both size are 0
    if (oldSize == 0 && newSize == 0)
    {
        return NULL;
    }

    // allocate a zero-initialized new block if oldSize is 0.
    if (oldSize == 0)
    {
        pointer = malloc(newSize);
        if (pointer == NULL)
        {
            LRT_Panic("allocation failure; may be out of memory");
        }
        memset(pointer, 0, newSize);
        return pointer;
    }

    // free block if newSize is 0
    if (newSize == 0)
    {
        free(pointer);
        return NULL;
    }

    // the system re-alloc.
    pointer = realloc(pointer, newSize);
    if (pointer == NULL)
    {
        LRT_Panic("reallocation failure; may be out of memory");
    }
    return pointer;
}