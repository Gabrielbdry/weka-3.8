from threading import Thread
import sys
from queue import Queue
import time
import subprocess


def do_work():
    while True:
        url = q.get()
        get_status()
        q.task_done()


def get_status():
    subprocess.call('curl -X POST "http://0.0.0.0:80/algorithm/BayesNet" -H "accept: text/uri-list" -H '
                    '"Content-Type: multipart/form-data" -F "estimator=SimpleEstimator" -F "estimatorParams=0.5" -F '
                    '"searchAlgorithm=local.K2" -F "useADTree=" -F "validationNum=10" -F "searchParams=-P 1 -S BAYES" '
                    '-F "datasetUri=" -F "validation=CrossValidation" -F "file=@dataset_100.arff"', shell=True)


q = Queue(0)
t = Thread(target=do_work)
t.daemon = True
t.start()

request_per_second = 20

try:
    while True:
        for i in range(0, request_per_second):
            q.put("task")
        time.sleep(1)
except KeyboardInterrupt:
    sys.exit(1)
