import os
import pytest
import moto
import boto3
from io import BytesIO
from PIL import Image


@pytest.fixture()
def image_filters_event():
    """ Generates 'apply-image-filters' event"""

    return {
        "id": "7bf73129-1428-4cd3-a780-95db273d1602",
        "detail-type": "mwarehouse.image.apply-filters",
        "source": "mwarehouse.lambda",
        "account": "123456789012",
        "time": "2015-11-11T21:29:54Z",
        "region": "us-east-2",
        "version": "0",
        "resources": [],
        "detail": {
            "key": "bsg3g3jd0cvm4p91ddf0"
        }
    }


@pytest.fixture()
def image():
    img = Image.new('RGB', (100, 1000), color='red')
    buffer = BytesIO()
    img.save(buffer, 'PNG')
    buffer.seek(0)
    return buffer


@pytest.fixture(scope='function')
def aws_credentials():
    """Mocked AWS Credentials for moto."""
    os.environ['AWS_ACCESS_KEY_ID'] = 'testing'
    os.environ['AWS_SECRET_ACCESS_KEY'] = 'testing'
    os.environ['AWS_SECURITY_TOKEN'] = 'testing'
    os.environ['AWS_SESSION_TOKEN'] = 'testing'


@pytest.fixture(scope='function')
def s3(aws_credentials):
    with moto.mock_s3():
        yield boto3.client('s3', region_name='us-east-1')
