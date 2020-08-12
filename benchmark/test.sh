#!/bin/bash

# API  127.0.0.1:8080/login_orm
# 1 100 500 1000 2000 5000
con_arr = (100 500)

exec_single_wrk(){
	for concurrency in $con_arr
	do 
		result=$(wrk -t$concurrency -c 1 -d 10 --latency --timeout 3 http://127.0.0.1:8080/login_orm >> "result.txt" 2>&1)
		# ./wrk -t ${concurrency} -c ${count} -d ${continueTime} --latency --timeout 3 "http://$input_api" >$RESULT_DIR$1$2".txt" 2>&1 &
	done 
}

exec_single_wrk