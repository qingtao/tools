#!/bin/sh
#
# Author: wqt.acc@gmail.com
#
### BEGIN INIT INFO
# Provides:          service_example
# Required-Start:    $remote_fs
# Required-Stop:
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: linux service example
# Description:       simple example for linux service daemon start|stop|restart|status
### END INIT INFO

. /lib/lsb/init-functions

BASE="/usr/local/vm_checker"
PIDFILE="${BASE}/run/checker.pid"
NAME="vm_checker"

do_status() {
    if test -f $1; then
        local pid=`cat $1`
        cmd=`ps --no-headers --pid ${pid} -o cmd|grep -o ${NAME}`
        if [ "x${cmd}" = "x${NAME}" ]; then
            return 0
        else
            return 1
        fi
    fi
    return 10
}

case $1 in
    start)
        if do_status $PIDFILE; then
            log_warning_msg "$NAME is running."
            exit 0
        fi
        ${BASE}/$NAME > /dev/null 2>&1 &
        sleep 1
        if do_status $PIDFILE; then
            log_success_msg "Start $NAME success."
            exit 0
        fi
        log_failure_msg "Start $NAME failed."
        exit 1
        ;;
    stop)
        if do_status $PIDFILE; then
            pid=`cat $PIDFILE`
            kill $pid
            sleep 2
            if do_status $PIDFILE; then
                log_failure_msg "Stop $NAME failed."
                exit 1
            else
                rm -f $PIDFILE
                log_success_msg "Stop $NAME success."
                exit 0
            fi
        fi
        log_warning_msg "$NAME not running."
        exit 0
        ;;
    restart)
        $0 stop
        $0 start 
        ;;
    status)
        if do_status $PIDFILE; then
            log_success_msg "$NAME is running: `cat $PIDFILE`."
        else
            log_warning_msg "$NAME not running."
        fi
        ;;
    *)
        echo "Usage: $NAME (start|stop|restart|status)"
        exit 1
esac


