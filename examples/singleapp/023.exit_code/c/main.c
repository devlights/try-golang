#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
    int i = atoi(argv[argc-1]);
    if (i >= 5) {
        exit(99-i);
    }
    
    if (i%2 == 0) {
        exit(EXIT_FAILURE);
    }

    exit(EXIT_SUCCESS);
}