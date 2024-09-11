#ifndef LOXCRT_ENTRYPOINT_H
#define LOXCRT_ENTRYPOINT_H

#include "prelude.h"

/**
 * Generated Lox entrypoint.
 *
 * This entrypoint is generated by [loxcc], and is basically AST translation from Lox
 * source code.
 *
 * Lox is a very simple language designed for the book Crafting Interpreters, which does
 * not support modules and packages. That makes compilation targeting C much easier.
 * However, Lox is a dynamically-typed language with OOP and GC support, which means
 * some runtime preparation is needed.
 */
void LOXMANGLE(entrypoint);

/**
 * The C standard entrypoint.
 *
 * Lox C Runtime (LOXCRT) will take over this "real" entrypoint, and do some
 * initialization work (e.g. GC preparation). When the runtime is ready, the generated
 * entrypoint will be called from this.
 */
int main(int argc, const char *argv[]);

#endif