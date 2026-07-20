package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a context that automatically cancels when SIGINT (Ctrl+C) or SIGTERM is received
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop() // Clean up signal listeners when main exits

	fmt.Println("Program running... Press Ctrl+C to stop.")

	// Pass the signal-aware context to long-running tasks
	if err := doWork(ctx); err != nil {
		fmt.Printf("Work stopped with error: %v\n", err)
	}

	fmt.Println("Graceful shutdown complete.")
}

func doWork(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			// Triggered when Ctrl+C is pressed
			fmt.Println("\nReceived interrupt signal! Cleaning up...")

			// Simulate cleanup work (e.g., closing DB connections, finishing pending writes)
			time.Sleep(5 * time.Second)

			return ctx.Err() // Returns context.Canceled

		default:
			// Perform regular work
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
		}
	}
}
