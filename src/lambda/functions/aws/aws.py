from timeit import default_timer as timer


def get_object(client, bucket, key):
    start = timer()
    print("aws event({})".format(key))
    buf = client.get_object(Bucket=bucket, Key=key)['Body'].read()
    print("read from bucket: latency={}".format(timer() - start))
    return buf


def put_object(client, bucket, key, buffer):
    start = timer()
    sent_data = client.put_object(Bucket=bucket, Key=key, Body=buffer)
    code = sent_data['ResponseMetadata']['HTTPStatusCode']
    print("write to the bucket: code={}, latency={}".format(code, timer() - start))
    return code == 200

