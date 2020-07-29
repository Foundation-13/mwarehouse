import pytest


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
