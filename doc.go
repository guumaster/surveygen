/*

# Usage as CLI:

	surveygen is a CLI tool to generate go code for surveys from a yaml definition file.

	Usage:
	  surveygen [flags]

	Flags:
	  -h, --help           help for surveygen
	  -p, --path strings   path to look for survey definitions (default [$CWD])


# Usage as module:

cd into your project root folder, then run:

```
	$> go run github.com/guumaster/surveygen . --path path/to/your/survey
```


# Usage with `go generate`

Add the following line to your main or any other .go file you prefer:

```
	//go:generate surveygen --path demo --path path/to/your/survey --path another/path
```

*/
package main
