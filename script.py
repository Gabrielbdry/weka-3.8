from threading import Thread
import sys
from Queue import Queue
import time
import subprocess

concurrent = 100

def doWork():
    while True:
        url = q.get()
        getStatus(url)
        q.task_done()

def getStatus(ourl):
    subprocess.call('curl -X POST "http://localhost:8081/algorithm/BayesNet" -H "accept: text/uri-list" -H "Content-Type: multipart/form-data" -F "estimator=SimpleEstimator" -F "estimatorParams=0.5" -F "searchAlgorithm=local.K2" -F "useADTree=" -F "validationNum=10" -F "searchParams=-P 1 -S BAYES" -F "datasetUri=" -F "validation=CrossValidation" -F "file=~/Documents/weka-3.8/contact-lenses.arff"', shell=True)

q = Queue(0)
for i in range(concurrent):
    t = Thread(target=doWork)
    t.daemon = True
    t.start()

request_per_second = 20

try:
    while(True):
        for i in range(0, request_per_second * concurrent):
            q.put("task")
        time.sleep(1)
except KeyboardInterrupt:
    sys.exit(1)

