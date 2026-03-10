# Financial Control

Financial Control is a multi-project repository for managing personal financial data through backend services and a web interface.

## About The Product

Financial Control helps users organize cash flow and card spending in one place.

It is built around these core product capabilities:

- Transaction management: create, read, update, delete, and mark transactions as paid.
- Transaction types: track `income`, `debit`, and `credit` operations.
- Categories: organize transactions by custom category, icon, and transaction type.
- Credit cards: register cards with billing cycle data (close day, expire day) and visual properties.
- Recurrence support:
	- monthly recurring transactions (fixed day each month),
	- annual recurring transactions (fixed month/day each year),
	- installment transactions (between initial and final dates).
- Reporting: generate monthly balance data including total income, debit, credit, and final balance.
- Spending analysis: aggregate spending by category and by credit card for monthly insights.

## Repository Structure

- `engine/`: Go service that contains core domain logic, data access, migrations, and API endpoints.
- `bff/`: Backend-for-Frontend layer (currently in early setup).
- `auto-service/`: Automation/service integration layer (currently in early setup).
- `webapp/`: Frontend web application (currently in early setup).