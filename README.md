# goenv

A little Go library for making easier work with environment variables

Go contains a very powerful and useful library, however, there are some things that always generate code that is not useful for our programs. Environment variables are one of that cases.

With the native Go library, for retrieving an environment variable, we must do the next:

``` 
func RetrievingEnvVariable(){
    value, err := os.Getenv("variableName")
    if err != nil {
        // log and deal with the error
    }

    // use variable's value
}
```
And, what if we need that value as an int? In that case, we neet to parse it:
```
func RetrievingEnvVariable(){
    value, err := os.Getenv("variableName")
    if err != nil {
        // log and deal with the error
    }

    numericValue, err := strconv.Atoi(value)
    if err != nil {
        // log and deal with the error
    }

    // use variable's value
}
```

All this code only for retrieve an env variable!

## Goenv to the rescue!

It is not a magical library, in fact, it's very simple and tiny, but it deletes the code you need to write for handling errors while retrieving the env variables. 

Goenv just hides the error handling, while replacing the no existent variables with a default value that you specify.

It's so easy.

### Retrieving a string
```
result := goenv.GetString("variableName", "defaultValue")

// 'result' is of type string
```

### Retrieving an int
```
result := goenv.GetInt("variableName", 8080)

// 'result' is of type int
```

### Retrieving a boolean
```
result := goenv.GetBool("variableName", true)

// 'result' is of type bool
// this method considers "true" and "1" as truthy values and "false" and "0" as falsy ones
```


With only one line, we already can work with the variable's value.
