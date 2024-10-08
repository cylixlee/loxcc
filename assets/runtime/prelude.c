#include "prelude.h"
#include <stdio.h>  // for error output
#include <stdlib.h> // for manually exit the program

void LRT_Panic(const char *message)
{
    fprintf(stderr, "runtime error: %s\n", message);
    exit(EXIT_FAILURE);
}