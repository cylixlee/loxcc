#ifndef LOXCRT_GC_H
#define LOXCRT_GC_H

#include "prelude.h"
#include "object.h"

// tells C++ compiler to treat the code as C source.
#ifdef __cplusplus
extern "C"
{
#endif

    /**
     * Type declarations and typedefs.
     *
     * Instead of writing typedefs everywhere a struct is defined, writing them together
     * at the beginning of a header file is more tidy and good for circular referencing.
     */

    void LRT_InitializeGC();
    void LRT_FinalizeGC();

    LRT_Object *LRT_AllocateObject(size_t size, LRT_ObjectType type);

    /**
     * The universal allocation function of LOXCRT.
     *
     * The table below describes the behavior of this function:
     *
     * `oldSize` | `newSize`              | Behavior
     * ----------|------------------------|----------------------------
     * 0         | 0                      | Nothing happens.
     * 0         | Non-zero               | Allocate new block.
     * Non-zero  | 0                      | Free allocation.
     * Non-zero  | Smaller than `oldSize` | Shrink existing allocation.
     * Non-zero  | Larger than `oldSize`  | Grow existing allocation.
     *
     * If a new block of memory is allocated, it is initialized with zero-value.
     *
     * Note that the parameter `pointer` does not affect the behavior. More specifically,
     * it's only useful in the "growing" behaviour: the allocator tries to extend the
     * memory in-place, but may fail and returns a different address.
     *
     * Unifying the allocation method is convenient for collecting statistics and tracing
     * objects, especially for GC.
     */
    void *LRT_Reallocate(void *pointer, size_t oldSize, size_t newSize);

// Convenient macro for allocation using `LRT_Reallocate`.
#define ALLOCATE(_Type, _Count) ((_Type *)(LRT_Reallocate(NULL, 0, sizeof(_Type) * _Count)))
#define FREE(_Ptr, _Type, _Count) ((_Type *)(LRT_Reallocate(_Ptr, sizeof(_Type) * _Count, 0)))

// convenient macro for allocating objects.
#define ALLOCATE_OBJ(_Type, _ObjectType) ((_Type *)LRT_AllocateObject(sizeof(_Type), _ObjectType))

// tells C++ compiler to treat the code as C source.
#ifdef __cplusplus
}
#endif

#endif