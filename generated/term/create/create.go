package create

//Creates a new term.

type Create struct {
    
    Taxonomy string // <taxonomy>
    
    Term string // <term>
    
    Slug string // [--slug=<slug>]
    
    Description string // [--description=<description>]
    
    Parent string // [--parent=<term-id>]
    
    Porcelain bool // [--porcelain]
    
}

//## OPTIONS
//
//<taxonomy>
//: Taxonomy for the new term.
//
//<term>
//: A name for the new term.
//
//[--slug=<slug>]
//: A unique slug for the new term. Defaults to sanitized version of name.
//
//[--description=<description>]
//: A description for the new term.
//
//[--parent=<term-id>]
//: A parent for the new term.
//
//[--porcelain]
//: Output just the new term id.
//
//## EXAMPLES
//
//    # Create a new category "Apple" with a description.
//    $ wp term create category Apple --description="A type of fruit"
//    Success: Created category 199.
//
//