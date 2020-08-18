/*
## INFO
	
	
 */
package config


// Config //Generates and reads the wp-config.php file.
type Config struct {
}

func (c Config) Args() []string {
    var args = []string{"config"}
    return args
}

