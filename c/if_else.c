#include <stdio.h>
#include <stdbool.h>

int main(void)
{
	int n=0;
    char is_even=false;
	printf("Enter a number between 1 and 15: ");
	scanf("%d", &n);

	if(n<1)
	{
		printf("The number should be greater than 1\n");
	}
	else if(n>15)	printf("The number should be lesser than 15\n");
	else
	{
		if(n%2==0)
        {
            printf("%d is even\n", n);
            is_even = true;
        }
		else printf("%d is odd\n", n);
	}
    if(!(is_even==true) && n>10) printf("The number is odd and greater than 10");
    else if(n==12 || n<10) printf("The number is 12 or less than 10");
    else printf("The number is %d", n);
	return 0;
}
