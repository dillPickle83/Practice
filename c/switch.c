#include <stdio.h>
#include <ctype.h>
#include <string.h>

int main(void)
{
    char s[] = "aaBCADefGhIjKlmNOoPp";
    int length = strlen(s);
    int a = 0;
    int vowel = 0;

    for(int i=0; i<length; i++)
    {
        switch (toupper(s[i]))
        {
            case 'A':
                a++;
            case 'E':
            case 'I':
            case 'O':
            case 'U':
                vowel++;
        }
        // The above logic is called fallthrough logic
        // When the case is matched we need a break at the end of the case's execution to stop it there
        // In cases like this where there is no case, the program keeps evaluating the restuntil a break or the end
        // Here, it checks for vowels and increments if its a vowel and increments a and vowel if the occurence is an A/a
    }
    printf("A count: %d, Vowel Count: %d", a, vowel);

    return 0;
}
