
Process Structure
-------------------
Process Control Block

    PID
    Parent ID
    State 
    Priority
    Memory Info
    Open Files
    CPU Registers
    Scheduling Info


Code (Text Segment)

Global Data

Heap

Threads

    Thread Control Block
    Program Counter
    Registers 
    stack (SP)
    State

Open Files

    file descriptors
    etc..

Network Sockets

    tcp connections
    udp connections
    sockets
    ports
    etc..

--------------------


1 a int = 10

2 b int = 20 

3 c int = a + b


## Go routines

GMP scheduler

G: Goroutine
M: Machine (OS Threads)
P: Processor (Logical Go processor, local queue/scheduling)

- Local Run Queue --> goroutines ready to run (256 Goroutines can run)

-> M has a P , P picks a G, G runs on M

When run go application, GOMAXPROCESS(),based on this, few number of threads are created.

- New Goroutine is created, go allocates a structure (info about the Goroutine, similar to Thread)
- Each P process contain Local Run Queue
- The runtime contains Global queue
- If all local queus become full, there is a Global queue


