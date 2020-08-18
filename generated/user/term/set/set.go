package set

//Set object terms.

type Set struct {
    
    Id string // <id>
    
    Taxonomy string // <taxonomy>
    
    Term []string // <term>...
    
    By string // [--by=<field>]
    
}

//Replaces existing terms on the object.
//
//<id>
//: The ID of the object.
//
//<taxonomy>
//: The name of the taxonomy type to be updated.
//
//<term>...
//: The slug of the term or terms to be updated.
//
//[--by=<field>]
//: Explicitly handle the term value as a slug or id.
//---
//options:
//  - slug
//  - id
//---
//
//