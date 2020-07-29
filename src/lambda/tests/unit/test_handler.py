from apply_image_filters import app


def test_lambda_handler(image_filters_event):
    key = app.lambda_handler(image_filters_event, "")

    assert key == "bsg3g3jd0cvm4p91ddf0"