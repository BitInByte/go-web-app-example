#!/bin/bash

http -vv --session=./session.json PUT :3000/v1/todo/status/1 
