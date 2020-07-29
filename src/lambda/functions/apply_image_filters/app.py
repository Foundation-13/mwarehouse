import boto3

from model import events
from aws import aws
from image import filters


def lambda_handler(event, context):
    """
    """

    bucket = "foundation-13-temporary"

    parsed_event = events.IncomingEvent(event)
    key = parsed_event.detail.key
    print("Received event({}, {}, {})".format(parsed_event.id, parsed_event.type, key))

    s3 = boto3.client('s3')

    src_buf = aws.get_object(s3, bucket, key)
    converted_buf = filters.monochrome(src_buf)

    new_id = parsed_event.detail.key + "-updated"
    aws.put_object(s3, bucket, new_id, converted_buf)

    return new_id

