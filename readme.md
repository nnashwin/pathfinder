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

### pathfinder.DoesExist(path string) bool

Returns a boolean denoting whether or not the path exists.

### pathfinder.CreateFile(path string) error

Returns an error.  This will be nil if the file was able to be created.

### pathfinder.CreateDir(path string) error

Returns an error.  This will be nil if the dir was able to be created.

### pathfinder.GetDir(path string) string

Returns the dir from a filepath string.  If the filepath string is simply a file, it will return a blank string

```
pathfinder.GetDir("/.dir/dir/dir.file")
// => "/.dir/dir/"

pathfinder.GetDir("dir.file")
// => ""
```

## License

MIT © Tyler Boright
