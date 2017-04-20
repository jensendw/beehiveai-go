# beehiveai-go
Simple golang library for BeeHiveAI integration API


## Installation

```
go get github.com/jensendw/beehiveai-go
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/jensendw/beehiveai-go"
)

var	authtoken = "" // Set your auth token here

func main() {
  client := bhive.NewClient(authtoken)
  //Create just the integration ID
  value, err := client.CreateIntegrationID("MyAwesomeIntegration")

  //Create an integration, this will automatically create the ID if it doesn't exist
  value, err := client.CreateIntegration("MyAwesomeIntegration", "text value", "tag1,tag2,tag3")

}
```

## License
[Apache 2](http://www.apache.org/licenses/LICENSE-2.0)

## Contributing

1. Fork it ( https://github.com/jensendw/beehiveai-go )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
