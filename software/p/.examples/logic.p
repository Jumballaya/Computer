/**
 *
 * Logic and Flow
 *
 */
import io
import fmt
import atoi from strings
import randInt from random


fn main() void {
  int answer = randInt(10)

  bool game = true
  for (game) {
    char* res = io.prompt("Guess a number between 1 and 10")
    int guess = atoi(res)

    if (guess == answer) {
      io.println("You are correct!")
      game = false
    } elif (guess < answer) {
      io.println("You guessed a little too low")
    } else {
      io.println("You guessed a little too high")
    }
  }
}
