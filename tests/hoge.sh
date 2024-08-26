#!/usr/bin/env bash

sayHello() {
    echo "Hello World!"
}

nestedFunction() {
    hoge() {
        echo "Hoge"
    }
}
