#ifndef LOXCRT_PRELUDE_H
#define LOXCRT_PRELUDE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifdef LOXMANGLE
#error "critical error: conflict definition of Lox runtime symbols"
#else
#define LOXMANGLE(_Name) __loxcrt__##_Name
#endif

#endif