# GoLang-Tutorials
Every go file should start with package "Folder name"


## Functions
for real:
for adding two numbers like 2 and 3 we don't need to depand on struct we can call a function and add them.

1. It's is a general action like it performs a specific task and can be called from anywhere.
2. Independent on its own means it is not tied to any type (or) struct.
3. A function need all the data as parameters to execute.


## Methods
for real:
In a phone i want to take a photo to take a photo i need a phone first here phone is the struct.

1. a block of code tied to a struct.
2. method is nothing but a function with a special receiver argument.
3. A method already has some data inside it struct and uses it automatically.


## 1. Using Go Routines to print wishes in different languages
I have created a function called birthdayWishes which runs two times according to the condition. And in the main I have used go keyword for the birthdayWishes fuction which runs simultaneously tells "Happy Birthday Sumanth" in two different languages and also used the waitgroup like the time.sleep to tell the OS to wait for the go routines and print the output.
go routines -- this are used to run simultaneously to the OS which are cheaper than the normal OS threads.

## 2. Palindrome in go 
Palindrome is nothing but the word which reads same from forward and backward.
Here i have created a function that the word is palindrome or not, I have used Sprintf keyword to convert numbers into strings and then told the for loop to stop when the length of word/2 then stop the loop and I have gave a condition that if the first letter of the word and last letter of the word is not equal to return false if they are equal to then return true. And in the main I gave a slice of numbers and in the for loop if the number matches the function condition return me that the number is palindrome or return me not a palindrome.

## 3. Factorial in go
Factorial is nothing but 5! = 5x4x3x2x1
Here I have Created a function that has a parameter n and the return value should be int and then in the for loop i told i less than or equal to n, if suppose if n is 3 then stop the loop when the value is 3 and the condition says multiply the result and store it in the same result. Then in the main I gave a value 3 and it will return me the factorial of 3.

## 4. Fibonacci in go
Fibonacci series is adding the preceding two numbers of the next number.
Created a function for the fibonacci series, intilized a, b and if the i is less than the series of numbers then stop the loop and printed a first and told to update the values of a, b by adding the preceding numbers. In the main i have just used the fibonacci function and told to give first 10 fibonacci series numbers.

## 4. Word Count in go
I have created words using slice and i have also created a word counter using maps which key as a word string and return in int means the count of the repeated words. In a for loop I told add +1 whenever the word repeats, again in a for i told to return what is the word and how many times it is repeating.