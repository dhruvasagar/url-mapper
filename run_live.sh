#!/bin/sh

PID_FILE='url-mapper.pid'
LOG_FILE='url-mapper.log'

restart() {
  if [ -f $PID_FILE ]; then
    echo "restarting server..."
    pkill -F $PID_FILE
    rm -f $PID_FILE
  fi
  ./dsis.me > $LOG_FILE 2>&1 &
  echo $! > $PID_FILE
}

restart
while read event
do
  echo "$event"
  restart
  # tail -f $LOG_FILE
done
