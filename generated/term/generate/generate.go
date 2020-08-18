package generate

//Generates some terms.

type Generate struct {
    
    Taxonomy string // <taxonomy>
    
    Count string // [--count=<number>]
    
    MaxDepth string // [--max_depth=<number>]
    
    Format string // [--format=<format>]
    
}

//Creates a specified number of new terms with dummy data.
//
//## OPTIONS
//
//<taxonomy>
//: The taxonomy for the generated terms.
//
//[--count=<number>]
//: How many terms to generate?
//---
//default: 100
//---
//
//[--max_depth=<number>]
//: Generate child terms down to a certain depth.
//---
//default: 1
//---
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: progress
//options:
//  - progress
//  - ids
//---
//
//## EXAMPLES
//
//    # Generate post categories.
//    $ wp term generate category --count=10
//    Generating terms  100% [=========] 0:02 / 0:02
//
//    # Add meta to every generated term.
//    $ wp term generate category --format=ids --count=3 | xargs -d ' ' -I % wp term meta add % foo bar
//    Success: Added custom field.
//    Success: Added custom field.
//    Success: Added custom field.
//
//