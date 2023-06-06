#include <stdio.h>
#include <stdlib.h>
int main(void)
{
	if (puts("Hello World!") == EOF) {
		return EXIT_FAILURE;
	}
	printf("%s\n", "Hello, World!");
	return EXIT_SUCCESS;
}
