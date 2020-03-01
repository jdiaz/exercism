#include "darts.h"
#include "math.h"

#define OUTER_CIRCLE_RADIUS 10
#define MIDDLE_CIRCLE_RADIUS 5
#define INNER_CIRCLE_RADIUS 1

int in_circle_bounds(coordinate_t c, float lower, float upper) {
	if (c.x >= lower && upper >= c.y)
		return 1;
	return 0;
}

int score(coordinate_t c) {

	if (in_circle_bounds(c, -INNER_CIRCLE_RADIUS, INNER_CIRCLE_RADIUS)) {
		return 10;
	}
	if (in_circle_bounds(c, -INNER_CIRCLE_RADIUS, MIDDLE_CIRCLE_RADIUS)) {
		return 5;
	}
	if (in_circle_bounds(c, -MIDDLE_CIRCLE_RADIUS, OUTER_CIRCLE_RADIUS)) {
		return 1;
	}

	return 0;
}
