/**
 *
 * Simple math in plang
 *
 *
 */
import math


fn add(int x, int y) int {
  return x + y
}

fn multiply(int x, int y) int {
  return x * y
}

fn triangle_area(float base, float height) float {
  return base * height * 0.5
}

fn circle_area(float r) float {
  return math.PI * math.pow(r, 2)
}
