import boto3

from model.events import IncomingEvent
from aws.aws import get_object, put_object
from image.filters import monochrome


def lambda_handler(event, context):
    """
    """

    bucket = "foundation-13-temporary-bohdan"

    parsed_event = IncomingEvent(event)
    key = parsed_event.detail.key
    print("Received event({}, {}, {})".format(parsed_event.id, parsed_event.type, key))

    s3 = boto3.client('s3')

    src_buf = get_object(s3, bucket, key)
    converted_buf = monochrome(src_buf)

    new_id = parsed_event.detail.key + "-updated"
    put_object(s3, bucket, new_id, converted_buf)

    return new_id

