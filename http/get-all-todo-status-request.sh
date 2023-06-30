#!/bin/bash

http -vv --session=./session.json GET :3000/v1/todo/status
