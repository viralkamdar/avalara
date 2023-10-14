# Avalara dictionary application
This application provides meaning of the word passed as a parameter

## Compile the application
```
./make build
```
This application can be built on different platforms depending on the operating systems used

## Run the application
This application requires an API key to query the Merriam-Webster dictionary. To run the application, set the API_KEY variable in the Makefile
```
./make run
```
Alternatively, after building the application, you can run it by it's name
```
API_KEY=<Your API_KEY> ./avalara -word=<word to search>
```