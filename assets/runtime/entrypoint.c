#include "entrypoint.h"

int main(int argc, const char *argv[])
{
    // currently, we have no runtime preparation needed.
    // GC is not yet introduced.
    LRT_Entrypoint();
    return 0;
}