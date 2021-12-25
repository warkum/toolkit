# GQL CLIENT

how to use 
```
import (
    "log"
    
    "github.com/warkum/toolkit/gqlclient"
)

func main() {
    cfg := gqlclient.Config{
        Address: "https://wkm.com/gql",
        Headers: map[string]string{
            "key":"value",
        }.
    }
    
    client := gqlclient.New(cfg)
    
    request := gqlclient.Request{
        Message: `
        query GetUser($userID: String!) {
          user(id: $userID) {
            uuid
            name
          }
        }`,
        Variables: map[string]interface{}{
            "userID":"xxx",
        },
        Headers: map[string]string{},
    }
    
    // you can also define struct response
    var response map[string]interface{}
    
    err := client.Run(context.Background, request, &response)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println(reponse)
}
```