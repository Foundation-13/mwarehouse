from functions.apply_image_filters.app import lambda_handler


def test_lambda_handler(image_filters_event, s3, image):
    bucket = "foundation-13-temporary"
    key = "bsg3g3jd0cvm4p91ddf0"

    s3.create_bucket(Bucket=bucket)
    s3.put_object(Bucket=bucket, Key=key, Body=image)

    key = lambda_handler(image_filters_event, "")

    assert key == "bsg3g3jd0cvm4p91ddf0-updated"
