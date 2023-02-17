# gameofpig

pig - A command line tool to simulate a game of pig. It is a two-player game played with a 6-sided die.

Usage:</br>
&emsp;pig [strategy] [strategy]

Args:</br>
&emsp;strategy   The number between 1 to 100

Description:</br>
&emsp;This command line application accepts two numbers between 1 to 100 as a positional argument. Strategies for player 1 and player 2, and performs the game of pig simulation on it. If no strategies are provided, then it will return an error. </br> 

&emsp;If the number is out of range, the application will exit with an error message. Otherwise, it will perform the simulation and output the result.

Example usage: </br>
```
$ pig 10 15
Result: Holding at   10 vs Holding at   15: wins: 3/10 (30.0%), losses: 7/10 (70.0%)
```
