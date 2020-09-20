#!/bin/bash

# API  127.0.0.1:8080/login_orm test script
# Set the concurrency value 100 500 1000 2000
con_arr="100 500 1000 2000"

exec_single_wrk(){
	for concurrency in $con_arr
	do
	  result_orm=$(wrk -t1 -c${concurrency} -d10 --latency --timeout 5 "http://127.0.0.1:8080/login_orm" >> "login_orm_result.txt" 2>&1)
	  echo >> "login_orm_result.txt" 2>&1
	  wait
	done
}

exec_single_wrk
