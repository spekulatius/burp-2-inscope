package main

import (
    "fmt"
    // "bufio"
    "log"
    "os"
    "io"
    "encoding/json"
)

type Entry struct {
    Enabled  bool   `json:"enabled"`
    File     string `json:"file"`
    Host     string `json:"host"`
    Port     string `json:"port"`
    Protocol string `json:"protocol"`
}

type BurpConfig struct {
    Target struct {
        Scope struct {
            AdvancedMode bool `json:"advanced_mode"`
            Exclude []Entry `json:"exclude"`
            Include []Entry `json:"include"`
        }
    }
}

func main() {
    // Get all variables we need to work.
    var filename string = os.Args[1]
    // var url string = os.Args[2]

    // Read the file
    content, err := os.ReadFile(filename)

    if err != nil {
        log.Fatal(err)
    }

    // Parse the JSON
    var burpConfig BurpConfig
    err = json.Unmarshal(content, &burpConfig)

    if err != nil {
        fmt.Println(err)
    }


    // fmt.Println("BurpConfig is:", burpConfig.Target.Scope)

    // info, _ := os.Stdin.Stat()
    r, w := io.Pipe()
fmt.Println(r)
w.Println("test");
    // reader := bufio.NewReader(os.Stdin)
    // reader.Read(data)

    var matched bool = false;
    for i := 0; i < len(burpConfig.Target.Scope.Include); i++ {
        // var entry Entry = burpConfig.Target.Scope.Include[i];

        // Check the protocol


        // Check the port, if set.


        // Check the regex


    }

    // Output the URL for further processing if it was matched.
    if matched {
        // fmt.Println(url)
    }
}
