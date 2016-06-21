![mashable](https://upload.wikimedia.org/wikipedia/commons/6/6d/Mashable.png)

**go client for [mashable](http://mashable.com)**
___

Someone on this community said there is no possible usecase of this library. I agree.
Probably not. Probably this is a library that will be a use to no one. Doesn't matter.
It was usefull to me. 
___

### What is mashable
___

Mashable was created by `Pete Cashmore` in July 2005. Time noted Mashable as one of the 25 
best blogs in 2009, and has been described as `one stop shop` for social media.
As of November 2015, it has over 6,000,000 Twitter followers and over 3,200,000 fans on Facebook.

### Background
___

While working on a research paper about data mining with one of my friend, i needed to parse
a lot (i really mean a lot) of mashable data as my test dataset, to extract the data i desires.
I was trying to write a web parser first, while i discovered the mashable api as far they have
and saves me lots of time. Instead of making a web parser i then write the api clients for 
mashable.

### Supported Resource
___

The api is undocumented, so the exact formats or some api endpoints can be missed. In case of
scenario like, i will be happy to get suggestions or PRs.

The supported resource list by the client is follows
- Post
   - [x] List
   - [x] Get
   - [x] Get From Url
   - [ ] Comments
- Topics
   - [x] List
   - [x] Get
- User
   - [x] Get
- Search
   - [x] Search with Query

If more api resources are found, will be added here.

### Get The Package
___

```bash
$ go get -u github.com/sadlil/mashable
```

### Usage
___

**Create Client:**
```go
import (
   "github.com/sadlil/mashable"
)

func main() {
   c, err := mashable.New()
   if err != nil {
      panic("client creation failed")
   }
}
```

**Get Posts:**
```go
// List returns the list of post order by post time. `ListOptions` contains two configurable
// field `page` and `per_page`. Setting `page` will return the specified page of the list. according
// `per_page`. Default values is 1 for the page and 20 for the per_page.
posts, err := c.Posts().List(&mashable.ListOptions{})

// Get a specified post by its id
post, err := c.Posts().Get(id)


// Get a specified post by its blog url. usefull if you do not know the id
post, err := c.Posts().GetFromUrl(url)

```

**Get Topics:**
```go
// List returns the list of topic. `ListOptions` contains two configurable
// field `page` and `per_page`. Setting `page` will return the specified page of the list. according
// `per_page`. Default values is 1 for the page and 20 for the per_page.
topics, err := c.Topics().List(&mashable.ListOptions{})

// Get a specified topic by its slug
topic, err := c.Topics().Get(slug)
```

**Get User:**
```go
// Get a specified user by his/her id
user, err := c.Users().Get(id)
```

**Search:**
```go
// Get search result according to a search query
res, err := c.Search().Query(query)
```

Details of all the response types can be found in [types.go](types.go) file.
Check tests for more examples.


### License
___

Licensed under the MIT license. See the [LICENSE](LICENSE) file for details.
