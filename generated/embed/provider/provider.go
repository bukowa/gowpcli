/*
## INFO
	
	
 */
package provider


// Provider //Retrieves oEmbed providers.
type Provider struct {
}

func (p Provider) Args() []string {
    var args = []string{"embed", "provider"}
    return args
}

