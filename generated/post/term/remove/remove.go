package remove

//Remove a term from an object.

type Remove struct {
    
    Id string // <id>
    
    Taxonomy string // <taxonomy>
    
    Term []string // [<term>...]
    
    By string // [--by=<field>]
    
    All bool // [--all]
    
}

//## OPTIONS
//
//<id>
//: The ID of the object.
//
//<taxonomy>
//: The name of the term's taxonomy.
//
//[<term>...]
//: The name of the term or terms to be removed from the object.
//
//[--by=<field>]
//: Explicitly handle the term value as a slug or id.
//---
//options:
//  - slug
//  - id
//---
//
//[--all]
//: Remove all terms from the object.
//
//