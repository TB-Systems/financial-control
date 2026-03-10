## Financial Control Engine

### What is?
The Financial Control Engine is an internal Go service that implements the business logic and data layer for financial control operations. It centralizes domain rules, persistence (DTOs / repositories), and migrations so other internal components can interact with a single, consistent source of truth for financial data.

Key points
- Purpose: encapsulate business logic, data models and persistence for all financial control concerns.
- Audience: internal services and authenticated clients inside the platform (not public-facing).
- Responsibilities: domain validation, transactional operations, DB access, DTOs, schema migrations and audit logging.
- Not responsible for authentication/authorization: the service expects requests to arrive already authenticated and to include the user identity (e.g., user id in a validated token). It will *
read
* the user id from the request context/token but does not validate credentials itself.

### Instalation & development build

#### 1 - create a .env file like the example.env

```env
FINANCIAL_CONTROL_DATABASE_PORT=5432
FINANCIAL_CONTROL_DATABASE_NAME=financialcontrol
FINANCIAL_CONTROL_DATABASE_USER=docker
FINANCIAL_CONTROL_DATABASE_PASSWORD=docker
FINANCIAL_CONTROL_DATABASE_HOST=localhost
FINANCIAL_CONTROL_CSRF_KEY=NhkEXjyS5ms3k7vNQ5fbk2Ffv0OIuQs6
```

#### 2 - run the docker compose image to create the database

```shell
docker compose up -d
```

#### 3 - install the dependencies

```shell
go mod download
go mod tidy
```

#### 4 - run the /terndotenv main.go to run the migrations using [tern](https://github.com/jackc/tern)

```shell
go run ./cmd/terndotenv
```

#### 4.1 - run the command below to rollback migrations in linux
```shell
 export $(cat example.env | xargs) && tern migrate --migrations ./internal/store/pgstore/migrations --config ./internal/store/pgstore/migrations/tern.conf --destination=0
```

#### 5 - run the command below for dev mode using [air](https://github.com/air-verse/air)

```shell
air
```

### Endpoints

All routes are prefixed with:

`/engine/v1`

Authentication / user scope:
- Every endpoint requires the `user_id` cookie with a valid UUID.
- If cookie is missing or invalid, the API returns `401` with the error response format.

## Request / Response Conventions

### Common Error Response

```json
{
	"status": 400,
	"messages": ["INVALID_DATA"]
}
```

### Common Success Response (delete/pay operations)

```json
{
	"message": "success"
}
```

### List Response Shape

```json
{
	"items": [],
	"total": 0
}
```

### Paginated Response Shape

```json
{
	"items": [],
	"page_count": 1,
	"page": 1
}
```

### Common Query Parameters

- `limit` (optional): defaults to `10`, max effective value is `10`.
- `page` (optional): defaults to `1`.
- `month` and `year`: required in monthly report endpoints (`month` in `1..12`, `year >= 1970`).
- `start_date` and `end_date` (`YYYY-MM-DD`): optional date filters for transaction listing.

## Endpoint Matrix

### Categories

| Method | Path | Body | Success Response |
|---|---|---|---|
| `POST` | `/categories/` | `CategoryRequest` | `201` + `CategoryResponse` |
| `GET` | `/categories/` | - | `200` + `ResponseList<CategoryResponse>` |
| `GET` | `/categories/:id` | - | `200` + `CategoryResponse` |
| `PUT` | `/categories/:id` | `CategoryRequest` | `200` + `CategoryResponse` |
| `DELETE` | `/categories/:id` | - | `200` + `ResponseSuccess` |

### Credit Cards

| Method | Path | Body | Success Response |
|---|---|---|---|
| `POST` | `/creditcards/` | `CreditCardRequest` | `201` + `CreditCardResponse` |
| `GET` | `/creditcards/` | - | `200` + `ResponseList<CreditCardResponse>` |
| `GET` | `/creditcards/:id` | - | `200` + `CreditCardResponse` |
| `PUT` | `/creditcards/:id` | `CreditCardRequest` | `200` + `CreditCardResponse` |
| `DELETE` | `/creditcards/:id` | - | `200` + `ResponseSuccess` |

### Transactions

| Method | Path | Body | Success Response |
|---|---|---|---|
| `POST` | `/transactions/` | `TransactionRequest` | `201` + `TransactionResponse` |
| `POST` | `/transactions/monthly` | `TransactionRequestFromRecurrentTransaction` | `201` + `TransactionResponse` |
| `POST` | `/transactions/annual` | `TransactionRequestFromRecurrentTransaction` | `201` + `TransactionResponse` |
| `POST` | `/transactions/installment` | `TransactionRequestFromRecurrentTransaction` | `201` + `TransactionResponse` |
| `GET` | `/transactions/` | - | `200` + `PaginatedResponse<TransactionResponse>` |
| `GET` | `/transactions/:id` | - | `200` + `TransactionResponse` |
| `GET` | `/transactions/report?month={m}&year={y}` | - | `200` + `PaginatedResponse<TransactionResponse>` |
| `PUT` | `/transactions/:id` | `TransactionRequest` | `200` + `TransactionResponse` |
| `DELETE` | `/transactions/:id` | - | `200` + `ResponseSuccess` |
| `PUT` | `/transactions/pay/:id` | - | `200` + `ResponseSuccess` |

Notes for `GET /transactions/`:
- Supports `?page=&limit=`.
- Supports optional date range filter with `?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`.

### Monthly Transactions

| Method | Path | Body | Success Response |
|---|---|---|---|
| `POST` | `/monthly_transactions/` | `MonthlyTransactionRequest` | `201` + `MonthlyTransactionResponse` |
| `GET` | `/monthly_transactions/` | - | `200` + `PaginatedResponse<MonthlyTransactionResponse>` |
| `GET` | `/monthly_transactions/:id` | - | `200` + `MonthlyTransactionResponse` |
| `PUT` | `/monthly_transactions/:id` | `MonthlyTransactionRequest` | `200` + `MonthlyTransactionResponse` |
| `DELETE` | `/monthly_transactions/:id` | - | `204` + `ResponseSuccess` |

### Annual Transactions

| Method | Path | Body | Success Response |
|---|---|---|---|
| `POST` | `/annual_transactions/` | `AnnualTransactionRequest` | `201` + `AnnualTransactionResponse` |
| `GET` | `/annual_transactions/` | - | `200` + `PaginatedResponse<AnnualTransactionResponse>` |
| `GET` | `/annual_transactions/:id` | - | `200` + `AnnualTransactionResponse` |
| `PUT` | `/annual_transactions/:id` | `AnnualTransactionRequest` | `200` + `AnnualTransactionResponse` |
| `DELETE` | `/annual_transactions/:id` | - | `204` + `ResponseSuccess` |

### Installment Transactions

| Method | Path | Body | Success Response |
|---|---|---|---|
| `POST` | `/installment_transactions/` | `InstallmentTransactionRequest` | `201` + `InstallmentTransactionResponse` |
| `GET` | `/installment_transactions/` | - | `200` + `PaginatedResponse<InstallmentTransactionResponse>` |
| `GET` | `/installment_transactions/:id` | - | `200` + `InstallmentTransactionResponse` |
| `PUT` | `/installment_transactions/:id` | `InstallmentTransactionRequest` | `200` + `InstallmentTransactionResponse` |
| `DELETE` | `/installment_transactions/:id` | - | `204` + `ResponseSuccess` |

### Monthly Report

| Method | Path | Body | Success Response |
|---|---|---|---|
| `GET` | `/monthly_report/?month={m}&year={y}` | - | `200` + `MonthlyReportResponse` |

## Body Schemas

### CategoryRequest

```json
{
	"transaction_type": 0,
	"name": "Salary",
	"icon": "wallet"
}
```

`transaction_type` values:
- `0` = income
- `1` = debit
- `2` = credit

### CreditCardRequest

```json
{
	"name": "Main Card",
	"first_four_numbers": "1234",
	"limit": 3000,
	"close_day": 10,
	"expire_day": 20,
	"background_color": "#1E3A8A",
	"text_color": "#FFFFFF"
}
```

### TransactionRequest

```json
{
	"name": "Groceries",
	"date": "2026-03-10T00:00:00Z",
	"value": 250.75,
	"paid": false,
	"category_id": "11111111-1111-1111-1111-111111111111",
	"creditcard_id": "22222222-2222-2222-2222-222222222222",
	"monthly_transaction_id": null,
	"annual_transaction_id": null,
	"installment_transaction_id": null
}
```

### TransactionRequestFromRecurrentTransaction

```json
{
	"id": "33333333-3333-3333-3333-333333333333"
}
```

### MonthlyTransactionRequest

```json
{
	"name": "Internet",
	"value": 120,
	"day": 5,
	"category_id": "11111111-1111-1111-1111-111111111111",
	"creditcard_id": null
}
```

### AnnualTransactionRequest

```json
{
	"name": "Car Tax",
	"value": 800,
	"day": 12,
	"month": 3,
	"category_id": "11111111-1111-1111-1111-111111111111",
	"creditcard_id": null
}
```

### InstallmentTransactionRequest

```json
{
	"name": "Laptop",
	"value": 4500,
	"initial_date": "2026-03-10T00:00:00Z",
	"final_date": "2026-08-10T00:00:00Z",
	"category_id": "11111111-1111-1111-1111-111111111111",
	"creditcard_id": "22222222-2222-2222-2222-222222222222"
}
```

## Response Schemas

### CategoryResponse

```json
{
	"id": "11111111-1111-1111-1111-111111111111",
	"transaction_type": 1,
	"name": "Food",
	"icon": "utensils",
	"created_at": "2026-03-10T12:00:00Z",
	"updated_at": "2026-03-10T12:00:00Z"
}
```

### CreditCardResponse

```json
{
	"id": "22222222-2222-2222-2222-222222222222",
	"name": "Main Card",
	"first_four_numbers": "1234",
	"limit": 3000,
	"close_day": 10,
	"expire_day": 20,
	"background_color": "#1E3A8A",
	"text_color": "#FFFFFF",
	"created_at": "2026-03-10T12:00:00Z",
	"updated_at": "2026-03-10T12:00:00Z"
}
```

### TransactionResponse

```json
{
	"id": "44444444-4444-4444-4444-444444444444",
	"name": "Groceries",
	"date": "2026-03-10T00:00:00Z",
	"value": 250.75,
	"paid": false,
	"category": {
		"id": "11111111-1111-1111-1111-111111111111",
		"transaction_type": 1,
		"name": "Food",
		"icon": "utensils",
		"created_at": "2026-03-10T12:00:00Z",
		"updated_at": "2026-03-10T12:00:00Z"
	},
	"creditcard": {
		"id": "22222222-2222-2222-2222-222222222222",
		"name": "Main Card",
		"first_four_numbers": "1234",
		"limit": 3000,
		"close_day": 10,
		"expire_day": 20,
		"background_color": "#1E3A8A",
		"text_color": "#FFFFFF",
		"created_at": "2026-03-10T12:00:00Z",
		"updated_at": "2026-03-10T12:00:00Z"
	},
	"monthly_transaction": {
		"id": "55555555-5555-5555-5555-555555555555",
		"name": "Internet",
		"value": 120,
		"day": 5,
		"created_at": "2026-03-10T12:00:00Z",
		"updated_at": "2026-03-10T12:00:00Z"
	},
	"annual_transaction": null,
	"installment_transaction": null,
	"created_at": "2026-03-10T12:00:00Z",
	"updated_at": "2026-03-10T12:00:00Z"
}
```

### MonthlyTransactionResponse

```json
{
	"id": "55555555-5555-5555-5555-555555555555",
	"name": "Internet",
	"value": 120,
	"day": 5,
	"category": {
		"id": "11111111-1111-1111-1111-111111111111",
		"transaction_type": 1,
		"name": "Utilities",
		"icon": "bolt",
		"created_at": "2026-03-10T12:00:00Z",
		"updated_at": "2026-03-10T12:00:00Z"
	},
	"creditcard": null,
	"created_at": "2026-03-10T12:00:00Z",
	"updated_at": "2026-03-10T12:00:00Z"
}
```

### AnnualTransactionResponse

```json
{
	"id": "66666666-6666-6666-6666-666666666666",
	"name": "Car Tax",
	"value": 800,
	"day": 12,
	"month": 3,
	"category": {
		"id": "11111111-1111-1111-1111-111111111111",
		"transaction_type": 1,
		"name": "Taxes",
		"icon": "receipt",
		"created_at": "2026-03-10T12:00:00Z",
		"updated_at": "2026-03-10T12:00:00Z"
	},
	"creditcard": null,
	"created_at": "2026-03-10T12:00:00Z",
	"updated_at": "2026-03-10T12:00:00Z"
}
```

### InstallmentTransactionResponse

```json
{
	"id": "77777777-7777-7777-7777-777777777777",
	"name": "Laptop",
	"value": 4500,
	"initial_date": "2026-03-10T00:00:00Z",
	"final_date": "2026-08-10T00:00:00Z",
	"category": {
		"id": "11111111-1111-1111-1111-111111111111",
		"transaction_type": 2,
		"name": "Electronics",
		"icon": "laptop",
		"created_at": "2026-03-10T12:00:00Z",
		"updated_at": "2026-03-10T12:00:00Z"
	},
	"creditcard": {
		"id": "22222222-2222-2222-2222-222222222222",
		"name": "Main Card",
		"first_four_numbers": "1234",
		"limit": 3000,
		"close_day": 10,
		"expire_day": 20,
		"background_color": "#1E3A8A",
		"text_color": "#FFFFFF",
		"created_at": "2026-03-10T12:00:00Z",
		"updated_at": "2026-03-10T12:00:00Z"
	},
	"created_at": "2026-03-10T12:00:00Z",
	"updated_at": "2026-03-10T12:00:00Z"
}
```

### MonthlyReportResponse

```json
{
	"total_income": 5000,
	"total_debit": 1800,
	"total_credit": 1200,
	"balance": 2000,
	"most_spent_category": {
		"id": "11111111-1111-1111-1111-111111111111",
		"name": "Food",
		"icon": "utensils",
		"transaction_type": 1,
		"value": 600
	},
	"most_spent_creditcard": {
		"id": "22222222-2222-2222-2222-222222222222",
		"name": "Main Card",
		"first_four_numbers": "1234",
		"limit": 3000,
		"close_day": 10,
		"expire_day": 20,
		"background_color": "#1E3A8A",
		"text_color": "#FFFFFF",
		"total_spent": 900
	},
	"categories": [],
	"creditcards": []
}
```