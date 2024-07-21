# Sence
As the name implies, it will provide with a signle instance of an object, and guarantee that there are no duplicates.

# Objectives
- Need a signle, shared value, of some particular type.
- Restrict object creation of some type to a single unit along the entire program

# Situations
- When you want to use the same connection to a database to make every query
- When you open a SSH connection to a server to do a few tasks, and don't want to reopen the connection for each task.
- If you need to limit the access to some variable or space, you use a singleton as the door to this variable.
- If you need to limit the number of calls to some places, you create a signleton instance to make the calls in the accepted window

# Example
- Write a program counter that holds the number of times it has been called during program execution
  
# Requirements
- When no counter has been created before, a new one is created with the value equal 0
- If a counter has already been created, return this instnace that holds the actual count
- If we call the method AddOne, the count must be increamented by 1