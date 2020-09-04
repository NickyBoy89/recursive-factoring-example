# Info

This repo contains two Go scripts, one that goes through the number to factor linearly, and the other that recursively goes through the numbers.

# What the scripts do

The primary functions of these programs is to take in a number and factor it

Ex: The number 4 factors out into 2 * 2, and the number 12 factors out into 2 * 2 * 3

# What the normal script does

The normal script takes the number and loops through it, from 2 (Since 1 is always a factor, and does not do anything to the final product) to the number to factor.

This program has a linear O(n) computation time

# What the recursive script does

The recursive script does the same thing as the normal script, but instead, when it encounters a factor >= 2, it creates a branch of recursion that reduces the number of computations by a factor of whatever number was the factor

Example:

1. For the input number 12, try to find the first factor
1. Found the factor 2, call the factor function again with the number that is left over (6), which reduces the number of computations by the factor (2)
1. Find the next factor, which is again 2, and do the same thing as the step above
1. THis time, the number remaining (3) is not factorable, so the function will reach the end of the loop, which is the exit case, and the higher logic will start getting answers

It is important to state that with a prime number, each solution will take the same number of steps, and therfore, recursion will not be of any help.

Where this solution really shines is with the use of a highly composite number ([Wikipedia](https://en.wikipedia.org/wiki/Superior_highly_composite_number)), which will take exponentially less time with the recursive formula
