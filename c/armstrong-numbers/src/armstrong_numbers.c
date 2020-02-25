#include "armstrong_numbers.h"
#include <math.h>

int digit_count(int n);

int is_armstrong_number(int candidate) {
	if (candidate <= 9 || candidate < 0) 
		return 1;

	int power = digit_count(candidate);
	int armstrong_sum = 0;

	do {
		int digit = candidate % 10;
		candidate /= candidate / 10;
		double raised = pow((double) digit, (double)power);		
		armstrong_sum += (int)raised;
	} while (candidate > 0);

	return armstrong_sum == candidate ? 1 : 0;
}

int digit_count(int n) {
	if (n < 10)
		return 1;

	int count = 0;
	do {
		n /= 10;
		count++;
	} while (n > 0);
	return count;
}
