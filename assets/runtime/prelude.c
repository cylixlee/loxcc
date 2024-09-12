#include "prelude.h"
#include <stdio.h>
#include <stdlib.h>

void LRT_Panic(const char *message)
{
    fprintf(stderr, "runtime error: %s\n", message);
    exit(EXIT_FAILURE);
}