#ifndef LOXCRT_PRELUDE_H
#define LOXCRT_PRELUDE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifdef LOXMANGLE
#error "critical error: conflict definition of Lox runtime symbols"
#else
/**
 * The Lox C Runtime (LOXCRT) Mangler Macro.
 *
 * To avoid naming conflicts, LOXCRT symbols is mangled. In the current version, any API
 * inside runtime (e.g. GC) is mangled with prefix "_LOXCRT_". Mangling of generated code
 * from Lox source is not specified.
 *
 * LOXCRT source uses [LOXMANGLE] macro to declare and define APIs, while the mangling of
 * user code is handled by [loxcc]. LOXCRT APIs do not depend on any user code except
 * entrypoint, so the entrypoint is mangled as a runtime API.
 *
 * Thus, calling a user-defined function will not lead to implicit errors of calling the
 * runtime API, or C standard library functions. If the generated C code is modified and
 * compiled with other 3rd-party libraries, the modifiers should be responsible for
 * dealing with name conflicts.
 */
#define LOXMANGLE(_Name) _LOXCRT_##_Name
#endif

#endif