go testing  
 
some commands  
`
go test
`  
`
go test -v
`  
test coverage  
`  
go test -cover
`  
generating coverage report  
`
go test -coverprofile=coverage.out
`  
show the detail report  
`
go tool cover -html=coverage.out
`  
performance test use Benchmark prefix and testing.B  
`
go test -bench=.
`  
Regex pattern  
`
go test -run=Calc -bench=.
`  
