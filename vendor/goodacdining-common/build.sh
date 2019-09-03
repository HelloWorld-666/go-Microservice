#!/bin/bash
#
#  Auto build shell script

# 此处需要设置或修改MAVEN_OPTS，否则在执行mvn install命令时可能会出现OutOfMemoryError错误
export MAVEN_OPTS="-Xmx512m -XX:MaxPermSize=128m -XX:MaxMetaspaceSize=128m"

if [ $# -lt 1 ]; then
  echo "Usage: $0 [r|d]";
  exit;

else
 
    case $1 in
      r | release)
        ENV="release"
        ;;
      d | debug)
        ENV="debug"
        ;;
      *)
        echo "Error! unknown parameter."
        exit 1
        ;;
    esac
     
    mvn clean install -Dmaven.test.skip=true -Dmaven.compile.fork=true -U -P${ENV}
fi