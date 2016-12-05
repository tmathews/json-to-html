# JSON to HTML

This tool currently takes single JSON objects through standard input delimited
by new lines and outputs, via standard output, an HTML version based on the template provided via
flag.

You currently have to separate the HTML output yourself by looking for the
`</html>` text.

There is no support for partials.

## Todo

 * Support partials.
 * Support multi-line JSON input.
 * Add flag for custom HTML output delimiter.
 * Add flag for HTML output file(s). Add filename generation support. 
 * Add flag for JSON input file(s). 
