#include "entrypoint.h"
#include "gc.h"

int main(int argc, const char *argv[])
{
    LRT_InitializeGC();
    LRT_Entrypoint();
    LRT_FinalizeGC();
    return 0;
}