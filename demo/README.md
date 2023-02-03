Example Queries

[Affordable Italian wine](https://link.weaviate.io/3XTelA5)

```
{
  Get{
    Review(
      nearText:{
        concepts:["an affordable wine representative of Italy"]
      },
      limit:1
    ){
      name
      review
      varietal
      vintage
    }
  }
}
```

[When are wines ready to drink?](https://console.semi.technology/console/query#weaviate_uri=http://localhost:8080&graphql_query=%7B%0A%20%20Get%7B%0A%20%20%20%20Review%7B%0A%20%20%20%20%20%20_additional%7B%0A%20%20%20%20%20%20%20%20generate(%0A%20%20%20%20%20%20%20%20%20%20singleResult%3A%7B%0A%20%20%20%20%20%20%20%20%20%20%20%20prompt%3A%20%22%22%22%0A%20%20%20%20%20%20%20%20%20%20%20%20based%20on%20%7Breview%7D%2C%0A%20%20%20%20%20%20%20%20%20%20%20%20when%20is%20this%20wine%20ready%20to%20drink%3F%0A%20%20%20%20%20%20%20%20%20%20%20%20%22%22%22%0A%20%20%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%20%20)%20%7B%0A%20%20%20%20%20%20%20%20%20%20singleResult%0A%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20name%0A%20%20%20%20%20%20review%0A%20%20%20%20%20%20varietal%0A%20%20%20%20%20%20vintage%0A%20%20%20%20%7D%0A%20%20%7D%0A%7D%0A)

```
{
  Get{
    Review{
      _additional{
        generate(
          singleResult:{
            prompt: """
            based on {review},
            when is this wine ready to drink?
            """
          }
        ) {
          singleResult
        }
      }
      name
      review
      varietal
      vintage
    }
  }
}
```

[An opinion on the style of wine](https://console.semi.technology/console/query#weaviate_uri=http://localhost:8080&graphql_query=%7B%0A%20%20Get%7B%0A%20%20%20%20Review(%0A%20%20%20%20%20%20nearText%3A%7B%0A%20%20%20%20%20%20%20%20concepts%3A%22an%20aged%20classic%20riesling%22%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20limit%3A1%0A%20%20%20%20)%7B%0A%20%20%20%20%20%20_additional%7B%0A%20%20%20%20%20%20%20%20generate(%0A%20%20%20%20%20%20%20%20%20%20singleResult%3A%7B%0A%20%20%20%20%20%20%20%20%20%20%20%20prompt%3A%20%22%22%22%0A%20%20%20%20%20%20%20%20%20%20%20%20based%20on%20%7Breview%7D%20would%20you%20consider%20this%20a%20fruit%20bomb%3F%0A%20%20%20%20%20%20%20%20%20%20%20%20%22%22%22%0A%20%20%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%20%20)%20%7B%0A%20%20%20%20%20%20%20%20%20%20singleResult%0A%20%20%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%20%20name%0A%20%20%20%20%20%20review%0A%20%20%20%20%20%20varietal%0A%20%20%20%20%20%20vintage%0A%20%20%20%20%7D%0A%20%20%7D%0A%7D%0A)

```
{
  Get{
    Review(
      nearText:{
        concepts:"an aged classic riesling"
      }
      limit:1
    ){
      _additional{
        generate(
          singleResult:{
            prompt: """
            based on {review} would you consider this a fruit bomb?
            """
          }
        ) {
          singleResult
        }
      }
      name
      review
      varietal
      vintage
    }
  }
}
```
