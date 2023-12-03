# Day 1

The problem for day 1 is basically finding number from a string respectively from left side and right of the string. I will discuss the approach that i use on each subsection.

## Section A

Section A requires us to find only number on left and right side on the string. The approach that i take is basically loop from the left until find one number and then stops the loop. The same goes on the right side, after that to get the combined number, we only need to concat and convert it into number.

## Section B

Section B changes the problem for looking numbers in form of literal number or numerical number. For example, "one", "two", 1, 2 and so on. There is 3 ways we can solve this

### Approach 1

Check each substring and if we find the corresponding number "one", "two" or "1", "2", etc then wefirst need to transform it into number representation. We also do this from the right side of the string and then concat both of the number.

### Approach 2

Instead of checking each substring, what we could do instead is convert the occurences of number into the numerical representation. For example "onetwothree" will be converted into "123" and this reduce the problem backs to section A.

But we also need to be careful on cases such as "twoone", "oneight", and so on. Since the numeric isn't that much, we could create map representation for left side and right side ( right side is rverted "eno", "owt" ).

### Approach 3

The third approach takes a complete different path. What we will be using is called as Aho Corasick alogrithm. Basically, if we have dictionary of word, we wanted to find the occurences of those word on input string. Our dictionary is the "1", "one", "two", and so on. 

I opted for this approach because i never actually implemented aho corasick, you can look up on the source code if you are interested on how it works. 

