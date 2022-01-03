package main

/*
Select versus switch in Go (choose one or more)?
a: A `select` is only used with channels
b: A `switch` is used with concrete type
c: A `select` will choose multiple valid options at random, while a `switch` will go in sequence
d: A `switch` is only used with channels, a `select` is used with concrete types
 */

// Answer: a, c
/*
Clarification:
In switch each of the case statement is an expression while in select each of the case statement is either send or receive operation on a channel.
Switch statement is non-blocking.
A switch will go in sequence to select the matching case.
Switch has two forms: expression switches and type switches while select has only one form.
A select statement  will block if send and receive operation is blocked in all the case statements and default block is not present. Default block makes the select non-blocking as default case will be executed if all the other cases are blocked.
*/