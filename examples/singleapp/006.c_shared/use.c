#include <stdio.h>
#include <stdlib.h>
#include "libgoadd.h"

int main() {
    int x = 111;
    int y = 222;
    int z = 0;

    z = GoAdd(x, y);
    printf("[FROM      C] %d\n", z);

    return EXIT_SUCCESS;
}