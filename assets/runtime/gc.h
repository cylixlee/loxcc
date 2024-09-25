#ifndef LOXCRT_GC_H
#define LOXCRT_GC_H

#include "prelude.h"

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

    struct LRT_Object; // external, in "object.h"

    typedef struct LRT_Object LRT_Object;

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

// tells C++ compiler to treat the code as C source.
#ifdef __cplusplus
}
#endif

#endif