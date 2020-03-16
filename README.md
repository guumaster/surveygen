[![GitHub Release](https://img.shields.io/github/release/guumaster/surveygen.svg?logo=github&labelColor=262b30)](https://github.com/guumaster/surveygen/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/guumaster/surveygen)](https://goreportcard.com/report/github.com/guumaster/surveygen)
[![License](https://img.shields.io/github/license/guumaster/surveygen)](https://github.com/guumaster/surveygen/LICENSE)

# surveygen

> generate interactive surveys from yaml files

Generate code for surveys that you can use with [github.com/AlecAivazis/survey/v2](https://github.com/AlecAivazis/survey/v2).

It will save you some time if you need to maintain large surveys and prefer to keep a simple yaml file instead of code.

Check [an example yaml](example/demo/my_awesome_survey.yaml) to see how to create it and use it in [code](example/main.go).


## Installation

### Binaries 
Go to [release page](https://github.com/guumaster/surveygen/releases) and download the binary you need.

### Module
go get github.com/guumaster/surveygen 



### Usage as CLI:

```
	surveygen is a CLI tool to generate go code for surveys from a yaml definition file.

	Usage:
	  surveygen [flags]

	Flags:
	  -h, --help           help for surveygen
	  -p, --path strings   path to look for survey definitions (default [$CWD])

```



## Usage as module:

cd into your project root folder, then run:

```
	$> go run github.com/guumaster/surveygen . --path path/to/your/survey
```


## Usage with `go generate`

Add the following line to your main or any other .go file you prefer:

```
	//go:generate surveygen --path demo --path path/to/your/survey --path another/path
```

### TODO

Features that I'd like to add: 

  * [ ] Some more documentation
  * [ ] Allow usage as module to generate questions on the fly instead of generating code
  * [ ] Validate yaml content, right now it may panic or generate wrong code


### References

  * [spf13/cobra](https://github.com/spf13/cobra)
  * [AlecAivazis/survey](https://github.com/AlecAivazis/survey/v2)


### LICENSE
 [MIT license](LICENSE)


### Author(s)
* [guumaster](https://github.com/guumaster)
