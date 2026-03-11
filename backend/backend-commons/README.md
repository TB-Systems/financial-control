# Financial Control Backend Commons

Shared domain contracts for Financial Control backend services.

This module centralizes:

- request/response DTOs (`dtos`)
- domain models used by services and repositories (`models`)
- mapping helpers between models and DTOs (`modelsdto`)
- cross-domain constants (`constants`)

## Module

```go
module backend-commons
```

Inside the monorepo, this module is resolved through `go.work`.

## Structure

- `constants`: common keys and shared strings used by DTO validation and flows.
- `dtos`: API contracts and request validation.
- `models`: internal backend entities.
- `modelsdto`: conversion helpers between `models` and `dtos`.

## DTOs (`dtos`)

Main request and response contracts:

- `CategoryRequest`, `CategoryResponse`
- `CreditCardRequest`, `CreditCardResponse`
- `TransactionRequest`, `TransactionRequestFromRecurrentTransaction`, `TransactionResponse`
- `MonthlyTransactionRequest`, `MonthlyTransactionResponse`
- `AnnualTransactionRequest`, `AnnualTransactionResponse`
- `InstallmentTransactionRequest`, `InstallmentTransactionResponse`
- `MonthlyReportResponse`

Validation methods are implemented directly in request DTOs and return `[]errors.ApiErrorItem` from `github.com/TB-Systems/go-commons/errors`.

Implemented validators:

- `CategoryRequest.Validate()`
- `CreditCardRequest.Validate()`
- `TransactionRequestFromRecurrentTransaction.Validate()`
- `TransactionRequest.Validate()`
- `MonthlyTransactionRequest.Validate()`
- `AnnualTransactionRequest.Validate()`
- `InstallmentTransactionRequest.Validate()`

## Models (`models`)

Domain entities include:

- `Category`, `CreateCategory`
- `CreditCard`, `CreateCreditCard`
- `Transaction`, `CreateTransaction`, `ShortTransaction`
- `MonthlyTransaction`, `CreateMonthlyTransaction`, `ShortMonthlyTransaction`
- `AnnualTransaction`, `CreateAnnualTransaction`, `ShortAnnualTransaction`
- `InstallmentTransaction`, `CreateInstallmentTransaction`, `ShortInstallmentTransaction`
- `MonthlyReport`, `CategoriesSpending`, `CreditCardsSpending`

`TransactionType` is defined as:

- `Income`
- `Debit`
- `Credit`

With helper:

- `TransactionType.IsValid()`

## Mappers (`modelsdto`)

Package `modelsdto` provides mapping helpers such as:

- create model from request:
	- `CreateCategoryFromRequest`
	- `CreateCreditCardFromCreditCardRequest`
	- `CreateTransactionFromTransactionRequest`
	- `CreateMonthlyTransactionFromRequest`
	- `CreateAnnualTransactionFromRequest`
	- `CreateInstallmentTransactionFromRequest`
- map model to response:
	- `CategoryResponseFromModel`
	- `CreditCardResponseFromCreditCard`
	- `TransactionResponseFromTransaction`
	- `MonthlyTransactionResponseFromModel`
	- `AnnualTransactionResponseFromModel`
	- `InstallmentTransactionResponseFromModel`
	- `MonthlyReportResponseFromModels`

## Usage Example

```go
package example

import (
		"backend-commons/dtos"
		"backend-commons/modelsdto"

		"github.com/google/uuid"
)

func buildCreateModel(req dtos.CategoryRequest, userID uuid.UUID) {
		_ = modelsdto.CreateCategoryFromRequest(req, userID)
}
```

## Development

Run tests:

```bash
go test ./...
```

Run coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Monorepo Notes

When used together with `engine`, keep both modules in the same workspace using `go.work`:

```txt
use (
		./backend-commons
		./engine
)
```

This allows `engine` imports such as `backend-commons/dtos` and `backend-commons/models` without publishing this module first.