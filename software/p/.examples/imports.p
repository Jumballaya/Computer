/**
 *
 * Imports
 *
 */
import io
import Point from './struct.p'

fn main() void {
  Point p = new Point(10, 10)

  io.println("Data: (%s)", p.toString())
}
