#include <stdio.h>

int main(void)
{
    int x = 5;
    printf("x: %d, x+1: %d\n", x, x+1);

    char c = 'Q';
    printf("c: %c\n", c);

    double d = 23;
    printf("d: %lf\n", d);

    float f = 4.56;
    printf("f: %f\n", f);

    char str[] = "Star Platina";
    printf("str: \"%s\"\n", str);

    //Make the value with a width of 10
    printf("|||%10f|||\n", f);

    // Make the value left aligned with width 8 and precision of 3
    printf("|||%-8.3f|||\n", f);

    return 0;
}
