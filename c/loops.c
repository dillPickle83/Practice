#include <stdio.h>

int main(void){
    int i = 0;

    while (i<5){
        printf("i: %d\n", i);
        i++;
    }

    printf("While loop is over.\n");

    int d = 0;
    do{
        printf("d: %d\n", d);
        d++;
    } while (d<5);

    printf("Do while loop is over.\n");

    int m = 6;
    do{
        printf("m: %d\n", m);
        d++;
    } while (m<5);

    printf("Do while loop is with d as 6 over.\n");

    for (int j=0; j<5; j++)
    {
        printf("j: %d\n", j);
    }

    printf("For loop is over.\n");

    return 0;
}
