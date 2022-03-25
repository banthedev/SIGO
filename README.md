`Version 1.0.1` [![Go Test](https://github.com/bryanmontalvan/sigo/actions/workflows/go-build-tests.yml/badge.svg)](https://github.com/bryanmontalvan/sigo/actions/workflows/go-build-tests.yml)

# Should I Go Out? (SIGO)

Programmers tend to not go out often simply because we're too busy trying to push a new feature to our side projects,
or we've been struggling with a bug fix at work. And whenever we do wish to go outside the weather may be unappetizing which 
may lower our motivation to buy more caffine.

Introducing **SIGO**, a command line interface with the convenient feature of telling you if you should go outside
depending by your location. 

Built using go and its [flag package](https://pkg.go.dev/flag) alongside with the [Weather API](https://www.weatherapi.com/)

# Installation Guide
Start by cloning the repository
```c
// HTTPS
git clone https://github.com/bryanmontalvan/sigo.git
```
```c
// GitHub CLI
gh repo clone bryanmontalvan/sigo
```
After cloning the repo you you have two options .. 
1. Run the CLI in a containerized environment (If you do not have Go installed or you do not wish to install Go)
2. Building the binary yourself (If you already have Go installed)
**Note:** Option 2 is more tailored for people who wish contribute/develop in a containerized environment

(*If you wish to run a containerized version of the CLI click [here]()*)

## Building Manually
Move into the `/src` directory and build the go binary
```go
go build
```
And that's it for the setup, now you roam free and explore

## Building using Docker
In the root directory of the project do the following
```bash
make docker-dev
```
You will now find yourself in a isolated containerized development environment where you can code freely without installing Go on your system. This container process is seperate from your own filesystem.

Now check to see if Go has been installed correctly 
```go
go version
```
## Add your own API key
Head over to [Weather API](https://www.weatherapi.com/) and create an account so you can retrieve your API key.
Now head over to the `src/store/` folder, here create a text file called `APIKEY.txt` and store your API key within that file

# Getting Started
In order to use the SIGO CLI you must move into the src folder `cd src`. There you have 3 main commands which all require a `./sigo` prefix

| Syntax  | Description | Example     |
| :---    |    :----:   |          ---: |
| get     | displays inputted city       | ./sigo get -city Boston |
| add     | stores inputted city into local-save        | ./sigo add -city Austin     |
| saved   | displays stored city |  ./sigo saved |

**get:** \
`./sigo get -city <cityname or zipcode>`
- Requires the `-city` flag 

**add:** \
`./sigo add -city <cityname or zipcode>`
- Requires the `-city` flag 
- Its recommended to use your zipcode for better accuracy since multiple countries have similar city names

**saved**: \
`./sigo saved`
- Requires no flags!!!

(**Note:** Currently only zipcodes from USA, Canda, and UK work!)

## Contributing and Community
Feel free to leave any suggestion, feedback, or opinion by opening an issue on this repo! 
Any and all contributions will be greatly appreciated!
