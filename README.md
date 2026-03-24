# Portfolio Rebalancer

A simple portfolio rebalancing engine implemented in Go.

This project calculates which stocks should be **bought or sold** in order to match a desired target allocation.

---

## Overview

A portfolio consists of:

- Current holdings (stocks with quantity and price)
- Target allocation (desired percentage distribution)

The system:

1. Calculates total portfolio value
2. Determines current allocation
3. Compares against target allocation
4. Generates actions: **BUY / SELL / HOLD**

---

## Example

### Current Portfolio

| Stock | Value |
| ----- | ----- |
| META  | 600   |
| AAPL  | 400   |

**Total = 1000**

---

### Target Allocation

| Stock | %   |
| ----- | --- |
| META  | 40% |
| AAPL  | 60% |

---

### Result

| Stock | Action | Amount |
| ----- | ------ | ------ |
| META  | SELL   | 200    |
| AAPL  | BUY    | 200    |

---

## Architecture

The project follows a simple layered structure:

```
cmd/            → entry point (main)
internal/
  domain/       → core business logic
  application/  → use case orchestration
  types/        → shared models
```

### Key Design Decisions

- Domain logic is isolated from application logic
- Rebalancing is deterministic and pure
- No external dependencies (kept intentionally simple)
- Action types use integers (DB and i18n friendly)

---

## Getting Started

### Requirements

- Go 1.20+

### Run the project

```bash
go run cmd/main.go
```

---

### Run tests

```bash
go test ./...
```

---

## Core Logic

For each stock:

```
target_value = total_portfolio_value * target_percentage
difference = target_value - current_value
```

- If difference > 0 → BUY
- If difference < 0 → SELL
- If difference = 0 → HOLD

---

## Edge Cases Considered

- Stocks present in holdings but not in target → SELL all
- Stocks present in target but not in holdings → BUY from zero
- Empty portfolio handling

---

## LLM Usage Disclosure

This project was developed with the assistance of an LLM for:

- validating architectural decisions in Golang constraints
- Identifying edge cases

All implementation and final decisions were reviewed and written manually.

---

## Documentantion

Additional development notes, including design decision and LLM usage, are available in:

```
 docs/development-notes.pdf
```

---

## Future Improvements

(Not implemented to avoid overengineering)

- Transaction fees handling
- Minimum trade thresholds
- Real-time price integration
- API interface (REST)
- AI-driven allocation (external system)

---

## Author

Jhoan Gutiérrez  
Software Engineer (Backend & Systems Focus)
