/**
 * For simplicity, we just manually write functions to match the LoxCC's requirement.
 * There's no difference from native functions and user-defined ones.
 */
#ifndef LOXCRT_NATIVE_H
#define LOXCRT_NATIVE_H

#include "prelude.h"
#include "value.h"
#include <stdarg.h> // for vararg functions

#ifdef __cplusplus
extern "C"
{
#endif

    // The FFI of clock() in C standard library.
    LRT_Value LFN_clock(size_t arity, va_list params);

#ifdef __cplusplus
}
#endif

#endif