![mashable](https://upload.wikimedia.org/wikipedia/commons/6/6d/Mashable.png)

**go client for [mashable](mashable.com)**

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
- Topics
   - [x] List
   - [x] Get
- User
   - [x] Get
- Search
   - [x] Search with Query

## Get The Package
___

```bash
$ go get -u github.com/sadlil/mashable
```