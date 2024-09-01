#include <stdio.h>

void print_array(int arr[]);
// Function declaration since the function is defined after main()

int main(void){
    int grade[5];

    for(int i=0; i<5; i++){
        printf("grade[%d]=", i);
	scanf("%d", &grade[i]);
    }

    printf("\nThe array is:\n");
    print_array(grade);

    printf("\nPrinting predefined array\n");
    int temp_array[3] = {23, 83, 654};
    // In this declaration, we can also not specify the number of elements
    // to be present in the array and the compiler will assume that in runtime
    print_array(temp_array);
}


void print_array(int arr[])
// Since the function only prints the values, return value is void
{
    for(int i=0; i<5; i++){
        printf("arr[%d]= %d\n", i, arr[i]);

    }
}
