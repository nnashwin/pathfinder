# pathfinder
> Easy, intuitive file creation and verification

[![Build Status](https://travis-ci.org/ru-lai/pathfinder.svg?branch=master)](https://travis-ci.org/ru-lai/pathfinder)

## Install
```
$ go get github.com/ru-lai/pathfinder
```

## Usage
```
import "github.com/ru-lai/pathfinder"

exists := pathfinder.DoesExist("path/to/file/Or/Dir")
if exists == false {
	// do some things
}

err := pathfinder.CreateFile("path/to/file")
if err != nil {
	// log the error however you want
}
```

## API

### pathfinder.DoesExist(path string)

Returns true if the path exists and false if it doesn't.

### pathfinder.CreateFile(path string)

Returns an error.  This will be nil if the file was able to be created.

### pathfinder.CreateDir(path string)

Returns an error.  This will be nil if the dir was able to be created.

## License

MIT © Tyler Boright
