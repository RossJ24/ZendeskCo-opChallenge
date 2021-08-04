# ZendeskCo-opChallenge

## About The Project

A CLI for Zendesk's ticket API built using Golang.

Built for the Zendesk CO-OP coding challenge.

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites


* go 16.16.5 or later [go installation link](https://golang.org/dl/)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/RossJ24/ZendeskCo-opChallenge.git
   ```
2. Install NPM packages
   ```sh
   go get
   ```
3. Include a `.env` file in the base directory with the environment variables 'EMAIL' and 'PASSWORD'
   
   Ex:
   ```
    EMAIL=youremail@email.com
    PASSWORD=yourpassword
   ```
4. Execute the program `go run main.go` or build an executable `go build -o <executable_name>` `./<executable_name>`
