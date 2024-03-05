#include <stdio.h>

int add(int a, int b);
int mult(int x, int y);
// The above statements are called function declarations which needs to be used if the function is defined after the function call is present.
// Eg: the main() function calls the add() function but add() is defined after the main function.
// The compiler readsd the code from tpo to bottom, so it will throw a compilation error is the function is called before the function header

int main(void)
// The above line is called the function header
{
    int a, b;
    printf("Input a:");
    scanf("%d", &a);
    printf("Input b:");
    scanf("%d", &b);

    int add_res = add(a, b);
    printf("Addition: %d\n", add_res);

    int mult_res=mult(a, b);
    printf("Multiplication: %d\n", mult_res);
    return 0;
}

int add(int a, int b)
{
    int result = a + b;
    return result;
}

int mult(int x, int y)
{
    int result = 0;
    for (int i=0; i<x; i++)
        result = add(result, y);
    return result;
}
