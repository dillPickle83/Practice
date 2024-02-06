#include <stdio.h>

int main(void)
{
	int n=0;
	printf("Enter a number between 1 and 15: ");
	scanf("%d", &n);

	if(n<1)
	{
		printf("The number should be greater than 1\n");
	}
	else if(n>15)	printf("The number should be lesser than 15\n");
	else
	{
		if(n%2==0) printf("%d is even\n", n);
		else printf("%d is odd\n", n);
	}
	return 0;
}
