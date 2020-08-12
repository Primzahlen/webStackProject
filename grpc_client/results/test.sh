#!/bin/bash

RESULT_DIR="/Users/chauncey/Desktop/Semester3/webStackProject/grpc_client/results"
#luaFileArr="ff0e_00.lua ff0e_02.lua ff0e_05.lua ff0e_06.lua ff0e_07.lua ff0e_08.lua ff0e_10.lua ff0e_11.lua" # 每个接口post需要的json参数的lua文件
# 127.0.0.1:8080/send_message

concurrency=$1 #开启线程数，用于控制速度
count=$2  #保持连接数
continueTime=$3  # 持续时间
input_api=$4    #请求API

OLD_IFS="$IFS"
IFS="/"
api_arr=($input_api)
IFS="$OLD_IFS"

api_func=${api_arr[2]}   #请求的函数
echo $api_func
exec_single_wrk(){
		result=$(wrk -t${concurrency} -c${count} -d${continueTime} --latency --timeout 3 "http://$input_api" >> "result.txt" 2>&1)
		# ./wrk -t ${concurrency} -c ${count} -d ${continueTime} --latency --timeout 3 "http://$input_api" >$RESULT_DIR$1$2".txt" 2>&1 &

}

exec_single_wrk $concurrency $count $continueTime $input_api