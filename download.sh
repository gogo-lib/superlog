#!/bin/bash
FILE() {
    file_name=("cronjob.go" "init.go" "kafka_writer.go" "logger.go" "structure_logger.go" "unstructure_logger.go")

    mkdir -p superlog

    for name in ${file_name[@]}; do
        file_path="https://raw.githubusercontent.com/gogo-lib/superlog/main/$name"
        wget $file_path -O "superlog/$name"
    done
}

FILE