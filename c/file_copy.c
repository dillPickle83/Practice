#include <stdio.h>

int main()
{
	int c;

	/*
	c = getchar();
	while (c != EOF)
	{
		putchar(c);
		c = getchar();
	}
	*/
	/* The above snippet of code can also be written like below */
	// The != operator takes precedence over the = operator and so the "c=getchar()" needs to be prioritized inside the paranthases
	// If it were "c=getchar() != EOF", getchar() != EOF would've been evaluated first before c were assigned with the value from getchar()
	// printf("%d", getchar() != EOF); => This expression prints 1, when the input is EOF
	while ((c = getchar()) != EOF)
		putchar(c);
}
