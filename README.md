# Some Useful utilties when building REST API services with golang

## Features included
- Validations for parameters, query strings, and request bodies
- Format validation errors into a user-friendly format
- Default presenter (currently, only JSON format is supported)
- Logger utility that prints with colors, filename, and line number

### Disclaimer:
This is not a new package. It is just a wrapper around [go-playground/validator](https://github.com/go-playground/validator) and [github.com/fatih/color](https://github.com/fatih/color). It makes them a little easier to use.
