# Regular expression To Nfa Tester

This is an app that converts a regex from infix to postfix and tests it against a string. 

## Getting Started

To start off my project I first followed video tutorials provided by Dr. Ian Mcloughlin which were provided on my college's learnonline hub at GMIT. Following the videos was an easy task but it took me a few watches of each one to fully grasp and understand the complexity of the code provided as well as actually getting my head wrapped around what the project done. 

### Design Process
```
1.) After following the videos I was left with a line of hard coded code that would test the postfix of a regular expression against an nfa. The result would either return back true or false on depending on whether or not the string matched. The regex that was tested most for me was "ab.c*|" which was tested with many different strings to see the test result such as "ab", "abc" and "ccccccc". Tests would return the desired result and all was well. 

2.) For the project I was also ment to add the ability to test with the chars "?" and "+" among other things but before I decided to add any of these I first thought it would make more sense to add the ability for the user to enter there own infix regex which would then be converted to pofix and the tested against a string, this is where the whole project fell apart for me. 

3.) I first added a getinput function to the program which would allow user input, I tested the function with a basic side program I made in go which allowed a user to enter and add two numbers together. 

4.) After I added the getinput function I then needed to add it to my main method. I decided that the user would first be asked to enter the infix of the regex and then they would been shown the pofix of the regular expression they entered. 

5.) I tested this first to see if I was on the right track and low and behold I was, what ever infix the user entered would be returned back as pofix on the next line. Onto the next step.

6.) I added the ability for the user to enter the string that needed to be tested.

7.) Next a match would be returned and would either be true or false, for some reason my tested strings and regex which should return as true would instead return as false, this meant there was clearly an error in my code.

8.) I added a tester function to my poregtonfa function that would return an error if the nfa was bigger than one

9.) on top of this, I then added back in the hardcoded line that was provided to me to just make sure it wasn't an error on another end. 

10.) after playing around with various expression and strings I noted that the programs nfa would become bigger than 1 for the input into the infix user line. I could test strings that were one char long in left for infix but anything bigger would return an error. 

12.) After multiple tests I could not pinpoint the source of my problem 


```

### Installing

A step by step series of examples that tell you have to get a development running



```
Download zip file of project.
```
```
Unzip folder to drive
```
```
navigate via cmd to the folder 
```
```
Type set GOPATH = c\GoProject2018 and hit enter
```
```
Type set GOBIN = c\GoProject2018\bin and hit enter
```

For Running the program:

```
Type go run src/main/nfa.go 
```




Please note: bin folder and pkg folder were not uploaded during initial commit, you can make them by just adding an empty folder called "bin" and one called "pkg" in the first directory with the src folder

## Running the tests

User will first be asked to enter a regex to convert from infix to pofix 
Infix will be returned as pofix 
User will be asked to enter a test string 
Match result will be displayed (still working on this, works for single chars and sometimes 3 chars, sometimes doesn't) 
Hardcoded match result will also be returned (this works) 

Example of pofix with true string: 
ab.c*| = ccccccc



## Built With

* Go lang 
* Vs Code



## Authors

* **Aron O' Malley** 





## Acknowledgments

* Dr. Ian Mcloughlin


