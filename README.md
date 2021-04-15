# golang-file-reader

To convert any _**multiline text file based**_ to a _**single line text**_. Initiated this project because my work project needs to convert a multiline python code file to a single line python code.

## Golang

Golang version used : 1.16.3

## Go Fiber

Framework Go fiber version used : 2.7.1

## Dependencies

- Golang's default dependencies
- Framework Go [fiber](https://github.com/gofiber/fiber)
- [viper (v1.7.1)](https://github.com/spf13/viper)

## Environment Variables Setting

Make file with name :
> config.env

And then put this in your config.env :

```
FOLDER_PATH="[spesific directory file, in string]"
PORT=[the port you want]
```

## Routes Explanation

* GET /file-reader : to read file content (text only) from your local directory. 
  - Query param given : ?filename=[your file name with file extension] 
* GET /file-receiver : upload your file from your local diretory, and magic will happen :wink:. 
  - Using form-data body type with key: `file` and value: `[your file]`
* All routes return JSON response
  - Example response : 
  ```
  {
    "content": "your text file content\nbut with '\n' as a new line"  
  }
  ```
