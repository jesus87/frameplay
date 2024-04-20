
# Go JSON Validator and HTTP Request Sender

This Go program validates input in structured JSON format, logs the input to STDOUT, and sends a request to an HTTP-based API with the input as part of the request message.

## Requirements

- Go installed on your machine.

## Installation

1. Clone this repository to your local machine:

git clone https://github.com/jesus87/frameplay.git
2. Navigate to the project directory:
cd frameplay

3. Update dependencies using Go modules:
go mod tidy


## Usage

1. Open the `main.go` file in a text editor.

2. Modify the `inputData` variable at the top of the file to include your JSON input.

3. Save the file.

4. Run the program:

go run main.go


5. The program will validate the input JSON, log it to STDOUT, and send a request to the API.

## Example

For example, if you have the following JSON input:

```json
{
    "name": "John",
    "age": 30
}

Running the program will validate the input, log it to STDOUT, and send a request to the API with the input as part of the request message.

## Unit Tests
To run unit tests, use the following command:
go test
