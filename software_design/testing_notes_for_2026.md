# Testing Notes for 2026

## Summary of "Testing in 2021" by Tim Bray
Tim Bray's 2021 article emphasizes that testing remains essential for software reliability, despite advances in tools and practices. Key points include:
- **Unit Tests**: Fast, cheap, and foundational; use frameworks for isolated testing.
- **Integration Tests**: Slower but critical for end-to-end validation; integrate with CI/CD pipelines.
- **Property-Based Testing**: Tools like QuickCheck generate edge-case tests automatically.
- **Fuzzing**: Essential for finding bugs in inputs; evolving with better automation.
- **Challenges**: Flaky tests, maintenance overhead, and balancing speed vs. coverage.
- **Trends**: Shift to continuous testing, better tooling for distributed systems, and integration with DevOps.

Bray advocates for pragmatic testing strategies, prioritizing developer experience and reliability over perfection.

## Updates for 2026
By 2026, testing has evolved with cloud-native architectures, and security focus. Key advancements, incorporating insights from Martin Fowler's "The Testing Pyramid" (emphasizing a balanced test suite with many fast unit tests, fewer integration tests, and minimal e2e tests for efficient feedback) and Ian Cooper's "TDD, Where Did It All Go Wrong" (focusing on testing behaviors over implementation details, the red-green-refactor cycle, and avoiding over-mocking):

- **Property-based Testing**:  Property-based testing now uses ML to optimize test cases, aligning with the pyramid's focus on fast unit tests and Cooper's emphasis on behavior-driven tests to avoid brittle, implementation-coupled suites.
- **Chaos Engineering and Reliability**: Widespread adoption of tools like Chaos Monkey for production simulations; testing for resilience in microservices and serverless environments, extending integration tests in the pyramid and Cooper's ports-and-adapters architecture for decoupling tests from internals.
- **Testing in Production**: Feature flags, canary releases, and observability (e.g., via OpenTelemetry) enable safe production testing without full rollouts, reducing reliance on slow e2e tests and supporting Cooper's critique of heavy, slow test suites.
- **Security Integration**: Fuzzing extended to security (e.g., OSS-Fuzz for vulnerabilities); automated SAST/DAST in CI pipelines, integrated into unit and integration layers, with AI aiding in generating behavior-focused security tests.
- **Distributed and Cloud Testing**: Enhanced for Kubernetes and multi-cloud; tools like Testkube for k8s-native testing, emphasizing contract testing for microservices to avoid the "ice cream cone" anti-pattern (too many e2e) and Cooper's advice on testing at ports rather than adapters.
- **Sustainability**: Energy-efficient testing practices, reducing compute waste in large test suites, prioritizing fast unit tests and Cooper's red-green-refactor cycle to minimize maintenance overhead.
- **Emerging Tools**: Low-code platforms (e.g., Testim) for non-developers; quantum-resistant testing for emerging tech, with AI helping maintain pyramid balance and refactor code without breaking behavior tests.
- **Challenges Persist**: managing test debt in fast-evolving codebases. Emphasis on "shift-left" with early, automated feedback, guided by the testing pyramid's structure and Cooper's reboot of TDD to test stable APIs, not implementation details—reducing mocking and enabling easier refactoring in microservices.

Overall, testing in 2026 is more intelligent, integrated, and production-oriented, but still requires human oversight for quality, reinforced by Cooper's insights on avoiding TDD pitfalls like over-testing internals and embracing behavior-focused, maintainable suites.

## Guidelines

Guidelines, not rules

- You should write tests.
- You should write tests at the same time as you write your code.
- Each Go package is a self contained unit.
- Your tests should assert the observable behaviour of your package, not its implementation.
- You should design your packages around their behaviour, not their implementation.
