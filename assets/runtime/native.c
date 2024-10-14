#include "native.h"
#include <time.h> // for clock()

LRT_Value LFN_clock(size_t arity, va_list params)
{
    if (arity != 0)
    {
        LRT_Panic("native function clock() requires no args.");
    }
    return NUMBER((double)clock() / CLOCKS_PER_SEC);
}