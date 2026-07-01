# Golang Concurrency Tasks


## Level 1: Goroutine Basics

### 1. Hello Goroutine

**Objective:** Write a program that starts one goroutine to print `Hello from goroutine` while the main function prints `Hello from main`.

**Concepts to teach:** goroutines, scheduler basics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 2. Multiple Goroutines

**Objective:** Start 5 goroutines, each printing its goroutine number.

**Concepts to teach:** goroutines, scheduler basics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 3. Goroutine With Function Argument

**Objective:** Create a function `printMessage(msg string)` and run it as a goroutine with different messages.

**Concepts to teach:** goroutines, scheduler basics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 4. Goroutine Loop Capture Problem

**Objective:** Start goroutines inside a loop and print the loop index. First reproduce the common closure problem, then fix it.

**Concepts to teach:** goroutines, scheduler basics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 5. Sleep-Based Waiting

**Objective:** Start a goroutine and use `time.Sleep` in `main` to wait for it. Then explain why this is not production-safe.

**Concepts to teach:** sync.WaitGroup, coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 6. WaitGroup Introduction

**Objective:** Rewrite task 5 using `sync.WaitGroup` instead of `time.Sleep`.

**Concepts to teach:** sync.WaitGroup, coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 7. Concurrent Number Printing

**Objective:** Launch 10 goroutines to print numbers from 1 to 10. Ensure the main function waits for all of them.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 8. Concurrent String Printing

**Objective:** Create a slice of names and print each name from a separate goroutine.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 9. Goroutine for CPU Work

**Objective:** Run a CPU-heavy calculation in a goroutine and print the result after it completes.

**Concepts to teach:** goroutines, scheduler basics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 10. Goroutine Count Observation

**Objective:** Use `runtime.NumGoroutine()` to print the number of goroutines before and after launching goroutines.

**Concepts to teach:** goroutines, scheduler basics, observability, safe concurrent logging.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 2: WaitGroup Practice

### 11. Wait for 3 Tasks

**Objective:** Run 3 independent tasks concurrently and wait for all of them using `sync.WaitGroup`.

**Concepts to teach:** sync.WaitGroup, coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 12. Download Simulation

**Objective:** Simulate downloading 5 files using goroutines. Each goroutine should sleep for a different duration.

**Concepts to teach:** channels, communication.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 13. Parallel Square Calculation

**Objective:** Given numbers 1 to 10, calculate their squares concurrently and print them.

**Concepts to teach:** channels, communication.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 14. Parallel API Call Simulation

**Objective:** Simulate 10 API calls concurrently using `time.Sleep` and print when each call completes.

**Concepts to teach:** concurrent I/O, HTTP client/server patterns.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 15. WaitGroup With Error Logging

**Objective:** Run multiple goroutines where some fail. Log errors safely and wait for all goroutines.

**Concepts to teach:** sync.WaitGroup, coordination, observability, safe concurrent logging.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 16. Nested Goroutines

**Objective:** Start goroutines from inside another goroutine and ensure all are waited for correctly.

**Concepts to teach:** goroutines, scheduler basics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 17. WaitGroup Misuse Demo

**Objective:** Intentionally call `wg.Add()` inside a goroutine and observe why it can be unsafe. Then fix it.

**Concepts to teach:** sync.WaitGroup, coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 18. Concurrent File Processing Simulation

**Objective:** Process 20 fake file names concurrently and print the processing status.

**Concepts to teach:** channels, communication.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 19. Parallel Sum by Chunks

**Objective:** Split a slice into chunks and calculate partial sums concurrently.

**Concepts to teach:** channels, communication.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 20. Controlled Goroutine Completion

**Objective:** Start multiple goroutines and print `All tasks completed` only after every goroutine finishes.

**Concepts to teach:** goroutines, scheduler basics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Difference between communication and shared memory.
- When sends and receives block.
- Who should close a channel and why.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 3: Basic Channels

### 21. Send and Receive One Value

**Objective:** Create a channel, send an integer from a goroutine, and receive it in main.

**Concepts to teach:** synchronization, shared state.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 22. String Channel

**Objective:** Send a string message from a goroutine to the main function using a channel.

**Concepts to teach:** channels, send/receive semantics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 23. Channel as Completion Signal

**Objective:** Use a channel to signal that a goroutine has completed its work.

**Concepts to teach:** channels, send/receive semantics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 24. Multiple Sends

**Objective:** Send numbers 1 to 5 from one goroutine and receive them in main.

**Concepts to teach:** synchronization, shared state.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 25. Close Channel

**Objective:** Send several values into a channel, close it, and range over the channel.

**Concepts to teach:** channels, send/receive semantics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 26. Receive From Closed Channel

**Objective:** Experiment with receiving from a closed channel and observe the zero value behavior.

**Concepts to teach:** channels, send/receive semantics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 27. Buffered Channel Basics

**Objective:** Create a buffered channel of size 3 and send 3 values without a receiver.

**Concepts to teach:** channels, send/receive semantics, buffered channels, backpressure.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 28. Buffered Channel Blocking

**Objective:** Create a buffered channel of size 2 and try sending 3 values. Observe the blocking behavior.

**Concepts to teach:** channels, send/receive semantics, buffered channels, backpressure.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 29. Channel Direction

**Objective:** Write functions using send-only and receive-only channel parameters.

**Concepts to teach:** channels, send/receive semantics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 30. Channel Deadlock Demo

**Objective:** Write a small program that causes a channel deadlock. Then fix it.

**Concepts to teach:** channels, send/receive semantics.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How data races happen.
- Mutex vs channel-based ownership.
- How to prove correctness using the race detector.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 4: Select Statement

### 31. Basic Select

**Objective:** Use `select` to receive from two channels.

**Concepts to teach:** select statement, non-blocking coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 32. Fastest Response Wins

**Objective:** Start two goroutines that send responses after different delays. Use `select` to receive the first response.

**Concepts to teach:** select, timeouts, cancellation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 33. Timeout With Select

**Objective:** Use `time.After` to implement a timeout while waiting for a channel response.

**Concepts to teach:** select statement, non-blocking coordination, timeouts, time.After/context deadlines.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 34. Default Case in Select

**Objective:** Use `default` in `select` to avoid blocking when no channel is ready.

**Concepts to teach:** select statement, non-blocking coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 35. Polling With Select

**Objective:** Build a loop that polls a channel using `select` with a default case.

**Concepts to teach:** select statement, non-blocking coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 36. Select With Multiple Workers

**Objective:** Start 3 workers sending results into different channels and consume them with `select`.

**Concepts to teach:** select statement, non-blocking coordination, worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 37. Select With Stop Channel

**Objective:** Use a stop channel to terminate a running goroutine.

**Concepts to teach:** channels, send/receive semantics, select statement, non-blocking coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 38. Select With Ticker

**Objective:** Use `time.Ticker` and `select` to print a message every second.

**Concepts to teach:** select statement, non-blocking coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 39. Select With Timer

**Objective:** Use `time.Timer` to stop a task after a fixed duration.

**Concepts to teach:** select statement, non-blocking coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 40. Select Priority Simulation

**Objective:** Simulate priority handling between high-priority and low-priority channels.

**Concepts to teach:** select statement, non-blocking coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- How `select` chooses ready cases.
- How to implement cancellation safely.
- How timeouts prevent permanently blocked goroutines.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 5: Mutex and Shared State

### 41. Race Condition Demo

**Objective:** Create a shared counter incremented by 100 goroutines without a mutex. Run with `go run -race`.

**Concepts to teach:** sync.Mutex, data race prevention.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Run the program with `go run -race` or `go test -race` and fix any race condition.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 42. Fix Counter With Mutex

**Objective:** Fix task 41 using `sync.Mutex`.

**Concepts to teach:** sync.Mutex, data race prevention.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Run the program with `go run -race` or `go test -race` and fix any race condition.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 43. Counter With RWMutex

**Objective:** Use `sync.RWMutex` to allow multiple readers and safe writers.

**Concepts to teach:** sync.Mutex, data race prevention.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Run the program with `go run -race` or `go test -race` and fix any race condition.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 44. Concurrent Map Write Problem

**Objective:** Write to a normal map from multiple goroutines and observe the issue.

**Concepts to teach:** worker pools, parallel processing.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 45. Safe Map With Mutex

**Objective:** Protect a map using `sync.Mutex`.

**Concepts to teach:** sync.Mutex, data race prevention.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Run the program with `go run -race` or `go test -race` and fix any race condition.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 46. Safe Map With sync.Map

**Objective:** Rewrite task 45 using `sync.Map`.

**Concepts to teach:** worker pools, parallel processing.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 47. Bank Account Simulation

**Objective:** Create a bank account balance and perform concurrent deposits and withdrawals safely.

**Concepts to teach:** worker pools, parallel processing.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 48. Inventory Update

**Objective:** Multiple goroutines update product inventory. Protect the shared inventory safely.

**Concepts to teach:** worker pools, parallel processing.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 49. Read-Heavy Cache

**Objective:** Build a simple read-heavy cache using `sync.RWMutex`.

**Concepts to teach:** worker pools, parallel processing.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 50. Atomic Counter

**Objective:** Use `sync/atomic` to implement a concurrent counter.

**Concepts to teach:** sync.Mutex, data race prevention.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Run the program with `go run -race` or `go test -race` and fix any race condition.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Throughput vs number of workers.
- Backpressure and bounded queues.
- Error collection and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 6: Worker Pools

### 51. Basic Worker Pool

**Objective:** Create 3 workers that process 10 jobs from a jobs channel.

**Concepts to teach:** worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 52. Worker Pool With Results

**Objective:** Workers should process jobs and send results to a results channel.

**Concepts to teach:** worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 53. Worker Pool With WaitGroup

**Objective:** Use `sync.WaitGroup` to wait for all workers to complete.

**Concepts to teach:** sync.WaitGroup, coordination, worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 54. Dynamic Job Submission

**Objective:** Submit jobs into a worker pool from another goroutine.

**Concepts to teach:** pipelines, stream processing.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 55. Worker Pool With Errors

**Objective:** Each worker may return an error. Collect successful results and errors separately.

**Concepts to teach:** worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 56. Worker Pool With Context Cancellation

**Objective:** Use `context.Context` to cancel all workers.

**Concepts to teach:** context.Context, cancellation propagation, worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 57. Worker Pool With Timeout

**Objective:** Stop processing if the entire worker pool exceeds a timeout.

**Concepts to teach:** timeouts, time.After/context deadlines, worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 58. Worker Pool With Rate Limit

**Objective:** Limit how frequently workers can process jobs using `time.Ticker`.

**Concepts to teach:** worker pool, job queues, rate limiting, ticker/token bucket.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 59. Worker Pool for File Processing

**Objective:** Process a list of fake file names using a configurable number of workers.

**Concepts to teach:** worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 60. Worker Pool Benchmark

**Objective:** Compare sequential processing and worker-pool processing using Go benchmarks.

**Concepts to teach:** worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Pipeline stage ownership.
- Fan-out/fan-in trade-offs.
- How to close downstream channels safely.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 7: Pipelines

### 61. Simple Pipeline

**Objective:** Build a pipeline with three stages: generate numbers, square numbers, print results.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 62. Pipeline With Channel Closing

**Objective:** Ensure every pipeline stage closes its output channel correctly.

**Concepts to teach:** channels, send/receive semantics, pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 63. Pipeline With Filtering

**Objective:** Generate numbers, filter even numbers, square them, and print results.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 64. Pipeline With Cancellation

**Objective:** Add `context.Context` cancellation to the pipeline.

**Concepts to teach:** context.Context, cancellation propagation, pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 65. Pipeline With Error Channel

**Objective:** Add an error channel to report errors from pipeline stages.

**Concepts to teach:** channels, send/receive semantics, pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 66. Pipeline Fan-Out

**Objective:** Send generated numbers to multiple worker goroutines for processing.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 67. Pipeline Fan-In

**Objective:** Merge results from multiple worker channels into one output channel.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 68. Pipeline Backpressure

**Objective:** Use buffered and unbuffered channels to observe backpressure behavior.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 69. Pipeline With Slow Consumer

**Objective:** Create a fast producer and slow consumer. Observe blocking and fix using buffering.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 70. Production-Style Pipeline

**Objective:** Build a pipeline that reads tasks, validates them, enriches them, and stores results.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 8: Fan-Out, Fan-In, and Coordination

### 71. Fan-Out HTTP Simulation

**Objective:** Send 20 fake HTTP requests across 5 workers and collect responses.

**Concepts to teach:** pipelines, fan-out/fan-in, concurrent I/O, HTTP client/server patterns.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 72. Fan-In Merge Function

**Objective:** Write a reusable `merge()` function that combines multiple channels into one.

**Concepts to teach:** pipelines, fan-out/fan-in.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 73. First Successful Result

**Objective:** Start multiple goroutines and return the first successful result while cancelling the rest.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 74. Broadcast Message

**Objective:** Send the same message to multiple goroutines using separate channels.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 75. Publish-Subscribe Simulation

**Objective:** Implement a simple pub-sub system using channels.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 76. Task Dispatcher

**Objective:** Create a dispatcher that sends jobs to different worker queues.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 77. Load Balancer Simulation

**Objective:** Distribute jobs among workers based on availability.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 78. Result Aggregator

**Objective:** Aggregate results from multiple goroutines and produce a final summary.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 79. Parallel Search

**Objective:** Search for a value across multiple slices concurrently and stop when found.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 80. Concurrent Web Crawler Simulation

**Objective:** Simulate crawling URLs concurrently with depth control and duplicate prevention.

**Concepts to teach:** advanced concurrency, system design.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 9: Context, Cancellation, and Timeouts

### 81. Context Cancellation Basics

**Objective:** Start a goroutine that stops when the context is cancelled.

**Concepts to teach:** context.Context, cancellation propagation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 82. Context Timeout

**Objective:** Create a context with timeout and stop work automatically after timeout.

**Concepts to teach:** timeouts, time.After/context deadlines, context.Context, cancellation propagation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 83. Context Deadline

**Objective:** Use `context.WithDeadline` to stop work at a specific time.

**Concepts to teach:** timeouts, time.After/context deadlines, context.Context, cancellation propagation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 84. Context Values

**Objective:** Pass request metadata using context values and read it inside goroutines.

**Concepts to teach:** context.Context, cancellation propagation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 85. Cancel Long-Running Task

**Objective:** Create a long-running task and cancel it from main.

**Concepts to teach:** context.Context, cancellation propagation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 86. Cancel Multiple Workers

**Objective:** Start multiple workers and cancel all using one context.

**Concepts to teach:** context.Context, cancellation propagation, worker pool, job queues.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.
- Close channels from the sender side and document the ownership of each channel.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.
- All jobs are processed exactly once and all workers stop cleanly.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 87. Cascading Cancellation

**Objective:** Create parent and child contexts and observe cancellation propagation.

**Concepts to teach:** context.Context, cancellation propagation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 88. HTTP Request With Context

**Objective:** Make an HTTP request using context timeout.

**Concepts to teach:** context.Context, cancellation propagation, concurrent I/O, HTTP client/server patterns.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 89. Database Query Simulation With Context

**Objective:** Simulate a DB query that stops when context is cancelled.

**Concepts to teach:** context.Context, cancellation propagation.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Avoid goroutine leaks by ensuring every goroutine has a clear exit path.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Program exits cleanly after timeout or cancellation without hanging.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 90. Graceful Shutdown

**Objective:** Implement graceful shutdown using context, signal handling, and worker cleanup.

**Concepts to teach:** sync.Mutex, data race prevention.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.
- Run the program with `go run -race` or `go test -race` and fix any race condition.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Production failure modes.
- Observability and debugging.
- Resource leaks, race conditions, and graceful shutdown.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Level 10: Advanced and Production-Grade Tasks

### 91. Semaphore With Buffered Channel

**Objective:** Limit concurrency to 5 goroutines using a buffered channel as a semaphore.

**Concepts to teach:** channels, send/receive semantics, buffered channels, backpressure.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.
- Race detector reports no data races after the fix.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 92. Bounded Parallelism

**Objective:** Process 1000 jobs but allow only 10 concurrent goroutines at a time.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 93. ErrGroup Usage

**Objective:** Use `errgroup.Group` to run multiple concurrent tasks and stop on first error.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 94. Retry With Backoff

**Objective:** Implement concurrent workers where failed jobs are retried with exponential backoff.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 95. Circuit Breaker Simulation

**Objective:** Build a simple circuit breaker around concurrent service calls.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 96. Concurrent Rate Limiter

**Objective:** Create a token-bucket-style rate limiter using channels and ticker.

**Concepts to teach:** rate limiting, ticker/token bucket.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 97. In-Memory Job Queue

**Objective:** Build an in-memory job queue with producers, consumers, retries, and graceful shutdown.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 98. Concurrent Log Processor

**Objective:** Read log lines from a channel, parse them concurrently, and aggregate statistics.

**Concepts to teach:** observability, safe concurrent logging.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 99. Concurrent Metrics Collector

**Objective:** Collect metrics from multiple fake services concurrently and expose aggregated results.

**Concepts to teach:** observability, safe concurrent logging.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

### 100. Mini Concurrent Task Scheduler

**Objective:** Build a mini scheduler that accepts tasks, runs them with bounded concurrency, supports cancellation, retries failed tasks, and shuts down gracefully.

**Concepts to teach:** goroutines, basic coordination.

**Requirements:**
- Create a small standalone Go program for this task.
- Use clear function names and keep the concurrency logic visible instead of hiding it inside large helper functions.
- Print meaningful output so the learner can observe ordering, blocking, completion, or cancellation behavior.
- Add comments explaining why goroutines/channels/synchronization are used.

**Acceptance criteria:**
- Program compiles and runs without panic.
- Main function waits correctly for all required goroutines.
- Output demonstrates the intended concurrency behavior.

**Trainer discussion points:**
- Why goroutines are lightweight but not free.
- Why `time.Sleep` is not proper synchronization.
- How Go scheduler behavior can make output order nondeterministic.

**Extension idea:** Convert the program into a table-driven test or benchmark where applicable.

---

## Suggested Training Flow

- **Day 1:** Goroutines, WaitGroups, basic channels, channel closing rules.
- **Day 2:** Buffered channels, select, timeout, cancellation, and context.
- **Day 3:** Mutexes, atomics, race detector, safe shared state, concurrent maps.
- **Day 4:** Worker pools, fan-out/fan-in, pipelines, error handling, graceful shutdown.
- **Day 5:** Production patterns: rate limiting, retries, observability, HTTP concurrency, leak detection, and capstone project.

## Evaluation Rubric

| Area | What to Check |
|---|---|
| Correctness | Program gives expected result and handles completion properly. |
| Safety | No data races, deadlocks, or goroutine leaks. |
| Idiomatic Go | Uses channels, WaitGroups, Mutexes, Context, and Select appropriately. |
| Readability | Code is small, clear, and explainable in a classroom. |
| Production Thinking | Includes timeout, cancellation, error handling, and shutdown where relevant. |