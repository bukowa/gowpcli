package delete

//Delete a meta field.

type Delete struct {
    
    Id string // <id>
    
    Key string // [<key>]
    
    Value string // [<value>]
    
    All bool // [--all]
    
}

//## OPTIONS
//
//<id>
//: The ID of the object.
//
//[<key>]
//: The name of the meta field to delete.
//
//[<value>]
//: The value to delete. If omitted, all rows with key will deleted.
//
//[--all]
//: Delete all meta for the object.
//
//