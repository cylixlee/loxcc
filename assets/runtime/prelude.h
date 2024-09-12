#ifndef LOXCRT_PRELUDE_H
#define LOXCRT_PRELUDE_H

/**
 * The Lox C Runtime (LOXCRT) Prelude.
 *
 * To avoid naming conflicts, LOXCRT symbols is mangled. In the current version, any API
 * inside runtime (e.g. GC) is mangled with prefix "LRT_". Mangling of generated code
 * from Lox source is not specified, but should be mangled in a different way.
 *
 * LOXCRT APIs do not depend on any user code except entrypoint, so the entrypoint is
 * mangled as a runtime API.
 *
 * Thus, calling a user-defined function will not lead to implicit errors of calling the
 * runtime API, or C standard library functions. If the generated C code is modified and
 * compiled with other 3rd-party libraries, the modifiers should be responsible for
 * dealing with name conflicts.
 */

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

void LRT_Panic(const char *message);

#endif