# go-sql-builder

[![Build Status](https://travis-ci.org/lukewilliamboswell/go-sql-builder.svg)](https://travis-ci.org/lukewilliamboswell/go-sql-builder) [![Coverage Status](https://coveralls.io/repos/lukewilliamboswell/go-sql-builder/badge.svg?branch=master)](https://coveralls.io/r/lukewilliamboswell/go-sql-builder?branch=master)

Compose SQL query strings

### Install

    go get github.com/lukewilliamboswell/go-sql-builder

### Tests

    go test

### Example
```go
package main

import (
	"fmt"
	"github.com/lukewilliamboswell/go-sql-builder"
)

func main() {

	AND_fields := make([]string, 0)

	AND_fields = append(AND_fields, sqlbuilder.EQUAL_STRING(`Apple`))
	AND_fields = append(AND_fields, sqlbuilder.EQUAL_INT(`OwnerID`))
	AND_fields = append(AND_fields, sqlbuilder.LIKE(`Peaches`))
	AND_fields = append(AND_fields, sqlbuilder.IN_INT(`FruitID`, 1, 2, 3))

	query := sqlbuilder.COMPOSE(
		sqlbuilder.SELECT_FROM(`Fruit`, []string{`Apples`, `Peaches`, `OwnerID`}...),
		` WHERE `,
		sqlbuilder.AND(AND_fields...),
		sqlbuilder.ORDER_BY(1, 250, `FruitID`),
	)

	fmt.Println(query)
	// SELECT [Apples],[Peaches],[OwnerID] FROM [dbo].[Fruit] 
	// WHERE  [Apple] = ? AND [OwnerID] = ?nnn AND [Peaches] LIKE ? AND [FruitID] IN(1,2,3)  
	// ORDER BY [FruitID] OFFSET 1 ROWS FETCH NEXT 250 ROWS ONLY 
}
```
