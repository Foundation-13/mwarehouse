from model import events


def test_parse_event(image_filters_event):
    e = events.IncomingEvent(image_filters_event)

    assert e.id == "7bf73129-1428-4cd3-a780-95db273d1602"
    assert e.type == "mwarehouse.image.apply-filters"

    assert e.detail.key == "bsg3g3jd0cvm4p91ddf0"
