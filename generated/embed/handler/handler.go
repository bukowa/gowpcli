/*
## INFO
	
	
 */
package handler


// Handler //Retrieves embed handlers.
type Handler struct {
}

func (h Handler) Args() []string {
    var args = []string{"embed", "handler"}
    return args
}

