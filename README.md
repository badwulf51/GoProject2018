# Regular expression To Nfa Tester

This is an app that converts a regex from infix to postfix and tests it againts a string. 

## Getting Started

To start off my project I first followed video tutorials provided by Dr. Ian Mcloughlin which were provided on my college's learnonline hub at gmit. Following the videos was an easy task but it took me a few watches of each one to fully grasp and understand the complexity of the code provided as well as actually getting my head wrapped around what the project done. 

### Design Process
```
1.) After following the videos I was left with a line of hard coded code that would test the postfix of a regular expression againts an nfa. The result would either return back true or false on depending on wheteher or not the string matched. The regex that was tested most for me was "ab.c*|" which was tested with many different strings to see the test result such as "ab", "abc" and "ccccccc". Tests would return the desired result and all was well. 

2.) For the project I was also ment to add the ability to test with the chars "?" and "+" amoung other things but before I decided to add any of these I first thought it would make more sense to add the ability for the user to enter there own infix regex which would then be converted to pofix and the tested againts a string, this is where the whole project fell apart for me. 

3.) I first added a getinput function to the program which would allow user input, I tested the function with a basic side program I made in go which allowed a user to enter and add two numbers together. 

4.) After I added the getinput function I then needed to add it to my main method. I decided that the user would first be asked to enter the infix of the regex and then they would been shown the pofix of the regular expression they entered. 

5.) I tested this first to see if I was on the right track and low and behold I was, what ever infix the user entered would be returned back as pofix on the next line. onto the next step.

6.) I added the ability for the user to enter the string that needed to be tested.

7.) Next a match would be returned and would either be true or false, for some reason my tested strings and regex which should return as true would instead return as false, this meant there was clearly an error in my code.

8.) I added a tester function to my poregtonfa function that would return an error if the nfa was bigger than one

9.) on top of this, I then added back in the hardcoded line that was provided to me to just make sure it wasnt an error on another end. 

10.) after playing around with various expression and strings I noted that the programs nfa would become bigger than 1 for the input into the infix user line. I could test strings that were one char long in left for infix but anything bigger would return an error. 


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




Please note: bin folder and pkg folder were not uploaded during intial commit, you can make them by just adding an empty folder called "bin" and one called "pkg" in the first directory with the src folder

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc

