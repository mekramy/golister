# Golister API Documentation

`golister` is a Go library designed to manage pagination, sorting, filtering, and metadata for data listings. This documentation details the API structure, methods, and usage patterns.

## Installation

```bash
go get github.com/mekramy/golister
```

## Constructor Functions

- **`New(options ...Options) Lister`**: Creates a new `Lister` instance with custom options.
- **`NewFromParams(params ListerParams, options ...Options) Lister`**: Creates a `Lister` from `ListerParams`.
- **`NewFromJson(data string, options ...Options) (Lister, error)`**: Creates a `Lister` from a JSON string.
- **`NewFromBase64Json(data string, options ...Options) (Lister, error)`**: Creates a `Lister` from a Base64-encoded JSON string.

## Types and Methods

### `ListerParams`

**Structure:**

```go
type ListerParams struct {
    Page    uint64         `json:"page"`
    Limit   uint           `json:"limit"`
    Sorts   []Sort         `json:"sorts"`
    Search  string         `json:"search"`
    Filters map[string]any `json:"filters"`
}
```

**Description:** Defines input parameters for creating a `Lister`. Includes pagination, limit, sorting, search, and filters.

### `Lister` Interface

The `Lister` interface provides methods to manage pagination, sorting, and metadata:

### Pagination

- `SetPage(page uint64) Lister`: Sets the current page.
- `Page() uint64`: Retrieves the current page.
- `Pages() uint64`: Returns the total number of pages. Available after `SetTotal` method called.
- `SetLimit(limit uint) Lister`: Sets the maximum number of items per page.
- `Limit() uint`: Retrieves the current limit.

### Sorting

- `AddSort(sort string, order Order) Lister`: Adds a sorting condition.
- `Sort() []Sort`: Returns the active sorting configuration.

### Search

- `SetSearch(search string) Lister`: Sets a search keyword or phrase.
- `Search() string`: Retrieves the search keyword.

### Filters

- `SetFilters(filters map[string]any) Lister`: Assigns a map of key-value pairs to filter data.
- `AddFilter(key string, value any) Lister`: Add a single filter condition.
- `Filters() map[string]any`: Retrieves all applied filters.
- `Filter(key string) any`: Retrieves a specific filter value.
- `HasFilter(key string) bool`: Checks if a specific filter exists.

### Metadata

- `AddMeta(key string, value any) Lister`: Sets metadata.
- `Meta(key string) any`: Retrieves a metadata value.
- `HasMeta(key string) bool`: Checks if metadata exists.
- `MetaData() map[string]any`: Retrieves the full metadata set.

### Record Counting

- `SetTotal(total uint64) Lister`: Sets the total number of records. This method must called after all sets.
- `Total() uint64`: Retrieves the total number of records.
- `From() uint64`: Calculates the starting position for records on the current page. Available after `SetTotal` method called.
- `To() uint64`: Calculates the ending position for records on the current page. Available after `SetTotal` method called.

### Responses

- `SQLSortOrder() string`: Generates an SQL-compatible ORDER BY clause.
- `Response() map[string]any`: Builds a structured JSON response.
- `ResponseWithData(data any) map[string]any`: Generates a JSON response with pagination and metadata details combined with additional data.

### Options

- **`WithLimits(def uint, valids ...uint) Options`**: Sets the default limit and valid limits.
- **`WithSorts(def string, valids ...string) Options`**: Sets the default sort and valid sorts.
- **`WithDefaultLimit(def uint) Options`**: Sets a default limit.
- **`WithDefaultSort(def string) Options`**: Sets a default sort.
- **`WithMySQLSorter() Options`**: Configures the SQL sorter for MySQL.
- **`WithPostgreSQLSorter() Options`**: Configures the SQL sorter for PostgreSQL.

## Example Usage

```go
import "github.com/mekramy/golister"

params := golister.ListerParams{
    Page:   1,
    Limit:  50,
    Search: "John Doe",
    Sorts: []Sort{
        {Field: "name", Order: golister.Ascending},
    }
    Filters: map[string]any{
        "status": "active",
    },
}
lister := golister.NewFromParams(params)
lister.SetTotal(101)
lister.AddMeta("server_time", time.Now().Unix())

response := lister.Response()
fmt.Println(response)
```
