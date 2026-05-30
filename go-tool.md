# escape analysis
go build -gcflags="-m"

# disable optimization
go build -gcflags="-N -l"

# assembly
go tool compile -S main.go

# symbols
go tool nm demo

# disassembly
go tool objdump demo

# macOS sections
otool -l demo

# SSA
GOSSAFUNC=main go build

# debugger
dlv exec ./demo