/**
 * The Lox C Runtime (LOXCRT) Prelude.
 *
 * To avoid naming conflicts, LOXCRT symbols is mangled. In the current version, any
 * API inside runtime (e.g. GC) is mangled with prefix "LRT_". Mangling of generated
 * code from Lox source is not specified, but should be mangled in a different way.
 *
 * LOXCRT APIs do not depend on any user code except entrypoint, so the entrypoint is
 * mangled as a runtime API.
 *
 * Thus, calling a user-defined function will not lead to implicit errors of calling
 * the runtime API, or C standard library functions. If the generated C code is
 * modified and compiled with other 3rd-party libraries, the modifiers should be
 * responsible for dealing with name conflicts.
 */

#ifndef LOXCRT_PRELUDE_H
#define LOXCRT_PRELUDE_H

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

    /**
     * The runtime panic handler.
     *
     * When a runtime error is detected, this function should be called. It prints the
     * given message to `stderr`, and exits the program with a non-zero value.
     */
    void LRT_Panic(const char *message);

    // Compilation Flags
    //
    // These flags are of great use when debugging the runtime implementation. However, as
    // embedded template, they should not be defined here; pass them to the system CC by
    // build-config.

#define GC_TRACE  // Output information when GC
#define GC_STRESS // Run GC as more as it can to detect bugs
#undef GC_STRESS
#undef GC_TRACE

#ifdef __cplusplus
}
#endif

#endif