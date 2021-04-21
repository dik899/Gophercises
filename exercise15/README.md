# Recover

Exercise 15 

This my implementation of the exercise 15 named Recover.

In the recover exercise we learned how to create some HTTP middleware that recovers from any panics in our application and renders a stack trace if we are in a local development environment.

This stack trace is displayed in browser. User should be able to navigate to these source code files. For this I'm using Chroma library which is capable of highlighting a line in source code.

For parsing line from stack trace I'm using custom parsing logic.

There is single file to do all this and the test coverage is 100%.
