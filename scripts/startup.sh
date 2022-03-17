#!/bin/bash
APPLICATION_NAME=$1
DELAY=$2
RUN_PID=''
PID_FILE=""
startup() {
    #把项目送入后台执行，然后把pid写入对应的文件中
    nohup "./$APPLICATION_NAME" -c conf/config.yml > nohup.out 2>&1 & echo $! > "${PID_FILE}"
    if [ -n "$!" ]; then
        echo "已启动${APPLICATION_NAME}..."
        echo "进程pid：$!"
    fi
    exit 0
}

getRunId() {
  RUN_PID=$(ps aux | grep "${APPLICATION_NAME}" | grep -v "grep" | grep -v "$0" | awk '{print $2}')
}

killProcess() {
  times=1
  while true
  do
      getRunId
      if [ -n "$RUN_PID" ]; then
          if [ ${times} -ge "${DELAY}" ]; then
            echo "强制杀死${APPLICATION_NAME}..."
            kill -9 $RUN_PID
            break
          else
            times=$((times + 1))
            echo "等待${times}秒..."
            sleep 1
          fi
      else
        break
      fi
  done
}
source /etc/profile
if [ -z "$APPLICATION_NAME" ]; then
    echo "请输入启动的任务名..."
    exit 1
fi
PID_FILE="${APPLICATION_NAME}.pid"
if [ -z "$DELAY" ]
then
    DELAY=30
fi
#pid存在时就从文件获取pid
if [ -e "${PID_FILE}" ]
then
    RUN_PID=$(cat "${PID_FILE}")
    if [ "$RUN_PID" ]
    then
      # 如果执行失败, 可能没有这个pid或者这个pid已经改变，重新查询后去kill
      if ! kill -15 "$RUN_PID" > /dev/null 2>&1; then
          echo "${RUN_PID} 进程停止失败!可能没有这个pid或者这个pid已经改变"
          getRunId
          echo "开始动态获取进程pid..."
          if [ -z "$RUN_PID" ] || ! kill -15 $RUN_PID > /dev/null 2>&1; then
              startup
          fi
      fi
      # 最多循环等待$DELAY秒，如果应用还没有退出就强制杀死
      killProcess && startup
    fi
else
    #第一次启动
    touch "${PID_FILE}"
    startup;
    exit 0
fi