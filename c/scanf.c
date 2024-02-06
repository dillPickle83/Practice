#include <stdio.h>

int main(void)
{
	// The char varable can't be empty
	char c = 'a';
    printf("Enter a char:");
    scanf("%c", &c);
	printf("%c\n", c);

    int n=0;
    printf("Enter a number:");
    scanf("%d", &n);
	printf("%d\n", n);

	// The placeholders for scanf and printf are different and is mandatory for scanf but can be either lf or f for printf
	double r = 0;
    printf("Enter a double:");
    scanf("%lf", &r);
	printf("%f\n", r);

	int n1,n2,n3;
	n1=n2=n3=0;
    printf("Enter 3 numbers:");
    scanf("%d %d %d", &n1, &n2, &n3);
	printf("sum: %d\n", n1+n2+n3);

	// If we enter more than 1 word as the input only the first word will be stored in the string variable
	char string[50];
	printf("Enter a string: ");
	scanf("%s", string);
	printf("%s\n", string);
    return 0;
}
