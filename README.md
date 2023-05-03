# iframeParser

iframeParser is a command-line tool that can be used to merge embedded HTML code into a single output file. This tool is useful for developers who work with embedded iframes and need to combine multiple iframes into a single HTML file.

## Installation

To use iframeParser, you can clone the repository from GitHub and build or run the program using Go.
```
git clone https://github.com/maxiputz/iframeParser.git
cd iframeParser
go run . /path/to/root/ index.html
```

Alternatively, you can build the executable file and run it:
```bash
go build .
./iframeParser /path/to/root/ index.html
```

## Usage

To use iframeParser, open a terminal or command prompt and navigate to the directory where the executable file is located. Then, run the following command:
```
./iframeParser /path/to/root/ index.html
```

Replace /path/to/root/ with the path to the root directory of your HTML files and index.html with the name of the input file you want to parse.
The program will automatically create an output.html file in the /path/to/root/ directory, with any iframes embedded in the index.html file.
