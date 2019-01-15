/**
 *
 *
 * Structs
 *
 */
 import fmt


struct Point {
  public:
    int x
    int y
    toString() char*

  private:
    char* name
}

fn Point::Point(int x, int y) void {
  this.x = x
  this.y = y
  this.name = "Point"
}

fn Point::toString() char* {
  return fmt.sformat("%s - (%d, %d)", this.name, this.x, this.y)
}
