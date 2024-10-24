#include "gc.h"
#include <stdlib.h> // for memory allocation & exiting program
#include <string.h> // for `memset`
#include <stdio.h>  // for (error) output
#include "table.h"  // for String interning

/**
 * The Garbage Collector.
 *
 * This is a private struct variable because other parts of LOXCRT should not rely on the
 * internal implementation of GC. They should call exposed interfaces instead.
 */
struct
{
    LRT_Table strings;   // string pool, all strings are interned in Lox
    LRT_Object *objects; // linked list of allocated objects
    size_t allocated;
    size_t freed;
} GC;

void LRT_InitializeGC()
{
    LRT_InitializeTable(&GC.strings);
    GC.objects = NULL;
    GC.allocated = 0;
    GC.freed = 0;
}

void LRT_FinalizeGC()
{
    LRT_FinalizeTable(&GC.strings);

    // clear heap
    //
    // this must be the final step, since other fields of GC may heap-allocate some
    // objects.
    LRT_Object *object = GC.objects;
    while (object != NULL)
    {
        LRT_Object *next = object->next;
        LRT_FinalizeObject(object);
        object = next;
    }

#ifdef GC_TRACE
    printf("=== GC Total Allocated: %llu\n", GC.allocated);
    printf("=== GC Total Freed: %llu\n", GC.freed);
#endif

    if (GC.allocated != GC.freed) // simple leak check
    {
        fprintf(stderr, "internal error: GC unclear");
        exit(EXIT_FAILURE);
    }
}

void LRT_GCInternString(LRT_StringObject *string)
{
    LRT_TableSet(&GC.strings, string, NIL);
}

LRT_StringObject *LRT_GCFindInterned(const char *chars, size_t length, uint32_t hash)
{
    return LRT_TableContainsKey(&GC.strings, chars, length, hash);
}

LRT_Object *LRT_AllocateObject(size_t size, LRT_ObjectType type)
{
    LRT_Object *object = LRT_Reallocate(NULL, 0, size);
    object->type = type;       // maintain type information.
    object->next = GC.objects; // add object to the linked list.
    GC.objects = object;
    return object;
}

void *LRT_Reallocate(void *pointer, size_t oldSize, size_t newSize)
{
    // no allocation if both size are identical.
    if (oldSize == newSize)
    {
        if (oldSize == 0)
        {
            return NULL;
        }
        return pointer;
    }

    // allocate a zero-initialized new block if oldSize is 0.
    if (oldSize == 0)
    {
        pointer = malloc(newSize);
        if (pointer == NULL)
        {
            LRT_Panic("allocation failure; may be out of memory");
        }
        GC.allocated += newSize;     // record size of allocated memory
        memset(pointer, 0, newSize); // zero initialize
        return pointer;
    }

    // free block if newSize is 0
    if (newSize == 0)
    {
        if (pointer == NULL)
        {
            LRT_Panic("trying to free NULL pointer with non-zero size");
        }
        GC.freed += oldSize; // record size of freed memory
        free(pointer);
        return NULL;
    }

    // the system re-alloc.
    pointer = realloc(pointer, newSize);
    if (pointer == NULL)
    {
        LRT_Panic("reallocation failure; may be out of memory");
    }

    if (newSize > oldSize) // record size of `realloc`ed memory
    {
        // existing allocation is growed
        GC.allocated += newSize - oldSize;
    }
    else
    {
        // existing allocation is shrinked
        GC.freed += oldSize - newSize;
    }
    return pointer;
}