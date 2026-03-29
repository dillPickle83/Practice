#include <stdio.h>

int main(){
    int *pa;
    int arr[10];

    // Assign values to the array
    for (int i=0; i<10; i++){
        arr[i] = i + 1;
    }

    int arr_size = sizeof(arr)/sizeof(arr[0]);
    printf("The elements of the array are:\n");
    for (int i=0; i<arr_size; i++){
        printf("arr[%d]: %d\n", i, arr[i]);
    }
    printf("\n");

    pa = &arr[3];
    int x = *pa;
    printf("Printing element 3 of the array:\n%d\n", x);
    x = x + 23;
    printf("X modified value: %d\n", x);
    printf("Original array value: %d\n", arr[3]);
    // pa is the address of arr[3] and since arrays are stored sequentially
    // pa++ or pa+1 points to the address of arr[4]
    printf("+4 value referenced by pointer: %d\n", *(pa+4));
}
