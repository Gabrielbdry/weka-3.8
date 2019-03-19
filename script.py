from threading import Thread
import sys
from queue import Queue
import time
import subprocess

def doWork():
    while True:
        url = q.get()
        getStatus(url)
        q.task_done()

def getStatus(ourl):
    subprocess.call('curl -X POST "http://0.0.0.0:8081/algorithm/BayesNet" -H "accept: text/uri-list" -H "Content-Type: multipart/form-data" -F "estimator=SimpleEstimator" -F "estimatorParams=0.5" -F "searchAlgorithm=local.K2" -F "useADTree=" -F "validationNum=10" -F "searchParams=-P 1 -S BAYES" -F "datasetUri=" -F "validation=CrossValidation" -F "file=@dataset_100.arff"', shell=True)

q = Queue(0)
t = Thread(target=doWork)
t.daemon = True
t.start()

request_per_second = 1

try:
    while(True):
        for i in range(0, request_per_second):
            q.put("task")
        time.sleep(1)
except KeyboardInterrupt:
    sys.exit(1)

