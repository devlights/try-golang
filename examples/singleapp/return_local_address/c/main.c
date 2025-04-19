#include <stdio.h>

int *getvalue(int x)
{
    return &x;
}

int main(void)
{
    int *p1 = getvalue(1);
    int *p2 = getvalue(2);
    printf("%d,%d\n", *p1, *p2);

    return 0;
}