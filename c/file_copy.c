#include <stdio.h>

int main()
{
	int c;
	double nc;

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
	//while ((c = getchar()) != EOF)
	//	putchar(c);

	printf("\nPrinting the number of characters in the string\n");
	// Can write the above while loop in a for loop like below
	for (nc = 0; getchar() != EOF; ++nc)
		;
	printf("%.0f", nc);
}
