package location

//Assigns, removes, and lists a menu's locations.

type Location struct {
    
}

//## EXAMPLES
//
//    # List available menu locations
//    $ wp menu location list
//    +----------+-------------------+
//    | location | description       |
//    +----------+-------------------+
//    | primary  | Primary Menu      |
//    | social   | Social Links Menu |
//    +----------+-------------------+
//
//    # Assign the 'primary-menu' menu to the 'primary' location
//    $ wp menu location assign primary-menu primary
//    Success: Assigned location to menu.
//
//    # Remove the 'primary-menu' menu from the 'primary' location
//    $ wp menu location remove primary-menu primary
//    Success: Removed location from menu.
//
//
//
//