# LinkedIn Automation – Technical Proof of Concept

## ⚠️ Disclaimer
This project is strictly for **educational and technical evaluation purposes**.
Automating LinkedIn violates their Terms of Service.
This code **must not** be used in production or on real accounts.

---

## Overview
This repository contains a **technical proof-of-concept** demonstrating
advanced browser automation techniques using **Go + Rod**.

The focus of this assignment is on:
- clean Go architecture
- human-like interaction patterns
- anti-detection / stealth techniques
- session persistence via browser cookies

For safety and compliance reasons, **only the authentication flow is implemented**.

---

## Implemented Scope

### Authentication (Core Focus)
- Human-like login flow
- Config-driven selectors
- Success & failure detection
- Timeout handling
- Mock login page used for safe demonstration

### Session Persistence
- Browser cookies captured via Chrome DevTools Protocol
- Cookies restored on subsequent runs
- Mirrors real-world browser session reuse logic

### Anti-Bot / Stealth Techniques
- Human-like mouse movement
- Human typing simulation
- Randomized delays & think time
- Random viewport sizing
- Scroll & hesitation behavior

> Advanced features such as search, messaging, and connection requests
> are intentionally excluded to keep this project aligned with
> ethical and evaluation constraints.

---

## Project Structure
```

internal/
auth/        # Generic login engine
browser/     # Browser lifecycle & wiring
config/      # Environment-based configuration
session/     # Cookie persistence
stealth/     # Human-like behavior & anti-detection

````

---

## Setup & Run

```bash
cp .env.example .env
go run main.go
````

---

## Notes

A mock login page is used for demonstration.
Session persistence is implemented at the browser level and reflects
real-world authentication workflows.

---

## Demo

Link : [Demo Link](soon)  
A short demo video is included showing: 

* browser launch
* human-like login interaction
* session persistence logic
