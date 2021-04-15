# golang-file-reader

To convert any multiline text file based to a single line text file. Initiated this project because my work project needs to convert a multiline python code file to a single line python code.

## Golang

Golang version : 1.16.3

## Go Fiber

Gofiber version : 2.7.1

## Dependencies

- Golang's default dependencies
- Gofiber
- viper (v1.7.1) (https://github.com/spf13/viper)

### Environment Variables Setting

Make file with name :
> config.env

And then put this in your config.env :

```
FOLDER_PATH="[spesific directory file, in string]"
PORT=[the port you want]
```

### Routes Explanation

* GET /file-reader : to read file from your local directory. Query param given : ?filename=[your file name with file extension] 
* GET /file-receiver : upload your file from your local diretory, and magic will happen :wink: . Using form-data body type with key: `file` and value: your file
