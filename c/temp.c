#include <stdio.h>

/* Print the conversion of fahrenheit to celsius from 0-300 */

void main()
{
		float fahr, cels;
		int upper, lower, step;

		lower = 0;
		upper = 300;
		step = 20;

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
}

