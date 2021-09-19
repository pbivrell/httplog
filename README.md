# httplog

# Overview 
httplog is a go package that attempts to simplify logging around http requests. This package exposes a structure containing the _useful_ data from any given http request. Then it defines an interface for writing this data. In the loggers director we have supplied some implementations for the logger to the following outputs.

|---|---|---|
| Implementation | Use case | 
| [logrus](https://github.com/sirupsen/logrus) | Text base logging | 
| [influx](https://www.influxdata.com/) | Time series database persistent logs | 

The package also offers a helper Middlewear function which can wrap an http.Handlerfunc to do the logging.

## Usage

This is a package only, it does nothing on it's own must be imported and used by another go package.

```
go get -u github.com/pbivrell/httplog
```

## Contributing

Feel free to add more implementations for various logging outputs, I only ask you supply a meaningful `use case` in the table above

## License 

IDK. Do whatever you want with this, its not special it's convenient
