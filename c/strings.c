#include <stdio.h>
#include <string.h>

int main(void){
    char test[5]="test";
    // When a string is defined as an array of characters, the final
    // index of the array will be the terminator character "\0" so if
    // the array size is 5, the length of the string needs to be 4
    for(int i=0; i<5; i++)
        printf("test[%d]=%d\n", i, test[i]);

    char random_str[]="Some random string?";
    str_length=strlen(random_str);
    printf("String length is: %d". str_length);

    int count_s = 0
    for(int i=0; i<str_length; i++)
        if (random_str[i] =="s" || random_str[i] == "S")
            count_s++;

    printf("No of s in the string: %d", count_s);
}
