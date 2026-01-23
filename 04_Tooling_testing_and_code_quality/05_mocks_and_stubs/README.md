# Mocks and Stubs

This folder demonstrates using interfaces to provide stubs and mocks for unit tests.

Summary:

- Stubs: simple implementations that return predefined responses. Use for happy-path tests and when you only need controlled inputs.
- Mocks: test doubles that record calls and allow assertions about interactions (call counts, parameters). Use when you need to verify behavior (e.g., that a retry was attempted).

Comparison (quick):

- Ease: Stubs are easiest to write by hand. Mocks take more code but give stronger verification.
- Coupling: Overusing mocks can tightly couple tests to implementation details; prefer stubs when possible.
- Tooling: `gomock` and `testify/mock` automate mock generation; handwritten mocks are often sufficient for small interfaces.

When to use which:

- Use a stub when you only need to control return values for dependency calls.
- Use a mock when your test must assert interactions (calls, order, arguments).
- Favor interface design that keeps dependencies small and focused so stubbing/mocking remains simple.

Helpful resources:

- Roadmap note on TDD and mocks: https://roadmap.sh/
- Article: Mock solutions for Golang unit test (Medium): https://laiyuanyuan-sg.medium.com/mock-solutions-for-golang-unit-test-a2b60bd3e157
- Article: Stubbing vs mocking discussion: https://blog.stackademic.com/test-driven-development-in-golang-stubbing-vs-mocking-vs-not-mocking-5f23f25b3a63
- Guide: Writing unit tests / mocking (Medium): https://medium.com/nerd-for-tech/writing-unit-tests-in-golang-part-2-mocking-d4fa1701a3ae
- Video explainer (recommended): https://www.youtube.com/watch?v=Ir7dl7XX9r4

Run the example tests:

```bash
cd golang_roadmap/04_Tooling_testing_and_code_quality/05_mocks_and_stubs
go test -v
```

Notes:

- For larger projects consider generated mocks (gomock) or `testify/mock` for expressive assertions. Keep tests readable and avoid over-specifying internal call order unless required.

Gomock vs Testify/Mock (summary)

- Approach: `gomock` uses generated, type-safe mocks (`mockgen`). `testify/mock` is reflection-based and more dynamic.
- Safety: `gomock` catches interface changes at compile time; `testify/mock` surfaces issues at test runtime.
- Ease: `testify/mock` is quicker to start with; `gomock` requires a generation step (suitable for CI).
- When to pick: use `gomock` for large codebases and strict interaction testing; use `testify/mock` for small projects and fast iteration.

Quick usage notes

- Testify: add `github.com/stretchr/testify` and use the `mock` subpackage or `assert/require` helpers.
- Gomock: install `github.com/golang/mock/mockgen`, generate mocks as part of build/test workflow, and use the generated mock types in tests.

Examples to add (optional):

- Small `testify/mock` example showing `On(...).Return(...)` then `AssertExpectations(t)`.
- `gomock` example showing `EXPECT().Return()` and using `gomock.InOrder` for ordered interactions.

