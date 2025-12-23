#include <stdio.h>

/* Print the conversion of fahrenheit to celsius from 0-300 */

# define lower 0
# define upper 300
# define step 20

/* Can define the variables like this at the top of the code to make the code changes later more accessible, instead of initializing in the function and then setting a value */

void main()
{
		float fahr, cels;

		printf("fahrenheit\tcelcius\n");
		fahr = lower;
		while (fahr <= upper)
		{
				// fahr - 32.0 => even if the fahr var was int, it will be considered as a float since the 32 is a float (by definition 32.0)
				cels = (5.0/9.0) * (fahr - 32.0);
				/* The formatting here says, fahrenheit is displayed in 3 character spaces with right side formatting and 0 digits after the decimal.
				   For celcius, it's displayed in 6 character spaces with right side justification and 1 value printed after the decimal. */
				printf("   %3.0f    \t%6.1f\n", fahr, cels);
				fahr = fahr + step;
		}

		// Print the same but for celsius
		fahr = lower;
		printf("\n\ncelsius \tfahrenheit\n");
		for (cels=upper; cels>=lower; cels=cels-step)
		{
       			fahr = ((9.0/5.0) * cels) + 32;
			printf("%6.0f\t\t    %3.0f\n", cels, fahr);
       		}
}

