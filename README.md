# censusgo

**This package is unofficial and is NOT in any way supported by Daybreak Games.**

All use of game data and other content must comply with [Daybreak Games Intellectual Property policy](http://census.daybreakgames.com/#intellectual-property)
and Daybreak Game Company's [Terms of Service](https://www.daybreakgames.com/termsofservice.vm?locale=en_US).

---

censusgo is a library that makes it much easier to use the Census API provided by Daybreak Games Company.
It uses dot notation chaining to make it easy to define and maintain queries.

For example, the following code will return all character faction stats saved after `lastLoginTime` and retrieve them
from the API in chunks of 500 at a time.

```go
b := censusgo.NewQueryBuilder("example", "ps2")

q := b.NewQuery("characters_stat_by_faction")
q.SetLimit(500)
q.Where("character_id").Equals("5428010917249639985")
q.ShowFields(
    "character_id",
    "stat_name",
    "profile_id",
    "value_forever_vs",
    "value_forever_nc",
    "value_forever_tr")
q.Where("last_save_date").IsGreaterThanOrEquals(lastLoginTime)

results, err := q.GetResultsBatch()
```

For a working example check out the [examples](https://github.com/Lampjaw/censusgo/examples)!

## Table of Contents

- [Installing](#installing)
- [Creating a new query](#creating-a-new-query)
- [Returning data from the query](#returning-data-from-the-query)
- [Defining a condition](#defining-a-condition)
- [Setting a language](#setting-a-language)
- [Show certain fields](#show-certain-fields)
- [Hide certain fields](#hide-certain-fields)
- [Set a limit for number of rows to return](#set-a-limit-for-number-of-rows-to-return)
- [Set the starting row](#set-the-starting-row)
- [Add a resolve](#add-a-resolve)
- [Join to another service](#join-to-another-service)
- [Tree results on a field](#tree-results-on-a-field)
- [Getting the url of the query](#getting-the-url-of-the-query)
- [Streaming data](#streaming-data)

### Installing

To download the censusgo library use

```go
go get github.com/Lampjaw/censusgo
```

### Creating a new query

To start creating queries, first create a `QueryBuilder`. `QueryBuilder` is a factory for generating queries that share
the same service id and namespace. Using `QueryBuilder` you can call `NewQuery(<collection_name>)` to begin building your
query definition.

```go
builder := censusgo.NewQueryBuilder("example", "ps2")
query := builder.NewQuery("character")
```

### Returning data from the query

There are two methods of query resolution: `GetResults()` and `GetResultsBatch()`.
Both methods return type `[]interface{}` but execute the query differently. 

`GetResults()` will only return the number of results or less of the limit specified or a single result if no limit is set.

`GetResultsBatch()` will attempt to return **every** record in the collection until either an error is thrown or the number
of records returns is less than the limit. If no limit is specified it defaults to 500 records per fetch attempt.

```go
results, err := query.GetResults()
```

### Defining a condition

```go
query.Where("name.lower").Equals("lampjaw")
```

The following operations and their equivalent syntax is below:

* `Equals`: =
* `NotEquals`: =!
* `IsLessThan`: =<
* `IsLessThanOrEquals`: =[
* `IsGreaterThan`: =>
* `IsGreaterThanOrEquals`: =]
* `StartsWith`: =^
* `Contains`: =*

### Setting a language

```go
query.SetLanguageString("en")

OR

query.SetLanguage(censusgo.LangEnglish)
```

No language is set by default so you will receive all localized strings if available.

### Show certain fields

```go
query.ShowFields("character_id"[, "field2", "field3", ...])
```

### Hide certain fields

```go
query.HideFields("currency"[, "field2", "field3", ...])
```

### Set a limit for number of rows to return

```go
query.SetLimit(10)
```

### Set the starting row

```go
query.SetStart(100)
```

### Add a resolve

```go
query.AddResolve("world"[, "resolve2", "resolve3", ...])
```

### Join to another service

```go
worldJoin := query.JoinCollection("characters_world")
```

Join objects have the following methods:

* `IsList(bool)`
* `IsOuterJoin(bool)`
* `ShowFields(...string)`: See the 'Show certain fields' section above
* `HideFields(...string)`: See the 'Hide certain fields' section above
* `OnField(string)`
* `ToField(string)`
* `WithInjectAt(string)`
* `Where(string)`: See the 'Defining a condition' section above
* `JoinCollection(string)`: Returns another join object for sub joining

### Tree results on a field

```go
vehicleTree := query.TreeField("type_id")
```

Tree objects have the following methods:

* `IsList(bool)`
* `GroupPrefix(string)`
* `StartField(string)`
* `TreeField(string)`: Returns another tree object for sub grouping

### Getting the url of the query

```go
url := query.GetUrl()
```