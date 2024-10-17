#ifndef LOXCRT_GC_H
#define LOXCRT_GC_H

#include "prelude.h"
#include "object.h" // for Object definitions

#ifdef __cplusplus
extern "C"
{
#endif

    /**
     * Initialize GC.
     *
     * The Garbage Collector is, in this implementation, a global variable with some
     * internal state recorded. Since C does not support constructors, we should manually
     * initialize it by calling this function.
     */
    void LRT_InitializeGC();

    /**
     * Finalize GC.
     *
     * GC has to do some clean-up operations when the program exits. LOXCRT ensures that
     * all memory allocated will be freed when the GC exits. Otherwise, an error is
     * reported and the program is exited with a non-zero status code.
     */
    void LRT_FinalizeGC();

    /**
     * Invoke GC.
     *
     * Manually or automatically called to run the garbage collection procedure.
     */
    void LRT_GarbageCollect();

    /**
     * String interning.
     *
     * Lox interns every string, so that strings that are literally equal is referential
     * equal. This saves a lot of time comparing strings, which is very important for the
     * Table.
     *
     * By the way, this feature is implemented by using a Table.
     */
    void LRT_GCInternString(LRT_StringObject *string);
    // Find whether GC has interned this string.
    LRT_StringObject *LRT_GCFindInterned(const char *chars, size_t length, uint32_t hash);

    /**
     * Allocate an Object, with internal status being set.
     *
     * LOXCRT has to insert some internal status (type information, GC mark, etc.) to make
     * GC work correctly. This should be the unified way to allocate an object.
     */
    LRT_Object *LRT_AllocateObject(size_t size, LRT_ObjectType type);

    /**
     * The universal allocation function of LOXCRT.
     *
     * Unifying the allocation method is convenient for collecting statistics and tracing
     * objects, especially for GC.
     *
     * The table below describes the behavior of this function:
     *
     * `oldSize` | `newSize`              | Behavior
     * ----------|------------------------|----------------------------
     * 0         | 0                      | Returns NULL.
     * 0         | Non-zero               | Allocate new block.
     * Non-zero  | 0                      | Free allocation.
     * Non-zero  | Smaller than `oldSize` | Shrink existing allocation.
     * Non-zero  | Larger than `oldSize`  | Grow existing allocation.
     * Non-zero  | Equals to `oldSize`    | Returns the given pointer.
     *
     * If a new block of memory is allocated, it is ensured to be initialized with
     * zero-value.
     *
     * Note that the parameter `pointer` does not affect the behavior. More specifically,
     * it's only useful in the "growing" behaviour: the allocator tries to extend the
     * memory in-place, but may fail and returns a different address.
     */
    void *LRT_Reallocate(void *pointer, size_t oldSize, size_t newSize);

// Convenient macro for allocating a block of memory, using `LRT_Reallocate`.
#define ALLOCATE(_Type, _Count) ((_Type *)(LRT_Reallocate(NULL, 0, sizeof(_Type) * _Count)))
// Convenient macro for freeing a block of memory, using `LRT_Reallocate`.
#define FREE(_Ptr, _Type, _Count) ((_Type *)(LRT_Reallocate(_Ptr, sizeof(_Type) * _Count, 0)))

/**
 * Convenient macro for allocating an Object, using `LRT_AllocateObject`.
 *
 * Note that there's no `FREE_OBJ` macro because GC calls finalizers of objects
 * internally. Other parts of source does not need to care about that.
 */
#define ALLOCATE_OBJ(_Type, _ObjectType) ((_Type *)LRT_AllocateObject(sizeof(_Type), _ObjectType))

#ifdef __cplusplus
}
#endif

#endif